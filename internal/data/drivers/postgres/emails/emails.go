package emails

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/data/queriers"
)

const Emails = "emails"

var emailsModel = model.Email{}

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

func (u emailsProvider) GetEmail(_ context.Context, email string) (*model.Email, error) {
	var out model.Email
	if err := u.db.Model(emailsModel).First(&email, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, data.ErrNotFound
		}
		return nil, errors.Wrap(err, "failed to get email")
	}
	return &out, nil
}

func (u emailsProvider) AddEmailsBatch(_ context.Context, emails []model.Email) error {
	if err := u.db.Create(&emails).Error; err != nil {
		return errors.Wrap(err, "failed to insert emails")
	}
	return nil
}

func (u emailsProvider) AddEmail(_ context.Context, email model.Email) (*model.Email, error) {
	if err := u.db.Create(&email).Error; err != nil {
		return nil, errors.Wrap(err, "failed to insert email")
	}
	return &email, nil
}
