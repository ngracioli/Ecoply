package repository

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/mlog"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTransaction(tx *gorm.DB) UserRepository

	Create(params UserCreateParams) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindByUuid(uuid string) (*models.User, error)
	PreloadUserType(user *models.User) error
	PreloadAddress(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) WithTransaction(tx *gorm.DB) UserRepository {
	return NewUserRepository(tx)
}

type UserCreateParams struct {
	Uuid     string
	Name     string
	Email    string
	Password string
	UserType *models.UserType
	Agent    *models.Agent
}

func (e *userRepository) Create(params UserCreateParams) (*models.User, error) {
	user := &models.User{
		Uuid:       params.Uuid,
		Name:       params.Name,
		Email:      params.Email,
		Password:   params.Password,
		UserTypeId: params.UserType.ID,
		AgentId:    params.Agent.ID,
	}

	if err := e.db.Create(user).Error; err != nil {
		mlog.Log("Failed to create user: " + err.Error())
		return nil, err
	}

	return user, nil
}

func (e *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := e.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			mlog.Log("Failed to find user by email: " + err.Error())
		}
		return nil, err
	}

	return &user, nil
}

func (e *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	err := e.db.First(&user, id).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			mlog.Log("Failed to find user by ID: " + err.Error())
		}
		return nil, err
	}

	return &user, nil
}

func (e *userRepository) FindByUuid(uuid string) (*models.User, error) {
	var user models.User
	err := e.db.Where("uuid = ?", uuid).First(&user).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			mlog.Log("Failed to find user by UUID: " + err.Error())
		}
		return nil, err
	}

	return &user, nil
}

func (e *userRepository) PreloadUserType(user *models.User) error {
	if err := e.db.Preload("UserType").First(user).Error; err != nil {
		mlog.Log("Failed to preload user type: " + err.Error())
		return err
	}
	return nil
}

func (e *userRepository) PreloadAddress(user *models.User) error {
	if err := e.db.Preload("Address").First(user).Error; err != nil {
		mlog.Log("Failed to preload address: " + err.Error())
		return err
	}
	return nil
}
