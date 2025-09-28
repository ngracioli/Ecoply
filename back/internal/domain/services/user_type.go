package services

import (
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"

	"gorm.io/gorm"
)

type UserTypeService interface {
	UserIsSupplier(user *models.User) bool
}

type userTypeService struct {
	userTypeRepo repository.UserTypeRepository
}

func NewUserTypeService(db *gorm.DB) UserTypeService {
	return &userTypeService{
		userTypeRepo: repository.NewUserTypeRepository(db),
	}
}

func (s *userTypeService) UserIsSupplier(user *models.User) bool {
	userType, err := s.userTypeRepo.FindById(user.UserTypeId)
	if err != nil {
		return false
	}

	return userType.Type == models.UserTypeSupplier
}
