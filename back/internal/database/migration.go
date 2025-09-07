package database

import (
	"ecoply/internal/domain/models"

	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) {
	con.AutoMigrate(
		&models.User{},
		&models.UserType{},
		&models.Address{},
	)

	initUserTypes()
}

func initUserTypes() {
	var count int64
	Con.Model(&models.UserType{}).Count(&count)

	if count == 0 {
		Con.Create(&models.UserType{Type: "supplier"})
		Con.Create(&models.UserType{Type: "buyer"})
	}
}
