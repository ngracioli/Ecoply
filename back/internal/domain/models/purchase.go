package models

import "gorm.io/gorm"

const (
	PurchaseStatusCompleted = "completed"
)

type Purchase struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	QuantityMwh int     `gorm:"type:decimal(10,3) not null"`
	PricePerMwh float64 `gorm:"type:decimal(10,2) not null"`

	Status string `gorm:"type:varchar(50);not null"`

	EnergyTypeId uint       `gorm:"references:ID;not null"`
	EnergyType   EnergyType `gorm:"foreignKey:EnergyTypeId"`

	BuyerAgentId uint  `gorm:"references:ID;not null"`
	BuyerAgent   Agent `gorm:"foreignKey:BuyerAgentId"`

	OfferId uint  `gorm:"references:ID;not null"`
	Offer   Offer `gorm:"foreignKey:OfferId"`
}
