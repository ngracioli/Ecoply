package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/scopes"
	"ecoply/internal/domain/utils"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type PurchaseRepository interface {
	WithTransaction(tx *gorm.DB) PurchaseRepository

	Create(purchase *models.Purchase) error
	Update(purchase *models.Purchase) error
	FindByUuid(uuid string) (*models.Purchase, error)
	ListPurchases(buyerId uint64, request *requests.ListPurchase) (*utils.PaginationWrapper[*models.Purchase], error)
	ListSold(sellerId uint64, request *requests.ListSold) (*utils.PaginationWrapper[*models.Purchase], error)
}

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{
		db: db,
	}
}

func (r *purchaseRepository) WithTransaction(tx *gorm.DB) PurchaseRepository {
	return NewPurchaseRepository(tx)
}

func (r *purchaseRepository) Create(purchase *models.Purchase) error {
	if err := r.db.Create(purchase).Error; err != nil {
		mlog.Log("Failed to create purchase: " + err.Error())
		return err
	}
	return nil
}

func (r *purchaseRepository) Delete(uuid string) error {
	if err := r.db.Where("uuid = ?", uuid).Delete(&models.Purchase{}).Error; err != nil {
		mlog.Log("Failed to delete purchase: " + err.Error())
		return err
	}
	return nil
}

func (r *purchaseRepository) FindByUuid(uuid string) (*models.Purchase, error) {
	var purchase models.Purchase
	if err := r.db.
		Preload("Buyer").
		Preload("Offer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, seller_id")
		}).
		Preload("Offer.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid")
		}).
		Where("uuid = ?", uuid).First(&purchase).Error; err != nil {
		mlog.Log("Failed to find purchase by uuid: " + err.Error())
		return nil, err
	}
	return &purchase, nil
}

func (r *purchaseRepository) Update(purchase *models.Purchase) error {
	if err := r.db.Save(purchase).Error; err != nil {
		mlog.Log("Failed to update purchase: " + err.Error())
		return err
	}
	return nil
}

func (r *purchaseRepository) ListPurchases(buyerId uint64, request *requests.ListPurchase) (*utils.PaginationWrapper[*models.Purchase], error) {
	var purchases []*models.Purchase

	result := r.db.Debug().
		Preload("Buyer").
		Preload("Offer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, seller_id")
		}).
		Preload("Offer.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, name")
		}).
		Where("purchases.buyer_id = ?", buyerId).
		Select("purchases.*, (purchases.price_per_mwh * purchases.quantity_mwh) AS purchase_value")

	if request.Status != "" {
		result = result.Where("purchases.status = ?", request.Status)
	}

	if request.PaymentMethod != "" {
		result = result.Where("purchases.payment_method = ?", request.PaymentMethod)
	}

	switch request.OrderPrice {
	case "asc":
		result = result.Order("purchase_value ASC")
	case "desc":
		result = result.Order("purchase_value DESC")
	}

	switch request.OrderQuantity {
	case "asc":
		result = result.Order("purchases.quantity_mwh ASC")
	case "desc":
		result = result.Order("purchases.quantity_mwh DESC")
	}

	result = result.Scopes(scopes.Paginate(r.db, request.Page, request.PageSize))

	if err := result.Find(&purchases).Error; err != nil {
		mlog.Log("Failed to list purchases: " + err.Error())
		return nil, err
	}

	var paginationWrapper = utils.NewPaginationWrapper(request.Page, request.PageSize, purchases)

	return paginationWrapper, nil
}

func (r *purchaseRepository) ListSold(sellerId uint64, request *requests.ListSold) (*utils.PaginationWrapper[*models.Purchase], error) {
	var purchases []*models.Purchase

	result := r.db.
		Preload("Buyer").
		Preload("Offer", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, seller_id").
				Where("seller_id = ?", sellerId)
		}).
		Preload("Offer.Seller", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, uuid, name")
		}).
		Select("purchases.*, (purchases.price_per_mwh * purchases.quantity_mwh) AS purchase_value")

	if request.Status != "" {
		result = result.Where("purchases.status = ?", request.Status)
	}

	if request.PaymentMethod != "" {
		result = result.Where("purchases.payment_method = ?", request.PaymentMethod)
	}

	switch request.OrderPrice {
	case "asc":
		result = result.Order("purchase_value ASC")
	case "desc":
		result = result.Order("purchase_value DESC")
	}

	switch request.OrderQuantity {
	case "asc":
		result = result.Order("purchases.quantity_mwh ASC")
	case "desc":
		result = result.Order("purchases.quantity_mwh DESC")
	}

	result = result.Scopes(scopes.Paginate(r.db, request.Page, request.PageSize))

	if err := result.Find(&purchases).Error; err != nil {
		mlog.Log("Failed to list purchases: " + err.Error())
		return nil, err
	}

	var paginationWrapper = utils.NewPaginationWrapper(request.Page, request.PageSize, purchases)

	return paginationWrapper, nil
}
