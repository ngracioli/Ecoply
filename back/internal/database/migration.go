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
}
