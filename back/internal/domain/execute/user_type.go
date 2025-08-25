package execute

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"net/http"

	"gorm.io/gorm"
)

type UserTypeExecute interface {
	FindById(id uint) (*models.UserType, *merr.ResponseError)
	FindByName(name string) (*models.UserType, *merr.ResponseError)
}

type userTypeExecute struct {
	db *gorm.DB
}

func NewUserTypeExecute(db *gorm.DB) UserTypeExecute {
	return &userTypeExecute{db: db}
}

func (e *userTypeExecute) FindById(id uint) (*models.UserType, *merr.ResponseError) {
	var userType models.UserType
	err := e.db.First(&userType, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user type by ID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &userType, nil
}

func (e *userTypeExecute) FindByName(name string) (*models.UserType, *merr.ResponseError) {
	var userType models.UserType
	err := e.db.Where("type = ?", name).First(&userType).Error

	if err == gorm.ErrRecordNotFound {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user type by name: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &userType, nil
}
