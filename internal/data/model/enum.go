package model

type Country struct {
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
}

func (Country) TableName() string {
	return "country"
}
