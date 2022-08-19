package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config interface {
	Logger
	Listener
	Databaser
	JwtIssuer
	SSOProvider
	Runtime
}

type config struct {
	Logger
	Listener
	Databaser
	JwtIssuer
	SSOProvider
	Runtime
}

type yamlConfig struct {
	LogLevel     string                `yaml:"log_level"`
	Address      string                `yaml:"address"`
	Database     yamlDatabaseConfig    `yaml:"database"`
	Jwt          yamlJwtConfig         `yaml:"jwt"`
	SSOProviders yamlSSOProviderConfig `yaml:"sso_providers"`
	Runtime      yamlRuntimeConfig     `yaml:"runtime"`
}

func New(path string) Config {
	cfg := yamlConfig{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "failed to read config %s", path))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrapf(err, "failed to unmarshal config %s", path))
	}

	return &config{
		Logger:      NewLogger(cfg.LogLevel),
		Listener:    NewListener(cfg.Address),
		Databaser:   NewDatabaser(cfg.Database.toPSQLPath(), cfg.Database.Driver),
		JwtIssuer:   NewJwtIssuer(cfg.Jwt.Issuer, cfg.Jwt.SecretKey),
		SSOProvider: NewSSOProvider(cfg.SSOProviders),
		Runtime:     NewRuntime(cfg.Runtime.Auth, cfg.Runtime.Environment, cfg.Runtime.Version),
	}
}
