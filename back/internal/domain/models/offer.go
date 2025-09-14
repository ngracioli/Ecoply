package models

import "gorm.io/gorm"

const (
	OfferTypeSell     = "sell"
	OfferTypePurchase = "buy"

	OfferStatusPending    = "open"
	OfferStatusFullFilled = "fullfilled"
	OfferStatusExpired    = "expired"

	OfferEnergyTypeConventional = "conventional"
	OfferEnergyTypeIncentivized = "incentivized"
)

type Offer struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	PricePerMwh float64 `gorm:"type:decimal(10,2);not null"`
	QuantityMwh int     `gorm:"type:decimal(10,3)not null"`

	PeriodStart string `gorm:"type:date;not null"`
	PeriodEnd   string `gorm:"type:date;not null"`

	Status string `gorm:"type:varchar(50);not null"`

	SubmarketId uint `gorm:"references:ID;not null"`
	Submarket   Submarket

	AgentId uint `gorm:"references:ID;not null"`
	Agent   Agent
}
