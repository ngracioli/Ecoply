package execute

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UserExecute interface {
	Create(params UserCreateParams) (*models.User, *merr.ResponseError)
	CreateWithAddressAndType(params UserCreateWithAddressAndTypeParams) (*models.User, *merr.ResponseError)
	FindByEmail(email string) (*models.User, *merr.ResponseError)
	FindById(id uint) (*models.User, *merr.ResponseError)
	FindUserByCredentials(email string, password string) (*models.User, *merr.ResponseError)
	FindByCpfCnpj(cpfCnpj string) (*models.User, *merr.ResponseError)
	PreloadUserType(user *models.User) *merr.ResponseError
	PreloadAddress(user *models.User) *merr.ResponseError
}

type userExecute struct {
	db *gorm.DB
}

func NewUserExecute(db *gorm.DB) UserExecute {
	return &userExecute{db: db}
}

type UserCreateParams struct {
	Name       string
	Email      string
	Password   string
	CpfCnpj    string
	UserTypeId uint
	AddressId  uint
}

func (e *userExecute) Create(params UserCreateParams) (*models.User, *merr.ResponseError) {
	result, err := e.FindByEmail(params.Email)

	if err != nil && err.StatusCode != http.StatusNotFound {
		return nil, err
	}

	if result != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserEmailAlreadyExists)
	}

	user := &models.User{
		UUID:       services.NewUuidV7String(),
		Name:       params.Name,
		Email:      params.Email,
		Password:   params.Password,
		CpfCnpj:    params.CpfCnpj,
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
	CpfCnpj  string
	UserType string
	Address  AddressCreateParams
}

func (e *userExecute) CreateWithAddressAndType(params UserCreateWithAddressAndTypeParams) (*models.User, *merr.ResponseError) {
	var user *models.User
	var responseError *merr.ResponseError

	err := e.db.Transaction(func(tx *gorm.DB) error {
		var addressExec AddressExecute = NewAddressExecute(tx)
		address, addressErr := addressExec.Create(params.Address)

		if addressErr != nil {
			responseError = addressErr
			return addressErr.Error
		}

		var userTypeExec UserTypeExecute = NewUserTypeExecute(e.db)
		userType, userTypeErr := userTypeExec.FindByName(params.UserType)

		if userTypeErr != nil {
			responseError = userTypeErr
			return userTypeErr.Error
		}

		var userExec UserExecute = NewUserExecute(tx)
		var userCreateParams = UserCreateParams{
			Name:       params.Name,
			Email:      params.Email,
			Password:   params.Password,
			CpfCnpj:    params.CpfCnpj,
			UserTypeId: userType.ID,
			AddressId:  address.ID,
		}

		createdUser, userErr := userExec.Create(userCreateParams)
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

func (e *userExecute) FindByEmail(email string) (*models.User, *merr.ResponseError) {
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

func (e *userExecute) FindById(id uint) (*models.User, *merr.ResponseError) {
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

func (e *userExecute) FindByCpfCnpj(cpfCnpj string) (*models.User, *merr.ResponseError) {
	var user models.User
	err := e.db.Where("cpf_cnpj = ?", cpfCnpj).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrNotFound)
	}

	if err != nil {
		mlog.Log("Failed to find user by CPF/CNPJ: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &user, nil
}

func (e *userExecute) FindUserByCredentials(email string, password string) (*models.User, *merr.ResponseError) {
	user, err := e.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrIncorrectPassword)
	}

	return user, nil
}

func (e *userExecute) PreloadUserType(user *models.User) *merr.ResponseError {
	if err := e.db.Preload("UserType").First(user).Error; err != nil {
		mlog.Log("Failed to preload user type: " + err.Error())
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return nil
}

func (e *userExecute) PreloadAddress(user *models.User) *merr.ResponseError {
	if err := e.db.Preload("Address").First(user).Error; err != nil {
		mlog.Log("Failed to preload address: " + err.Error())
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return nil
}
