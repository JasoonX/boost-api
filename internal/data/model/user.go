package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type UserStatus string
type UserRole string

const (
	UserStatusActive     UserStatus = "active"
	UserStatusInactive   UserStatus = "inactive"
	UserStatusUnverified UserStatus = "unverified"

	UserRoleViewer      UserRole = "viewer"
	UserRoleReactViewer UserRole = "react_viewer"
	UserRoleSuperAdmin  UserRole = "super_admin"
)

func (r UserRole) Equals(other UserRole) bool {
	return r == other
}

var (
	UserStatuses = []UserStatus{UserStatusActive, UserStatusInactive, UserStatusUnverified}
	UserRoles    = []UserRole{UserRoleViewer, UserRoleReactViewer, UserRoleSuperAdmin}
)

type User struct {
	ID        uuid.UUID  `gorm:"column:id;default:uuid_generate_v4()"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	Username     *string `gorm:"column:username"`
	FirstName    *string `gorm:"column:first_name"`
	LastName     *string `gorm:"column:last_name"`
	PasswordHash string  `gorm:"column:password_hash"`

	Status UserStatus     `gorm:"column:status"`
	Role   pq.StringArray `gorm:"column:role;type:text[]"`

	Emails []Email `gorm:"foreignkey:user_id"`
	Phones []Phone `gorm:"foreignkey:user_id"`
}

func (u User) IsRole(role UserRole) bool {
	for _, r := range u.Role {
		if role.Equals(UserRole(r)) {
			return true
		}
	}
	return false
}

func (User) TableName() string {
	return "users"
}

func (u User) GetID() string {
	return u.ID.String()
}

// TODO: maybe refactor, don't like this way
func (u User) Roles() []UserRole {
	roles := make([]UserRole, len(u.Role))
	for i, r := range u.Role {
		roles[i] = UserRole(r)
	}
	return roles
}

func (u User) RolesStrings() []string {
	roles := make([]string, len(u.Role))
	for i, r := range u.Role {
		roles[i] = r
	}
	return roles
}
