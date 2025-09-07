package models

import (
	"gorm.io/gorm"
)

type User struct {
	Uuid       string `gorm:"type:uuid;uniqueIndex;not null"`
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:text;not null;unique"`
	Password   string `gorm:"type:varchar(255);not null"`
	Cnpj       string `gorm:"type:varchar(14);not null;unique"`
	UserTypeId uint   `gorm:"references:ID;not null"`
	UserType   UserType
	AddressId  uint `gorm:"references:ID;not null"`
	Address    Address

	gorm.Model
}
