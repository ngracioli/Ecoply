package models

type Address struct {
	ID      uint   `gorm:"primarykey"`
	Cep     string `gorm:"varchar(8);not null"`
	State   string `gorm:"varchar(255);not null"`
	City    string `gorm:"varchar(255);not null"`
	Country string `gorm:"varchar(255);not null"`
}
