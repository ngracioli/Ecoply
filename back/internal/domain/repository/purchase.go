package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type PurchaseRepository interface {
	WithTransaction(tx *gorm.DB) PurchaseRepository

	Create(purchase *models.Purchase) error
	Delete(uuid string) error
	FindByUuid(uuid string) (*models.Purchase, error)
	List() ([]*models.Purchase, error)
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
	if err := r.db.Preload("Buyer").
		Preload("Offer").
		Where("uuid = ?", uuid).First(&purchase).Error; err != nil {
		mlog.Log("Failed to find purchase by uuid: " + err.Error())
		return nil, err
	}
	return &purchase, nil
}

func (r *purchaseRepository) List() ([]*models.Purchase, error) {
	return nil, nil
}
