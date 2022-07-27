package emails

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

const Emails = "emails"

type emailsProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.EmailsProvider {
	return &emailsProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Emails),
	}
}

func (u emailsProvider) AddEmailsBatch(_ context.Context, emails []model.Email) error {
	if err := u.db.Omit(postgres.ID).Create(&emails).Error; err != nil {
		return errors.Wrap(err, "failed to insert emails")
	}
	return nil
}
