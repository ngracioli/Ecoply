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

	insertUserTypes(con)
	insertSubmarkets(con)
	insertEnergyTypes(con)
}

func insertUserTypes(con *gorm.DB) {
	var count int64
	con.Model(&models.UserType{}).Count(&count)

	if count == 0 {
		con.Create(&models.UserType{Type: models.UserTypeSupplier})
		con.Create(&models.UserType{Type: models.UserTypeBuyer})
	}
}

func insertSubmarkets(con *gorm.DB) {
	var count int64
	con.Model(&models.Submarket{}).Count(&count)

	if count == 0 {
		con.Create(&models.Submarket{Name: models.SubmarketSECO})
		con.Create(&models.Submarket{Name: models.SubmarketS})
		con.Create(&models.Submarket{Name: models.SubmarketNE})
		con.Create(&models.Submarket{Name: models.SubmarketN})
	}
}

func insertEnergyTypes(con *gorm.DB) {
	var count int64
	con.Model(&models.EnergyType{}).Count(&count)

	if count == 0 {
		con.Create(&models.EnergyType{Type: models.EnergyTypeSolar})
		con.Create(&models.EnergyType{Type: models.EnergyTypeWind})
		con.Create(&models.EnergyType{Type: models.EnergyTypeHydro})
		con.Create(&models.EnergyType{Type: models.EnergyTypeGeothermal})
	}
}
