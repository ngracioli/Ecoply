package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type SubmarketRepository interface {
	WithTransaction(tx *gorm.DB) SubmarketRepository

	FindByName(name string) (*models.Submarket, error)
}

type submarketRepository struct {
	db *gorm.DB
}

func (r *submarketRepository) WithTransaction(tx *gorm.DB) SubmarketRepository {
	return NewSubmarketRepository(tx)
}

func NewSubmarketRepository(db *gorm.DB) SubmarketRepository {
	return &submarketRepository{db: db}
}

func (r *submarketRepository) FindByName(name string) (*models.Submarket, error) {
	var submarket models.Submarket

	if err := r.db.Where("name = ?", name).First(&submarket).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			mlog.Log("Failed to find submarket by name: " + err.Error())
		}
		return nil, err
	}

	return &submarket, nil
}
