package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus string
type UserRole string

const (
	Active     UserStatus = "active"
	Inactive   UserStatus = "inactive"
	Unverified UserStatus = "unverified"

	Viewer      UserRole = "viewer"
	ReactViewer UserRole = "react_viewer"
	Admin       UserRole = "admin"
)

var (
	UserStatuses = []UserStatus{Active, Inactive, Unverified}
	UserRoles    = []UserRole{Viewer, ReactViewer, Admin}
)

type User struct {
	ID        uuid.UUID  `gorm:"column:id"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	Username     *string `gorm:"column:username"`
	FirstName    *string `gorm:"column:first_name"`
	LastName     *string `gorm:"column:last_name"`
	IsAnonymous  bool    `gorm:"column:is_anonymous"`
	PasswordHash string  `gorm:"column:password_hash"`

	Status UserStatus `gorm:"column:status"`
	Role   UserRole   `gorm:"column:role"`
}

func (User) TableName() string {
	return "users"
}
