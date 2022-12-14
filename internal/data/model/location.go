package model

import "github.com/google/uuid"

type Location struct {
	ID          uuid.UUID `gorm:"column:id;default:uuid_generate_v4()"`
	CountryCode *string   `gorm:"column:country_code"`
	Region      *string   `gorm:"column:region"`
	City        *string   `gorm:"column:city"`
	Street      *string   `gorm:"column:street"`
}

func (Location) TableName() string {
	return "locations"
}
