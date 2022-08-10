package source

import (
	"embed"
	"io/fs"

	"github.com/pkg/errors"

	"github.com/BOOST-2021/boost-app-back/internal/config"
)

type ScriptRunnerSource struct {
	FileSystem embed.FS
}

func (s ScriptRunnerSource) Run(cfg config.Config, name string) error {
	cfg.Logging().WithField("path", name).Info("Running script")
	sqlScript, err := s.FileSystem.ReadFile(name)
	if err != nil {
		return errors.Wrap(err, "failed to read script")
	}
	return cfg.DB().Exec(string(sqlScript)).Error
}

func (s ScriptRunnerSource) RunAll(cfg config.Config) error {
	return fs.WalkDir(s.FileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if err := s.Run(cfg, path); err != nil {
			return err
		}
		return nil
	})
}
