package repository

import (
	"ecoply/internal/domain/models"
	"time"

	"gorm.io/gorm"
)

type OfferRepository interface {
	GetOfferByUuid(uuid string) (*models.Offer, error)
	GetUserOffers(userUuid string) ([]*models.Offer, error)
	CreateOffer(user *models.User, params OfferCreateParams) error
	UpdateOffer(offer *models.Offer) error
	DeleteOffer(id string) error
}

type offerRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) OfferRepository {
	return &offerRepository{db: db}
}

func (r *offerRepository) GetOfferByUuid(uuid string) (*models.Offer, error) {
	var offer models.Offer
	if err := r.db.Where("uuid = ?", uuid).First(&offer).Error; err != nil {
		return nil, err
	}
	return &offer, nil
}

func (r *offerRepository) GetUserOffers(userID string) ([]*models.Offer, error) {
	var offers []*models.Offer
	if err := r.db.Where("user_id = ?", userID).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}

type OfferCreateParams struct {
	Uuid               string
	PricePerMwh        float64
	InitialQuantityMwh float64
	PeriodStart        time.Time
	PeriodEnd          time.Time
	Description        string
	EnergyTypeId       uint
}

func (r *offerRepository) CreateOffer(user *models.User, params OfferCreateParams) error {
	offer := &models.Offer{
		Uuid:                 params.Uuid,
		SellerAgentId:        user.ID,
		SubmarketId:          user.Agent.SubmarketId,
		PricePerMwh:          params.PricePerMwh,
		InitialQuantityMwh:   params.InitialQuantityMwh,
		RemainingQuantityMwh: params.InitialQuantityMwh,
		PeriodStart:          params.PeriodStart,
		PeriodEnd:            params.PeriodEnd,
		Description:          params.Description,
		Status:               models.OfferStatusOpen,
		EnergyTypeId:         params.EnergyTypeId,
	}

	return r.db.Create(offer).Error
}

func (r *offerRepository) UpdateOffer(offer *models.Offer) error {
	return r.db.Save(offer).Error
}

func (r *offerRepository) DeleteOffer(id string) error {
	return r.db.Delete(&models.Offer{}, id).Error
}
