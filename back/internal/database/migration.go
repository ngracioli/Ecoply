package database

import (
	"ecoply/internal/domain/models"

	"gorm.io/gorm"
)

func Migrate(con *gorm.DB) {
	con.Migrator().AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Agent{},
		&models.Submarket{},
		&models.Offer{},
		&models.Purchase{},
	)

	initUserTypes()
	initSubmarkets()
}

func initUserTypes() {
	var count int64
	Con.Model(&models.UserType{}).Count(&count)

	if count == 0 {
		Con.Create(&models.UserType{Type: models.UserTypeSupplier})
		Con.Create(&models.UserType{Type: models.UserTypeBuyer})
	}
}

func initSubmarkets() {
	var count int64
	Con.Model(&models.Submarket{}).Count(&count)

	if count == 0 {
		Con.Create(&models.Submarket{Name: models.SubmarketSECO})
		Con.Create(&models.Submarket{Name: models.SubmarketS})
		Con.Create(&models.Submarket{Name: models.SubmarketNE})
		Con.Create(&models.Submarket{Name: models.SubmarketN})
	}
}
