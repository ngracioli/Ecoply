package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	OfferStatusFresh      uint8 = 1
	OfferStatusOpen       uint8 = 2
	OfferStatusFullFilled uint8 = 3
	OfferStatusExpired    uint8 = 4
)

type Offer struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	PricePerMwh          float64 `gorm:"type:decimal(10,2);not null"`
	InitialQuantityMwh   float64 `gorm:"type:decimal(10,3);not null"`
	RemainingQuantityMwh float64 `gorm:"type:decimal(10,3);not null"`

	Description string `gorm:"type:text;not null"`

	PeriodStart time.Time `gorm:"type:datetime;not null"`
	PeriodEnd   time.Time `gorm:"type:datetime;not null"`

	Status uint8 `gorm:"type:int;not null"`

	EnergyTypeId uint       `gorm:"references:ID;not null"`
	EnergyType   EnergyType `gorm:"foreignKey:EnergyTypeId"`

	SubmarketId uint      `gorm:"references:ID;not null"`
	Submarket   Submarket `gorm:"foreignKey:SubmarketId"`

	SellerId uint `gorm:"references:ID;not null"`
	Seller   User `gorm:"foreignKey:SellerId"`

	Purchase []Purchase `gorm:"foreignKey:OfferId"`
}
