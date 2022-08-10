package cli

import (
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/BOOST-2021/boost-app-back/internal/config"
	"github.com/BOOST-2021/boost-app-back/internal/listener"
	"github.com/BOOST-2021/boost-app-back/utils/fake"
	"github.com/BOOST-2021/boost-app-back/utils/migrate"
	"github.com/BOOST-2021/boost-app-back/utils/scripts"
)

func Run(args []string) bool {
	cfg := config.New(os.Getenv("CONFIG"))
	log := cfg.Logging()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.Error("internal panicked: ", rvr)
		}
	}()

	svc := listener.New(cfg)

	app := &cli.App{
		Commands: cli.Commands{
			{
				Name:  "run",
				Usage: "run listener daemon",
				Action: func(c *cli.Context) error {
					if err := svc.Listen(); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate service database",
				Subcommands: cli.Commands{
					{
						Name:  "up",
						Usage: "migrate database up",
						Action: func(c *cli.Context) error {
							log.Debug("Migrating up")
							if err := migrate.Migrate(cfg, migrate.Up); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "down",
						Usage: "migrate database down",
						Action: func(c *cli.Context) error {
							log.Debug("Migrating down")
							if err := migrate.Migrate(cfg, migrate.Down); err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
			{
				Name:  "scripts",
				Usage: "apply sql scripts to db",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Value:   "default",
						Usage:   "sql scripts name to apply (default runs all), separated by comma",
					},
				},
				Action: func(c *cli.Context) error {
					log.Debug("Running scripts...")
					if fName := c.String("name"); fName != "default" {
						if err := scripts.Run(cfg, strings.Split(fName, ",")...); err != nil {
							return err
						}
						return nil
					}
					if err := scripts.Run(cfg); err != nil {
						return err
					}
					return nil
				},
			},
			{
				// TODO: for dev purposes only, consider removing this in the future
				Name:  "fake",
				Usage: "fake data in the database",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "mode",
						Aliases: []string{"m"},
						Value:   "default",
						Usage:   "mode of the fake to run",
					},
				},
				Action: func(c *cli.Context) error {
					log.Debug("generating fake data...")
					fakeCfg, err := fake.NewConfig(os.Getenv("FAKE_CONFIG"))
					if err != nil {
						return err
					}
					fakeGenerator := fake.New(cfg, fakeCfg)
					switch c.Args().First() {
					// additional fake modes can be added here
					default:
						if err := fakeGenerator.Default(); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(args); err != nil {
		log.Fatal(err, ": service initialization failed")
		return false
	}

	return true
}
