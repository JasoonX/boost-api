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

const Phone = "phone"

type phonesProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.PhonesProvider {
	return &phonesProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Phone),
	}
}

func (u phonesProvider) AddPhonesBatch(_ context.Context, phones []model.Phone) error {
	if err := u.db.Omit(postgres.ID).Create(&phones).Error; err != nil {
		return errors.Wrap(err, "failed to insert phones")
	}
	return nil
}
