package repository

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"net/http"

	"gorm.io/gorm"
)

type SubmarketRepository interface {
	FindByName(name string) (*models.Submarket, *merr.ResponseError)
}

type submarketRepository struct {
	db *gorm.DB
}

func NewSubmarketRepository(db *gorm.DB) SubmarketRepository {
	return &submarketRepository{db: db}
}

func (r *submarketRepository) FindByName(name string) (*models.Submarket, *merr.ResponseError) {
	var submarket models.Submarket

	if err := r.db.Where("name = ?", name).First(&submarket).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
		}
		mlog.Log("Failed to find submarket by name: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &submarket, nil
}
