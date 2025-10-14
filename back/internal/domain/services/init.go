package services

import (
	"gorm.io/gorm"
)

var (
	Auth     AuthService
	Offer    OfferService
	User     UserService
	UserType UserTypeService
)

func InitServices(db *gorm.DB) {
	Auth = NewAuthService(db)
	Offer = NewOfferService(db)
	User = NewUserService(db)
	UserType = NewUserTypeService(db)
}
