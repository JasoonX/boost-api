package assets

import (
	"embed"

	"github.com/rubenv/sql-migrate"

	"github.com/BOOST-2021/boost-app-back/util/source"
)

//go:embed migrations/postgres/*.sql
var migrationsSQL embed.FS

//go:embed scripts/postgres/*.sql
var scriptsSQL embed.FS

var Migrations = migrate.EmbedFileSystemMigrationSource{
	FileSystem: migrationsSQL,
	Root:       "migrations/postgres",
}

var Scripts = source.ScriptRunnerSource{
	FileSystem: scriptsSQL,
}
