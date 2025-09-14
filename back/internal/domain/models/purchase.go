package models

import "gorm.io/gorm"

const (
	PurchaseStatusCompleted = "completed"
)

type Purchase struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	QuantityMwh int     `gorm:"type:decimal(10,3) not null"`
	TotalPrice  float64 `gorm:"type:decimal(10,2) not null"`

	Status string `gorm:"type:varchar(50);not null"`

	BuyerAgentId uint `gorm:"references:ID;not null"`
	BuyerAgent   Agent

	OfferId uint `gorm:"references:ID;not null"`
	Offer   Offer
}
