package convert

import (
	"fmt"

	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/resources"
)

func EntityTypeToPath(entityType resources.EntityType, entityID string) string {
	switch entityType {
	case resources.USERS:
		return fmt.Sprintf("/v1/users/%s", entityID)
	case resources.NEWS:
		return fmt.Sprintf("/v1/news/%s", entityID)
	default:
		return "no specific path"
	}
}

func ToResponseUserStatus(status model.UserStatus) resources.UserStatus {
	switch status {
	case model.UserStatusActive:
		return resources.ACTIVE
	case model.UserStatusInactive:
		return resources.INACTIVE
	case model.UserStatusUnverified:
		return resources.UNVERIFIED
	default:
		panic("unknown user status")
	}
}

func ToResponseUserRole(role model.UserRole) resources.UserRole {
	switch role {
	case model.UserRoleSuperAdmin:
		return resources.ADMIN
	case model.UserRoleViewer:
		return resources.VIEWER
	case model.UserRoleReactViewer:
		return resources.REACT_VIEWER
	default:
		panic("unknown user role")
	}
}

func ToResponseUserRoles(roles []model.UserRole) []resources.UserRole {
	rolesResponse := make([]resources.UserRole, len(roles))
	for i, r := range roles {
		rolesResponse[i] = ToResponseUserRole(r)
	}
	return rolesResponse
}
