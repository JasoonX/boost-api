package convert

import "github.com/BOOST-2021/boost-app-back/internal/data/model"

func FromRequestRole(role string) model.UserRole {
	return model.UserRole(role)
}
