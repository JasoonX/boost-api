package phones

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

const Phones = "phones"

type phonesProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.PhonesProvider {
	return &phonesProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Phones),
	}
}

func (u phonesProvider) AddPhonesBatch(_ context.Context, phones []model.Phone) error {
	if err := u.db.Create(&phones).Error; err != nil {
		return errors.Wrap(err, "failed to insert phones")
	}
	return nil
}

func (u phonesProvider) AddPhone(_ context.Context, phone model.Phone) (*model.Phone, error) {
	if err := u.db.Create(&phone).Error; err != nil {
		return nil, errors.Wrap(err, "failed to insert phone")
	}
	return &phone, nil
}
