package store

import (
	"github.com/BOOST-2021/boost-app-back/internal/common"
	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const Store = "store"

type DataProvider interface {
	// GetNewsProvider
	// GetUsersProvider
	// GetCommentsProvider
	// GetReactionsProvider
	// ...
	Ping() error
}

type dataProvider struct {
	db *gorm.DB
}

// TODO: remove
func (p dataProvider) Ping() error {
	return errors.New("failed to ping")
}

func New(cfg config.Config) DataProvider {
	logging := cfg.Logging().WithField(common.SQLTag, cfg.Driver())
	logging.Info("DB connecting...")

	switch cfg.Driver() {
	case config.PostgresDriver:
		return &dataProvider{
			db: cfg.DB(),
		}
	default:
		// should never happen normally
		panic(errors.Errorf("provided driver unsupported: %s", cfg.Driver()))
	}
}
