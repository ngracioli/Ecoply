package models

import (
	"ecoply/internal/domain/utils"
	"time"

	"gorm.io/gorm"
)

const (
	OfferStatusFresh     string = "fresh"
	OfferStatusOpen      string = "open"
	OfferStatusFulfilled string = "fulfilled"
	OfferStatusExpired   string = "expired"
)

type Offer struct {
	gorm.Model

	Uuid string `gorm:"type:uuid;uniqueIndex;not null"`

	PricePerMwh          float64 `gorm:"type:decimal(10,2);not null"`
	InitialQuantityMwh   float64 `gorm:"type:decimal(10,3);not null"`
	RemainingQuantityMwh float64 `gorm:"type:decimal(10,3);not null"`

	Description string `gorm:"type:text;not null"`

	PeriodStart time.Time `gorm:"type:date;not null"`
	PeriodEnd   time.Time `gorm:"type:date;not null"`

	Status string `gorm:"type:varchar(20);not null"`

	EnergyTypeId uint       `gorm:"references:ID;not null"`
	EnergyType   EnergyType `gorm:"foreignKey:EnergyTypeId"`

	SubmarketId uint      `gorm:"references:ID;not null"`
	Submarket   Submarket `gorm:"foreignKey:SubmarketId"`

	SellerId uint `gorm:"references:ID;not null"`
	Seller   User `gorm:"foreignKey:SellerId"`

	Purchase []Purchase `gorm:"foreignKey:OfferId"`
}

func (o *Offer) IsExpired() bool {
	var now time.Time = utils.NowInLocalZeroHour()
	var periodEnd time.Time = utils.TruncateDateToLocalZeroHour(o.PeriodEnd)

	return now.After(periodEnd)
}

func (o *Offer) IsFresh() bool {
	return o.Status == OfferStatusFresh
}

func (o *Offer) IsOpen() bool {
	return o.Status == OfferStatusOpen
}

func (o *Offer) IsFulfilled() bool {
	return o.RemainingQuantityMwh == 0
}
