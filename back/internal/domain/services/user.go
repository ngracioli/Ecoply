package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"net/http"

	"gorm.io/gorm"
)

type UserService interface {
	FindByUuid(uuid string) (*models.User, *merr.ResponseError)
}

type userService struct {
	userRepo repository.UserRepository
	db       *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		userRepo: repository.NewUserRepository(db),
		db:       db,
	}
}

func (s *userService) FindByUuid(uuid string) (*models.User, *merr.ResponseError) {
	user, err := s.userRepo.FindByUuid(uuid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrUserNotFound)
		}
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}
	return user, nil
}
