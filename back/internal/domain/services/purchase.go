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
	var offer *models.Offer
	var err error

	offer, err = s.offerRepo.GetByUuid(offerUuid)
	if err != nil {
		return merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
	}

	if offer.RemainingQuantityMwh < request.QuantityMwh {
		return merr.NewResponseError(http.StatusUnprocessableEntity, ErrInsufficientOfferQuantity)
	}

	if offer.SellerId == user.ID {
		return merr.NewResponseError(http.StatusForbidden, ErrCannotPurchaseOwnOffer)
	}

	if offer.IsFulfilled() || offer.IsExpired() {
		return merr.NewResponseError(http.StatusUnprocessableEntity, ErrOfferHasEnded)
	}

	offer.RemainingQuantityMwh -= request.QuantityMwh

	if offer.IsFresh() {
		offer.Status = models.OfferStatusOpen
	} else if offer.RemainingQuantityMwh == 0 {
		offer.Status = models.OfferStatusFulfilled
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		var offerRepository repository.OfferRepository = repository.NewOfferRepository(tx)
		var purchaseRepository repository.PurchaseRepository = repository.NewPurchaseRepository(tx)
		var purchase *models.Purchase
		var err error

		if err = offerRepository.Update(offer); err != nil {
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

		if err = purchaseRepository.Create(purchase); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return nil
}
