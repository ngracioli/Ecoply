package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/resources"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type ContractService interface {
	Get(user *models.User, purchaseUuid string) (*resources.Contract, *merr.ResponseError)
}

type contractService struct {
	db           *gorm.DB
	offerRepo    repository.OfferRepository
	userRepo     repository.UserRepository
	purchaseRepo repository.PurchaseRepository
}

func NewContractService(db *gorm.DB) ContractService {
	return &contractService{
		db:           db,
		offerRepo:    repository.NewOfferRepository(db),
		userRepo:     repository.NewUserRepository(db),
		purchaseRepo: repository.NewPurchaseRepository(db),
	}
}

func (s *contractService) Get(user *models.User, purchaseUuid string) (*resources.Contract, *merr.ResponseError) {
	var purchase *models.Purchase
	var offer *models.Offer
	var buyer *models.User
	var supplier *models.User

	var err error

	if purchase, err = s.purchaseRepo.FindByUuid(purchaseUuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrPurchaseNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if !purchase.IsCompleted() {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrPurchaseIsNotCompleted)
	}

	if offer, err = s.offerRepo.GetById(purchase.OfferId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrPurchaseNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if buyer, err = s.userRepo.FindById(purchase.BuyerId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrPurchaseNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if supplier, err = s.userRepo.FindById(offer.SellerId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrPurchaseNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if !purchase.IsOwner(user) && !offer.IsOwner(user) {
		return nil, merr.NewResponseError(http.StatusForbidden, ErrUserIsNotContractMember)
	}

	if err = s.db.Preload("Agent").Preload("Agent.Submarket").First(supplier).Error; err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	if err = s.db.Preload("Agent").Preload("Agent.Submarket").First(buyer).Error; err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	if err = s.db.Preload("EnergyType").Preload("Submarket").First(offer).Error; err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var resource *resources.Contract = &resources.Contract{
		Supplier: resources.ContractSupplier{
			Uuid:          supplier.Uuid,
			Cnpj:          supplier.Agent.Cnpj,
			CceeCode:      supplier.Agent.CceeCode,
			SubmarketName: supplier.Agent.Submarket.Name,
			CompanyName:   supplier.Agent.CompanyName,
		},
		Buyer: resources.ContractBuyer{
			Uuid:          buyer.Uuid,
			Cnpj:          buyer.Agent.Cnpj,
			CceeCode:      buyer.Agent.CceeCode,
			SubmarketName: buyer.Agent.Submarket.Name,
			CompanyName:   buyer.Agent.CompanyName,
		},
		Offer: resources.ContractOffer{
			Uuid:                  offer.Uuid,
			PricePerMwh:           offer.PricePerMwh,
			InitialQuantityMwh:    offer.InitialQuantityMwh,
			ContractedQuantityMwh: purchase.QuantityMwh,
			Description:           offer.Description,
			PeriodStart:           offer.PeriodStart.Format(time.DateOnly),
			PeriodEnd:             offer.PeriodEnd.Format(time.DateOnly),
			EnergyType:            offer.EnergyType.Type,
			Submarket:             offer.Submarket.Name,
			CreatedAt:             offer.CreatedAt,
		},
	}

	return resource, nil
}
