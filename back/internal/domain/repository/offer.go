package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/scopes"
	"ecoply/internal/domain/utils"
	"ecoply/internal/mlog"
	"time"

	"gorm.io/gorm"
)

type OfferRepository interface {
	GetByUuid(uuid string) (*models.Offer, error)
	GetBySellerId(userId uint) ([]*models.Offer, error)
	Create(params *OfferCreateParams) (*models.Offer, error)
	List(request *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*models.Offer], error)
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
	Status               string
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

func (r *offerRepository) List(request *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*models.Offer], error) {
	var offers []*models.Offer

	result := r.db.Joins("Submarket").
		Joins("EnergyType").
		Joins("Seller").
		Where("seller_id != ?", user.ID).
		Where("status NOT IN (?)", []string{models.OfferStatusExpired, models.OfferStatusFullFilled})

	if request.Submarket != "" {
		result = result.Where("\"Submarket\".name = ?", request.Submarket)
	}

	if request.EnergyType != "" {
		result = result.Where("\"EnergyType\".type = ?", request.EnergyType)
	}

	switch {
	case request.PeriodStart != "" && request.PeriodEnd != "":
		result = result.Where("period_start >= ? AND period_end <= ?", request.PeriodStart, request.PeriodEnd)
	case request.PeriodStart != "":
		result = result.Where("period_start >= ?", request.PeriodStart)
	case request.PeriodEnd != "":
		result = result.Where("period_end <= ?", request.PeriodEnd)
	}

	result = result.Scopes(scopes.Paginate(r.db, request.Page, request.PageSize))

	if err := result.Find(&offers).Error; err != nil {
		mlog.Log("Failed to list offers: " + err.Error())
		return nil, err
	}

	var paginationWrapper = utils.NewPaginationWrapper[*models.Offer](request.Page, request.PageSize, offers)

	return paginationWrapper, nil
}
