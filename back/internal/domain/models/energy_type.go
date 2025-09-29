package models

const (
	EnergyTypeSolar      = "Solar"
	EnergyTypeWind       = "Eólica"
	EnergyTypeHydro      = "Hidroelétrica"
	EnergyTypeGeothermal = "Geotérmica"
)

type EnergyType struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(50);unique;not null"`
}
