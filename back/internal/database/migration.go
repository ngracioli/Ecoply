package database

import (
	"ecoply/internal/domain/models"

	"gorm.io/gorm"
)

type Migration struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(255);not null;uniqueIndex"`
}

func Migrate(con *gorm.DB) {
	con.Migrator().AutoMigrate(
		&models.UserType{},
		&models.Submarket{},
		&models.User{},
		&models.Agent{},

		&models.Address{},
		&models.AddressStreet{},
		&models.AddressNeighborhood{},
		&models.AddressCity{},
		&models.AddressState{},

		&models.EnergyType{},
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
