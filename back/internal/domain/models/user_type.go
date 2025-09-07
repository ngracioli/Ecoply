package models

type UserType struct {
	ID   uint   `gorm:"primarykey"`
	Type string `gorm:"varchar(20);not null;unique"`
}
