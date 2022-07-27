package queriers

import (
	"context"

	"github.com/google/uuid"

	"github.com/BOOST-2021/boost-app-back/internal/data/model"
)

// How methods namings are built:
// OpName+EntityName+Args
// We use: add/delete/update/get/list/paginate/search/filter

type UsersProvider interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)

	ListUsers(ctx context.Context) ([]model.User, error)

	AddUsersBatch(ctx context.Context, users []model.User) error
}

type NewsProvider interface {
	ListNews(ctx context.Context) ([]model.News, error)

	IndexNews(ctx context.Context, id uuid.UUID, title, text, meta string) error

	AddNewsBatch(ctx context.Context, news []model.News) error
}

type EmailsProvider interface {
	AddEmailsBatch(ctx context.Context, emails []model.Email) error
}

type PhonesProvider interface {
	AddPhonesBatch(ctx context.Context, phones []model.Phone) error
}

type LocationsProvider interface {
	ListLocations(ctx context.Context) ([]model.Location, error)

	AddLocationsBatch(ctx context.Context, locations []model.Location) error
}

type UserLocationsProvider interface {
	AddUserLocationsBatch(ctx context.Context, userLocations []model.UserLocation) error
}

type EnumsProvider interface {
	ListCountries(ctx context.Context) ([]model.Country, error)
}
