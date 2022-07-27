package user_locations

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/data/queriers"
)

const UserLocations = "user_locations"

type userLocationsProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.UserLocationsProvider {
	return &userLocationsProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, UserLocations),
	}
}

func (u userLocationsProvider) AddUserLocationsBatch(_ context.Context, userLocations []model.UserLocation) error {
	if err := u.db.Omit(postgres.ID).Create(&userLocations).Error; err != nil {
		return errors.Wrap(err, "failed to insert user locations")
	}
	return nil
}
