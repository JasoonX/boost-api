package scripts

import (
	"github.com/BOOST-2021/boost-app-back/assets"
	"github.com/BOOST-2021/boost-app-back/internal/config"
)

func Run(cfg config.Config, names ...string) error {
	for _, name := range names {
		if err := assets.Scripts.Run(cfg, name); err != nil {
			return err
		}
	}
	if len(names) == 0 {
		if err := assets.Scripts.RunAll(cfg); err != nil {
			return err
		}
	}
	return nil
}
