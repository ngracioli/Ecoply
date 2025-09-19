package models

type Address struct {
	ID         uint   `gorm:"primarykey"`
	Cep        string `gorm:"varchar(8);not null;index"`
	Complement string `gorm:"varchar(255)"`
	Number     string `gorm:"varchar(20);not null"`

	StreetId uint          `gorm:"references:ID;not null"`
	Street   AddressStreet `gorm:"foreignKey:StreetId"`
}

type AddressStreet struct {
	ID     uint   `gorm:"primarykey"`
	Street string `gorm:"varchar(255);not null;uniqueIndex"`

	NeighborhoodId uint                `gorm:"references:ID;not null"`
	Neighborhood   AddressNeighborhood `gorm:"foreignKey:NeighborhoodId"`
}

type AddressNeighborhood struct {
	ID           uint   `gorm:"primarykey"`
	Neighborhood string `gorm:"varchar(255);not null;uniqueIndex"`

	CityId uint        `gorm:"references:ID;not null"`
	City   AddressCity `gorm:"foreignKey:CityId"`
}

type AddressCity struct {
	ID   uint   `gorm:"primarykey"`
	City string `gorm:"varchar(255);not null;uniqueIndex"`

	StateId uint         `gorm:"references:ID;not null"`
	State   AddressState `gorm:"foreignKey:StateId"`
}

type AddressState struct {
	ID       uint   `gorm:"primarykey"`
	State    string `gorm:"varchar(255);not null;uniqueIndex"`
	Initials string `gorm:"varchar(2);not null;uniqueIndex"`
}
