package main

import (
	"go-ddd/pkg/db"
)

type config struct {
	DBDataSource    string
	DBDriver        string
	DBMigrationPath string
}

func main() {
	cfg := config{
		DBDriver:        "sqlite3",
		DBDataSource:    "./database.db",
		DBMigrationPath: "./db/migrations",
	}

	db.Migrate(db.MigrateArgs{
		DBDriver:        cfg.DBDriver,
		DBDataSource:    cfg.DBDataSource,
		DBMigrationPath: cfg.DBMigrationPath,
		IsDown:          true,
	})
}
