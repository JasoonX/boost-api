//go:build fake

package fake

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Users         int `yaml:"users"`
	News          int `yaml:"news"`
	Emails        int `yaml:"emails"`
	Phones        int `yaml:"phones"`
	Locations     int `yaml:"locations"`
	UserLocations int `yaml:"user_locations"`
}

func NewConfig(path string) (*Config, error) {
	cfg := Config{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "failed to read fake config %s", path))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrapf(err, "failed to unmarshal fake config %s", path))
	}

	return &cfg, nil
}
