package model

import (
	"github.com/google/uuid"
)

type Email struct {
	ID         uuid.UUID `gorm:"column:id"`
	Email      string    `gorm:"column:email"`
	IsVerified bool      `gorm:"column:is_verified"`
	IsPrimary  bool      `gorm:"column:is_primary"`
	UserID     uuid.UUID `gorm:"column:user_id"`
}

func (Email) TableName() string {
	return "emails"
}

func (e Email) GetID() string {
	return e.ID.String()
}
