package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/scopes"
	"ecoply/internal/domain/utils"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type OfferRepository interface {
	WithTransaction(tx *gorm.DB) OfferRepository

	GetByUuid(uuid string) (*models.Offer, error)
	GetById(id uint) (*models.Offer, error)
	GetBySellerId(userId uint) ([]*models.Offer, error)
	Create(*models.Offer) (*models.Offer, error)
	List(request *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*models.Offer], error)
	Purchases(offerUuid string, request *requests.ListPurchasesFromOffer) (*utils.PaginationWrapper[*models.Purchase], error)
	Update(offer *models.Offer) error
	Delete(uuid string) error
	UpdateExpiredOffers() error
}

type offerRepository struct {
	db *gorm.DB
}

func NewOfferRepository(db *gorm.DB) OfferRepository {
	return &offerRepository{db: db}
}

func (r *offerRepository) WithTransaction(tx *gorm.DB) OfferRepository {
	return NewOfferRepository(tx)
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

func (r *offerRepository) Create(offer *models.Offer) (*models.Offer, error) {
	if err := r.db.Create(offer).Error; err != nil {
		mlog.Log("Failed to create offer: " + err.Error())
		return nil, err
	}
	return offer, nil
}

func (r *offerRepository) Update(offer *models.Offer) error {
	var err = r.db.Save(offer).Error
	if err != nil {
		mlog.Log("Failed to update offer: " + err.Error())
		return err
	}

	return nil
}

func (r *offerRepository) Delete(uuid string) error {
	var err = r.db.Where("uuid = ?", uuid).Delete(&models.Offer{}).Error
	if err != nil {
		mlog.Log("Failed to delete offfer: " + err.Error())
		return err
	}

	return nil
}

func (r *offerRepository) List(request *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*models.Offer], error) {
	var offers []*models.Offer

	result := r.db.
		Preload("Submarket").
		Preload("EnergyType").
		Preload("Seller").
		InnerJoins("Submarket").
		InnerJoins("EnergyType").
		Where("seller_id != ?", user.ID).
		Where("status NOT IN (?)", []string{models.OfferStatusExpired, models.OfferStatusFulfilled})

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

	var paginationWrapper = utils.NewPaginationWrapper(request.Page, request.PageSize, offers)

	return paginationWrapper, nil
}

func (r *offerRepository) UpdateExpiredOffers() error {
	return r.db.Model(&models.Offer{}).
		Where("period_end < ? AND status IN (?)", utils.NowInLocal(), []string{
			models.OfferStatusFresh,
			models.OfferStatusOpen,
		}).
		Update("status", models.OfferStatusExpired).Error
}

func (r *offerRepository) GetById(id uint) (*models.Offer, error) {
	var offer models.Offer
	if err := r.db.Preload("Submarket").
		Preload("EnergyType").
		Preload("Seller").
		First(&offer, id).Error; err != nil {
		return nil, err
	}
	return &offer, nil
}

func (r *offerRepository) Purchases(offerUuid string, request *requests.ListPurchasesFromOffer) (*utils.PaginationWrapper[*models.Purchase], error) {
	var purchases []*models.Purchase

	offer, err := r.GetByUuid(offerUuid)
	if err != nil {
		return nil, err
	}

	result := r.db.Debug().
		Preload("Buyer").
		Preload("Offer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, seller_id")
		}).
		Preload("Offer.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid")
		}).
		Where("purchases.offer_id = ?", offer.ID).
		Where("purchases.status IN (?)", []string{models.PurchaseStatusCompleted})

	if request.StartDate != "" {
		result = result.Where("purchases.created_at >= ?", request.StartDate)
	}

	if request.EndDate != "" {
		result = result.Where("purchases.created_at <= ?", request.EndDate)
	}

	if request.AscendingPrice {
		result = result.Order("purchases.price_per_mwh ASC")
	} else {
		result = result.Order("purchases.price_per_mwh DESC")
	}

	if request.AscendingQuantity {
		result = result.Order("purchases.quantity_mwh ASC")
	}

	result = result.Scopes(scopes.Paginate(r.db, request.Page, request.PageSize))

	if err := result.Find(&purchases).Error; err != nil {
		mlog.Log("Failed to list purchases from offer: " + err.Error())
		return nil, err
	}

	var paginationWrapper = utils.NewPaginationWrapper(request.Page, request.PageSize, purchases)

	return paginationWrapper, nil
}
