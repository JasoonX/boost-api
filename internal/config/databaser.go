package config

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDriver list of supported drivers
const (
	PostgresDriver = "postgres"
)

type Databaser interface {
	DB() *gorm.DB
	Driver() string
}

type databaser struct {
	db     *gorm.DB
	driver string
}

type yamlDatabaseConfig struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	SslMode  string `yaml:"ssl_mode"`
	Driver   string `yaml:"driver"`
}

func (c yamlDatabaseConfig) toPSQLPath() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SslMode)
}

func NewDatabaser(conn, driver string) Databaser {

	var sess *gorm.DB
	var err error

	switch driver {
	case PostgresDriver:
		sess, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
		if err != nil {
			panic(errors.Wrapf(err, "failed to open database connection: %s", conn))
		}

		rawDB, err := sess.DB()
		if err != nil {
			panic(errors.Wrapf(err, "failed to establish database connection: %s", conn))
		}
		if rawDB.Ping() != nil {
			panic(errors.Wrapf(err, "failed to ping database: %s", conn))
		}
	default:
		panic(errors.Errorf("provided driver unsupported: %s", driver))
	}
	return &databaser{
		db:     sess,
		driver: driver,
	}
}

func (d *databaser) DB() *gorm.DB {
	// TODO: turn this off for prod
	return d.db.Debug()
}

func (d *databaser) Driver() string {
	return d.driver
}
