package enums

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

const Enums = "enums"

var countryModel = model.Country{}

type enumsProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.EnumsProvider {
	return &enumsProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Enums),
	}
}

func (p enumsProvider) ListCountries(_ context.Context) ([]model.Country, error) {
	var out []model.Country
	if err := p.db.Model(countryModel).Find(&out).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list countries")
	}
	return out, nil
}
