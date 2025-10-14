package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"time"

	"gorm.io/gorm"
)

type OfferRepository interface {
	GetByUuid(uuid string) (*models.Offer, error)
	GetBySellerId(userId uint) ([]*models.Offer, error)
	Create(params *OfferCreateParams) (*models.Offer, error)
	Update(offer *models.Offer) error
	Delete(uuid string) error
}

type offerRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) OfferRepository {
	return &offerRepository{db: db}
}

func (r *offerRepository) GetByUuid(uuid string) (*models.Offer, error) {
	var offer models.Offer
	if err := r.db.Preload("Submarket").
		Preload("EnergyType").
		Preload("Seller").
		Where("uuid = ?", uuid).First(&offer).Error; err != nil {
		return nil, err
	}
	return &offer, nil
}

func (r *offerRepository) GetBySellerId(userId uint) ([]*models.Offer, error) {
	var offers []*models.Offer
	if err := r.db.Preload("Submarket").
		Preload("EnergyType").
		Preload("Seller").
		Where("seller_id = ?", userId).
		Find(&offers).Error; err != nil {
		mlog.Log("Failed to get by seller id: " + err.Error())
		return nil, err
	}
	return offers, nil
}

type OfferCreateParams struct {
	Uuid                 string
	PricePerMwh          float64
	InitialQuantityMwh   float64
	RemainingQuantityMwh float64
	Description          string
	PeriodStart          time.Time
	PeriodEnd            time.Time
	Status               uint8
	EnergyTypeId         uint
	SellerId             uint
	SubmarketId          uint
}

func (r *offerRepository) Create(params *OfferCreateParams) (*models.Offer, error) {
	offer := &models.Offer{
		Uuid:                 params.Uuid,
		PricePerMwh:          params.PricePerMwh,
		InitialQuantityMwh:   params.InitialQuantityMwh,
		RemainingQuantityMwh: params.RemainingQuantityMwh,
		Description:          params.Description,
		PeriodStart:          params.PeriodStart,
		PeriodEnd:            params.PeriodEnd,
		Status:               params.Status,
		EnergyTypeId:         params.EnergyTypeId,
		SellerId:             params.SellerId,
		SubmarketId:          params.SubmarketId,
	}

	if err := r.db.Create(offer).Error; err != nil {
		mlog.Log("Failed to create offer: " + err.Error())
		return nil, err
	}
	return offer, nil
}

func (r *offerRepository) Update(offer *models.Offer) error {
	return r.db.Save(offer).Error
}

func (r *offerRepository) Delete(uuid string) error {
	return r.db.Where("uuid = ?", uuid).Delete(&models.Offer{}).Error
}
