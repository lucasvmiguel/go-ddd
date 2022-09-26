package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
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

	db, err := sql.Open(cfg.DBDriver, cfg.DBDataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open(cfg.DBMigrationPath)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, cfg.DBDriver, instance)
	if err != nil {
		log.Fatal(err)
	}

	// modify for Down
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
}
