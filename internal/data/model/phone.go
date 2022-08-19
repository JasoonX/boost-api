package model

import (
	"github.com/google/uuid"
)

type Phone struct {
	ID               uuid.UUID `gorm:"column:id"`
	SubscriberNumber string    `gorm:"column:subscriber_number"`
	Extension        *string   `gorm:"column:extension"`
	CountryCode      *string   `gorm:"column:country_code"`
	IsVerified       bool      `gorm:"column:is_verified"`
	IsPrimary        bool      `gorm:"column:is_primary"`
	UserID           uuid.UUID `gorm:"column:user_id"`
}

func (Phone) TableName() string {
	return "phones"
}

func (p Phone) GetID() string {
	return p.ID.String()
}
