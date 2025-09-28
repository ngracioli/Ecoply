package models

const (
	EnergyTypeSolar      = "solar"
	EnergyTypeWind       = "wind"
	EnergyTypeHydro      = "hydroelectric"
	EnergyTypeGeothermal = "geothermal"
)

type EnergyType struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(50);not null"`
}
