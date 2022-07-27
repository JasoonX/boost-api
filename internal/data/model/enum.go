package model

type Country struct {
	CountryCode string `gorm:"column:country_code"`
	Name        string `gorm:"column:name"`
}

func (Country) TableName() string {
	return "country"
}
