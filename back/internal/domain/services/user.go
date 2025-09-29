package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"net/http"

	"gorm.io/gorm"
)

type UserService interface {
	FindByUuid(uuid string) (*models.User, *merr.ResponseError)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) FindByUuid(uuid string) (*models.User, *merr.ResponseError) {
	var user models.User
	if err := s.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, merr.NewResponseError(http.StatusNotFound, ErrUserNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return &user, nil
}
