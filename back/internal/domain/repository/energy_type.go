package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type EnergyTypeRepository interface {
	WithTransaction(tx *gorm.DB) EnergyTypeRepository

	GetByType(energy string) (*models.EnergyType, error)
}

type energyTypeRepository struct {
	db *gorm.DB
}

func NewEnergyRepository(db *gorm.DB) EnergyTypeRepository {
	return &energyTypeRepository{db: db}
}

func (r *energyTypeRepository) WithTransaction(tx *gorm.DB) EnergyTypeRepository {
	return NewEnergyRepository(tx)
}

func (r *energyTypeRepository) GetByType(energy string) (*models.EnergyType, error) {
	var energyType models.EnergyType

	var err = r.db.Where("type = ?", energy).Find(&energyType).Error
	if err != nil {
		mlog.Log("Failed to find energy type: " + err.Error())
		return nil, err
	}

	if energyType.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &energyType, nil
}
