package models

type EnergyType struct {
	ID   uint   `gorm:"primaryKey;"`
	Name string `gorm:"type:varchar(50);primaryKey"`
}
