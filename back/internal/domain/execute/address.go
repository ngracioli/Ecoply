package execute

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"net/http"

	"gorm.io/gorm"
)

type AddressExecute interface {
	Create(params AddressCreateParams) (*models.Address, *merr.ResponseError)
	FindById(id uint) (*models.Address, *merr.ResponseError)
}

type addressExecute struct {
	db *gorm.DB
}

func NewAddressExecute(db *gorm.DB) AddressExecute {
	return &addressExecute{db: db}
}

type AddressCreateParams struct {
	Cep     string
	State   string
	City    string
	Country string
}

func (a *addressExecute) Create(params AddressCreateParams) (*models.Address, *merr.ResponseError) {
	address := &models.Address{
		Cep:     params.Cep,
		State:   params.State,
		City:    params.City,
		Country: params.Country,
	}

	if err := a.db.Create(address).Error; err != nil {
		mlog.Log("Failed to create address: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return address, nil
}

func (a *addressExecute) FindById(id uint) (*models.Address, *merr.ResponseError) {
	var address models.Address
	err := a.db.First(&address, id).Error

	if err == gorm.ErrRecordNotFound {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find address by ID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &address, nil
}
