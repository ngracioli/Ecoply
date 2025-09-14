package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"

	"gorm.io/gorm"
)

type UserTypeRepository interface {
	FindById(id uint) (*models.UserType, error)
	FindByName(name string) (*models.UserType, error)
}

type userTypeRepository struct {
	db *gorm.DB
}

func NewUserTypeRepository(db *gorm.DB) UserTypeRepository {
	return &userTypeRepository{db: db}
}

func (e *userTypeRepository) FindById(id uint) (*models.UserType, error) {
	var userType models.UserType
	err := e.db.First(&userType, id).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			mlog.Log("Failed to find user type by ID: " + err.Error())
		}
		return nil, err
	}

	return &userType, nil
}

func (e *userTypeRepository) FindByName(name string) (*models.UserType, error) {
	var userType models.UserType
	err := e.db.Where("type = ?", name).First(&userType).Error

	if err != nil {
		if err != gorm.ErrRecordNotFound {
			mlog.Log("Failed to find user type by name: " + err.Error())
		}
		return nil, err
	}

	return &userType, nil
}
