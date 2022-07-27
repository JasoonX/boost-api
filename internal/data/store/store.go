package store

import (
	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/emails"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/enums"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/locations"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/news"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/phones"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/user_locations"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/users"
	"github.com/BOOST-2021/boost-app-back/internal/data/queriers"
)

const Store = "store"

type DataProvider interface {
	UsersProvider() queriers.UsersProvider
	NewsProvider() queriers.NewsProvider
	EmailsProvider() queriers.EmailsProvider
	PhonesProvider() queriers.PhonesProvider
	LocationsProvider() queriers.LocationsProvider
	UserLocationsProvider() queriers.UserLocationsProvider
	EnumProvider() queriers.EnumsProvider
}

type dataProvider struct {
	cfg config.Config
}

func (d dataProvider) UsersProvider() queriers.UsersProvider {
	return users.New(d.cfg)
}

func (d dataProvider) NewsProvider() queriers.NewsProvider {
	return news.New(d.cfg)
}

func (d dataProvider) EmailsProvider() queriers.EmailsProvider {
	return emails.New(d.cfg)
}

func (d dataProvider) PhonesProvider() queriers.PhonesProvider {
	return phones.New(d.cfg)
}

func (d dataProvider) LocationsProvider() queriers.LocationsProvider {
	return locations.New(d.cfg)
}

func (d dataProvider) UserLocationsProvider() queriers.UserLocationsProvider {
	return user_locations.New(d.cfg)
}

func (d dataProvider) EnumProvider() queriers.EnumsProvider {
	return enums.New(d.cfg)
}

func New(cfg config.Config) DataProvider {
	logging := cfg.Logging().WithField(common.SQLTag, cfg.Driver())
	logging.Info("Data provider connecting...")

	return &dataProvider{
		cfg: cfg,
	}
}
