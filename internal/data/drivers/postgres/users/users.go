package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/data"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres"
	"github.com/BOOST-2021/boost-app-back/internal/data/drivers/postgres/emails"
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

func (u usersProvider) AddUser(_ context.Context, user model.User) (*model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, errors.Wrap(err, "failed to insert user")
	}
	return &user, nil
}

func (u usersProvider) GetUserByID(_ context.Context, id uuid.UUID) (*model.User, error) {
	var out model.User
	if err := u.db.Model(userModel).Where("id = ?", id).First(&out).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, data.ErrNotFound
		}
		return nil, errors.Wrap(err, "failed to get user by id")
	}
	return &out, nil
}

func (u usersProvider) GetUserByEmail(_ context.Context, email string) (*model.User, error) {
	var out model.User
	// TODO: refactor all joins
	if err := u.db.Model(userModel).Joins(fmt.Sprintf("LEFT JOIN %s on %s.user_id = %s.id", emails.Emails, emails.Emails, Users)).First(&out, fmt.Sprintf("%s.email = ?", emails.Emails), email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, data.ErrNotFound
		}
		return nil, errors.Wrap(err, "failed to get user by email")
	}
	return &out, nil
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
