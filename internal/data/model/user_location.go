package model

import (
	"time"

	"github.com/google/uuid"
)

type UserLocation struct {
	ID         uuid.UUID `gorm:"column:id"`
	UserID     uuid.UUID `gorm:"column:user_id"`
	LocationID uuid.UUID `gorm:"column:location_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (UserLocation) TableName() string {
	return "user_location"
}
