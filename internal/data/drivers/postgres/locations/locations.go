package locations

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

const Locations = "locations"

var locationModel = model.Location{}

type locationsProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.LocationsProvider {
	return &locationsProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Locations),
	}
}

func (p locationsProvider) AddLocationsBatch(_ context.Context, locations []model.Location) error {
	if err := p.db.Omit(postgres.ID).Create(&locations).Error; err != nil {
		return errors.Wrap(err, "failed to insert locations")
	}
	return nil
}

func (p locationsProvider) ListLocations(_ context.Context) ([]model.Location, error) {
	var out []model.Location
	if err := p.db.Model(locationModel).Find(&out).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list locations")
	}
	return out, nil
}
