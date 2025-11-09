package models

import "gorm.io/gorm"

const (
	PurchaseStatusCompleted = "completed"
)

type Purchase struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	QuantityMwh float64 `gorm:"type:decimal(10,3) not null"`
	PricePerMwh float64 `gorm:"type:decimal(10,2) not null"`

	Status string `gorm:"type:varchar(50);not null"`

	BuyerId uint `gorm:"references:ID;not null"`
	Buyer   User `gorm:"foreignKey:BuyerId"`

	OfferId uint  `gorm:"references:ID;not null"`
	Offer   Offer `gorm:"foreignKey:OfferId"`
}

func (p *Purchase) IsCompleted() bool {
	return p.Status == PurchaseStatusCompleted
}
