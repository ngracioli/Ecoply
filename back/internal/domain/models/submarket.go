package models

const (
	SubmarketSECO = "SE_CO" // Southeast and Center-West
	SubmarketS    = "S"     // South
	SubmarketNE   = "NE"    // Northeast
	SubmarketN    = "N"     // North
)

type Submarket struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(10);not null;unique"`
}
