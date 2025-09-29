package models

import "gorm.io/gorm"

const (
	OfferStatusOpen       uint8 = 1
	OfferStatusFullFilled uint8 = 2
	OfferStatusExpired    uint8 = 3
)

type Offer struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	PricePerMwh          float64 `gorm:"type:decimal(10,2);not null"`
	InitialQuantityMwh   float64 `gorm:"type:decimal(10,3);not null"`
	RemainingQuantityMwh float64 `gorm:"type:decimal(10,3);not null"`

	PeriodStart string `gorm:"type:date;not null"`
	PeriodEnd   string `gorm:"type:date;not null"`

	Status string `gorm:"type:int;not null"`

	EnergyTypeId uint       `gorm:"references:ID;not null"`
	EnergyType   EnergyType `gorm:"foreignKey:EnergyTypeId"`

	SubmarketId uint      `gorm:"references:ID;not null"`
	Submarket   Submarket `gorm:"foreignKey:SubmarketId"`

	SellerAgentId uint  `gorm:"references:ID;not null"`
	SellerAgent   Agent `gorm:"foreignKey:SellerAgentId"`

	Purchase []Purchase `gorm:"foreignKey:OfferId"`
}
