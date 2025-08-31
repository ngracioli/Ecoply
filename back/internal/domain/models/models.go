package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UUID       string `gorm:"type:uuid;uniqueIndex;not null"`
	Name       string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:text;not null;unique"`
	Password   string `gorm:"type:varchar(255);not null"`
	CpfCnpj    string `gorm:"type:varchar(14);not null;unique"`
	UserTypeId uint   `gorm:"references:ID;not null"`
	UserType   UserType
	AddressId  uint `gorm:"references:ID;not null"`
	Address    Address
}

type UserType struct {
	ID   uint   `gorm:"primarykey"`
	Type string `gorm:"varchar(20);not null;unique"`
}

type Address struct {
	ID      uint   `gorm:"primarykey"`
	Cep     string `gorm:"varchar(8);not null"`
	State   string `gorm:"varchar(255);not null"`
	City    string `gorm:"varchar(255);not null"`
	Country string `gorm:"varchar(255);not null"`
}
