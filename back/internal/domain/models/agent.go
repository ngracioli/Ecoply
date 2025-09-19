package models

import "gorm.io/gorm"

type Agent struct {
	gorm.Model

	Cnpj        string `gorm:"primaryKey,type:varchar(14);not null;unique"`
	CompanyName string `gorm:"type:varchar(255);not null"`
	CceeCode    string `gorm:"type:varchar(50);not null;uniqueIndex"`

	SubmarketId uint      `gorm:"references:ID;not null"`
	Submarket   Submarket `gorm:"foreignKey:SubmarketId"`

	AddressId uint    `gorm:"references:ID;not null"`
	Address   Address `gorm:"foreignKey:AddressId"`
}
