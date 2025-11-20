package models

import "gorm.io/gorm"

const (
	PurchaseStatusCompleted = "completed"
	PurchaseStatusWaiting   = "waiting"
	PurchaseStatusCanceled  = "canceled"

	PurchasePaymentPix    = "pix"
	PurchasePaymentCard   = "card"
	PurchasePaymentBillet = "billet"
)

type Purchase struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	QuantityMwh float64 `gorm:"type:decimal(10,3) not null"`
	PricePerMwh float64 `gorm:"type:decimal(10,2) not null"`

	Status string `gorm:"type:varchar(50);not null"`

	PaymentMethod string `gorm:"type:varchar(20);not null"`

	BuyerId uint `gorm:"references:ID;not null"`
	Buyer   User `gorm:"foreignKey:BuyerId"`

	OfferId uint  `gorm:"references:ID;not null"`
	Offer   Offer `gorm:"foreignKey:OfferId"`
}

func (p *Purchase) IsCompleted() bool {
	return p.Status == PurchaseStatusCompleted
}

func (p *Purchase) IsWaiting() bool {
	return p.Status == PurchaseStatusWaiting
}

func (p *Purchase) IsCancelled() bool {
	return p.Status == PurchaseStatusCanceled
}

func (p *Purchase) IsPix() bool {
	return p.PaymentMethod == PurchasePaymentPix
}

func (p *Purchase) IsCard() bool {
	return p.PaymentMethod == PurchasePaymentCard
}

func (p *Purchase) IsBillet() bool {
	return p.PaymentMethod == PurchasePaymentBillet
}

func (p *Purchase) IsOwner(user *User) bool {
	return user.ID == p.BuyerId
}
