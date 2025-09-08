package repository

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(params AddressCreateParams) (*models.Address, *merr.ResponseError)
	FindById(id uint) (*models.Address, *merr.ResponseError)
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db: db}
}

type AddressCreateParams struct {
	Cep           string
	Complement    string
	Street        string
	Number        string
	Neighborhood  string
	City          string
	State         string
	StateInitials string
}

func (a *addressRepository) Create(params AddressCreateParams) (*models.Address, *merr.ResponseError) {
	var state models.AddressState
	var city models.AddressCity
	var neighborhood models.AddressNeighborhood
	var street models.AddressStreet
	var address *models.Address

	err := a.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.FirstOrCreate(&state, models.AddressState{
			State: params.State,
		}).Error; err != nil {
			return err
		}

		if err := tx.FirstOrCreate(&city, models.AddressCity{
			City:    params.City,
			StateId: state.ID,
		}).Error; err != nil {
			return err
		}

		if err := tx.FirstOrCreate(&neighborhood, models.AddressNeighborhood{
			Neighborhood: params.Neighborhood,
			CityId:       city.ID,
		}).Error; err != nil {
			return err
		}

		if err := tx.FirstOrCreate(&street, models.AddressStreet{
			Street:         params.Street,
			NeighborhoodId: neighborhood.ID,
		}).Error; err != nil {
			return err
		}

		createdAddress := &models.Address{
			Cep:        params.Cep,
			Complement: params.Complement,
			Number:     params.Number,
			StreetId:   street.ID,
		}

		if err := tx.Create(createdAddress).Error; err != nil {
			return err
		}

		address = createdAddress

		return nil
	})

	if err != nil {
		mlog.Log("Failed to create address: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return address, nil
}

func (a *addressRepository) FindById(id uint) (*models.Address, *merr.ResponseError) {
	var address models.Address
	err := a.db.First(&address, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find address by ID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &address, nil
}
