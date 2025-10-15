package models

const (
	EnergyTypeSolar      = "solar"
	EnergyTypeEolic      = "eolic"
	EnergyTypeHydro      = "hydroelectric"
	EnergyTypeGeothermal = "geothermal"
)

type EnergyType struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(50);not null"`
}
