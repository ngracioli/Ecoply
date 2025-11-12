package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/requests"
	"net/http"

	"gorm.io/gorm"
)

type PurchaseService interface {
	Create(request *requests.CreatePurchase, offerUuid string, user *models.User) *merr.ResponseError
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
			Uuid:        NewUuidV7String(),
			OfferId:     offer.ID,
			PricePerMwh: offer.PricePerMwh,
			Status:      models.PurchaseStatusCompleted,
			BuyerId:     user.ID,
			QuantityMwh: request.QuantityMwh,
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
