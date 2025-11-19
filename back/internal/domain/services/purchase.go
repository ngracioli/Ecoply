package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/resources"
	"ecoply/internal/domain/utils"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type PurchaseService interface {
	Create(request *requests.CreatePurchase, offerUuid string, user *models.User) *merr.ResponseError
	List(request *requests.ListPurchase, user *models.User) (*utils.PaginationWrapper[*resources.Purchase], *merr.ResponseError)
	Cancel(pruchaseUuid string, user *models.User) *merr.ResponseError
}

type purchaseService struct {
	db           *gorm.DB
	purchaseRepo repository.PurchaseRepository
	offerRepo    repository.OfferRepository
}

func NewPurchaseService(db *gorm.DB) PurchaseService {
	return &purchaseService{
		db:           db,
		purchaseRepo: repository.NewPurchaseRepository(db),
		offerRepo:    repository.NewOfferRepository(db),
	}
}

func (s *purchaseService) Create(
	request *requests.CreatePurchase,
	offerUuid string,
	user *models.User,
) *merr.ResponseError {
	var errResponse *merr.ResponseError
	var offer *models.Offer

	var err error = s.db.Transaction(func(tx *gorm.DB) error {
		var purchase *models.Purchase
		var err error

		offer, err = s.offerRepo.WithTransaction(tx).GetByUuid(offerUuid)
		if err != nil {
			errResponse = merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
			return err
		}

		if offer.RemainingQuantityMwh < request.QuantityMwh {
			errResponse = merr.NewResponseError(http.StatusUnprocessableEntity, ErrInsufficientOfferQuantity)
			return err
		}

		if offer.SellerId == user.ID {
			errResponse = merr.NewResponseError(http.StatusForbidden, ErrCannotPurchaseOwnOffer)
			return err
		}

		if offer.IsFulfilled() || offer.IsExpired() {
			errResponse = merr.NewResponseError(http.StatusUnprocessableEntity, ErrOfferHasEnded)
			return err
		}

		offer.RemainingQuantityMwh -= request.QuantityMwh

		if offer.IsFresh() {
			offer.Status = models.OfferStatusOpen
		} else if offer.RemainingQuantityMwh == 0 {
			offer.Status = models.OfferStatusFulfilled
		}

		if err = s.offerRepo.WithTransaction(tx).Update(offer); err != nil {
			return err
		}

		purchase = &models.Purchase{
			Uuid:          NewUuidV7String(),
			OfferId:       offer.ID,
			PricePerMwh:   offer.PricePerMwh,
			PaymentMethod: request.PaymentMethod,
			Status:        models.PurchaseStatusCompleted,
			BuyerId:       user.ID,
			QuantityMwh:   request.QuantityMwh,
		}

		if err = s.purchaseRepo.WithTransaction(tx).Create(purchase); err != nil {
			return err
		}

		return nil
	})

	if errResponse != nil {
		return errResponse
	}

	if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return nil
}

func (s *purchaseService) List(
	request *requests.ListPurchase,
	user *models.User,
) (*utils.PaginationWrapper[*resources.Purchase], *merr.ResponseError) {
	var list *utils.PaginationWrapper[*models.Purchase]
	var err error

	list, err = s.purchaseRepo.List(uint64(user.ID), request)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var response utils.PaginationWrapper[*resources.Purchase]

	response.Page = list.Page
	response.PageSize = list.PageSize
	response.HasNext = list.HasNext
	response.HasPrev = list.HasPrev
	response.Data = make([]*resources.Purchase, 0, len(list.Data))

	for _, purchase := range list.Data {
		response.Data = append(response.Data, makePurchaseResourceFromModel(purchase))
	}

	return &response, nil
}

func makePurchaseResourceFromModel(purchase *models.Purchase) *resources.Purchase {
	var createdAt time.Time = utils.TruncateDateToLocal(purchase.CreatedAt)

	return &resources.Purchase{
		Uuid:          purchase.Uuid,
		QuantityMwh:   purchase.QuantityMwh,
		PricePerMwh:   purchase.PricePerMwh,
		Status:        purchase.Status,
		PaymentMethod: purchase.PaymentMethod,
		OfferUuid:     purchase.Offer.Uuid,
		SellerUuid:    purchase.Offer.Seller.Uuid,
		CreatedAt:     createdAt.Format(time.RFC3339),
	}
}

func (s *purchaseService) Cancel(purchaseUuid string, user *models.User) *merr.ResponseError {
	var purchase *models.Purchase
	var err error
	var responseErr *merr.ResponseError

	s.db.Transaction(func(tx *gorm.DB) error {
		purchase, err = s.purchaseRepo.WithTransaction(tx).FindByUuid(purchaseUuid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				responseErr = merr.NewResponseError(http.StatusNotFound, ErrPurchaseNotFound)
				return ErrPurchaseNotFound
			}
			responseErr = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			return ErrInternal
		}

		if user.ID != purchase.BuyerId {
			responseErr = merr.NewResponseError(http.StatusForbidden, ErrUserIsNotThePurchaseOwner)
			return ErrUserIsNotThePurchaseOwner
		}

		if !isPurchaseCancelable(purchase) {
			responseErr = merr.NewResponseError(http.StatusUnprocessableEntity, ErrPurchaseCannotBeCancelled)
			return ErrPurchaseCannotBeCancelled
		}

		purchase.Status = models.PurchaseStatusCanceled
		if err := s.purchaseRepo.Update(purchase); err != nil {
			responseErr = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			return ErrInternal
		}

		offer, err := s.offerRepo.GetById(purchase.OfferId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				responseErr = merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
				return ErrPurchaseNotFound
			}
			responseErr = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			return ErrInternal
		}

		offer.RemainingQuantityMwh += purchase.QuantityMwh
		if offer.IsFulfilled() {
			offer.Status = models.OfferStatusOpen
		}

		if err := s.offerRepo.Update(offer); err != nil {
			responseErr = merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
			return ErrInternal
		}

		return nil
	})

	if responseErr != nil {
		return responseErr
	}

	return nil
}

func isPurchaseCancelable(purchase *models.Purchase) bool {
	var createdAt time.Time = utils.TruncateDateToLocal(purchase.CreatedAt)
	var maxCancelDate = createdAt.Add(time.Hour * 2)
	var now = utils.NowInLocal()

	return !now.After(maxCancelDate) && !purchase.IsCancelled()
}
