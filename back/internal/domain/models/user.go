package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Uuid     string `gorm:"type:uuid;uniqueIndex;not null"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:text;not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`

	UserTypeId uint     `gorm:"references:ID;not null"`
	UserType   UserType `gorm:"foreignKey:UserTypeId"`

	AgentId uint  `gorm:"references:ID;not null"`
	Agent   Agent `gorm:"foreignKey:AgentId"`
}
