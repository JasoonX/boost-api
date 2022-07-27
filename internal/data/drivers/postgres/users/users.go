package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres"
	"github.com/BOOST-2021/boost-app-back/internal/data/model"
	"github.com/BOOST-2021/boost-app-back/internal/data/queriers"
)

const Users = "users"

var userModel = model.User{}

type usersProvider struct {
	log *logrus.Entry
	db  *gorm.DB
}

func New(cfg config.Config) queriers.UsersProvider {
	return &usersProvider{
		db:  cfg.DB(),
		log: cfg.Logging().WithField(postgres.Querier, Users),
	}
}

func (u usersProvider) GetUserByID(_ context.Context, id uuid.UUID) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u usersProvider) ListUsers(_ context.Context) ([]model.User, error) {
	var out []model.User
	if err := u.db.Model(userModel).Find(&out).Error; err != nil {
		return nil, errors.Wrap(err, "failed to list users")
	}
	return out, nil
}

func (u usersProvider) AddUsersBatch(_ context.Context, users []model.User) error {
	if err := u.db.Omit(postgres.ID).Create(&users).Error; err != nil {
		return errors.Wrap(err, "failed to insert users")
	}
	return nil
}
