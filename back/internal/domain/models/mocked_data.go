package models

type MockedData struct {
	ID       uint64 `gorm:"primaryKey"`
	Inserted uint8  `gorm:"type:int;not null"`
}
