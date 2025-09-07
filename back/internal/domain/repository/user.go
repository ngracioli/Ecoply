package repository

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(params UserCreateParams) (*models.User, *merr.ResponseError)
	CreateWithAddressAndType(params UserCreateWithAddressAndTypeParams) (*models.User, *merr.ResponseError)
	FindByEmail(email string) (*models.User, *merr.ResponseError)
	FindById(id uint) (*models.User, *merr.ResponseError)
	FindByUuid(uuid string) (*models.User, *merr.ResponseError)
	FindUserByCredentials(email string, password string) (*models.User, *merr.ResponseError)
	FindByCnpj(cnpj string) (*models.User, *merr.ResponseError)
	PreloadUserType(user *models.User) *merr.ResponseError
	PreloadAddress(user *models.User) *merr.ResponseError
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

type UserCreateParams struct {
	Name       string
	Email      string
	Password   string
	Cnpj       string
	UserTypeId uint
	AddressId  uint
}

func (e *userRepository) Create(params UserCreateParams) (*models.User, *merr.ResponseError) {
	result, err := e.FindByEmail(params.Email)

	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	if result != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserEmailAlreadyExists)
	}

	result, err = e.FindByCnpj(params.Cnpj)

	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	if result != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserCnpjAlreadyExists)
	}

	user := &models.User{
		Uuid:       services.NewUuidV7String(),
		Name:       params.Name,
		Email:      params.Email,
		Password:   params.Password,
		Cnpj:       params.Cnpj,
		UserTypeId: params.UserTypeId,
		AddressId:  params.AddressId,
	}

	if err := e.db.Create(user).Error; err != nil {
		mlog.Log("Failed to create user: " + err.Error())
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInternal)
	}

	return user, nil
}

type UserCreateWithAddressAndTypeParams struct {
	Name     string
	Email    string
	Password string
	Cnpj     string
	UserType string
	Address  AddressCreateParams
}

func (e *userRepository) CreateWithAddressAndType(params UserCreateWithAddressAndTypeParams) (*models.User, *merr.ResponseError) {
	var user *models.User
	var responseError *merr.ResponseError

	err := e.db.Transaction(func(tx *gorm.DB) error {
		var addressRepo AddressRepository = NewAddressRepository(tx)
		address, addressErr := addressRepo.Create(params.Address)

		if addressErr != nil {
			responseError = addressErr
			return addressErr.Error
		}

		var userTypeRepo UserTypeRepository = NewUserTypeRepository(tx)
		userType, userTypeErr := userTypeRepo.FindByName(params.UserType)

		if userTypeErr != nil {
			responseError = userTypeErr
			return userTypeErr.Error
		}

		var userRepo UserRepository = NewUserRepository(tx)
		var userCreateParams = UserCreateParams{
			Name:       params.Name,
			Email:      params.Email,
			Password:   params.Password,
			Cnpj:       params.Cnpj,
			UserTypeId: userType.ID,
			AddressId:  address.ID,
		}

		createdUser, userErr := userRepo.Create(userCreateParams)
		if userErr != nil {
			responseError = userErr
			return userErr.Error
		}

		user = createdUser

		return nil
	})

	if err != nil {
		if responseError != nil {
			return nil, responseError
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return user, nil
}

func (e *userRepository) FindByEmail(email string) (*models.User, *merr.ResponseError) {
	var user models.User
	err := e.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user by email: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &user, nil
}

func (e *userRepository) FindById(id uint) (*models.User, *merr.ResponseError) {
	var user models.User
	err := e.db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user by ID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &user, nil
}

func (e *userRepository) FindByCnpj(cnpj string) (*models.User, *merr.ResponseError) {
	var user models.User
	err := e.db.Where("cnpj = ?", cnpj).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user by CPF/CNPJ: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &user, nil
}

func (e *userRepository) FindUserByCredentials(email string, password string) (*models.User, *merr.ResponseError) {
	user, err := e.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrIncorrectPassword)
	}

	return user, nil
}

func (e *userRepository) FindByUuid(uuid string) (*models.User, *merr.ResponseError) {
	var user models.User
	err := e.db.Where("uuid = ?", uuid).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user by UUID: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &user, nil
}

func (e *userRepository) PreloadUserType(user *models.User) *merr.ResponseError {
	if err := e.db.Preload("UserType").First(user).Error; err != nil {
		mlog.Log("Failed to preload user type: " + err.Error())
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return nil
}

func (e *userRepository) PreloadAddress(user *models.User) *merr.ResponseError {
	if err := e.db.Preload("Address").First(user).Error; err != nil {
		mlog.Log("Failed to preload address: " + err.Error())
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return nil
}
