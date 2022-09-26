package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
)

type MigrateArgs struct {
	DBDataSource    string
	DBDriver        string
	DBMigrationPath string
	IsDown          bool
}

func Migrate(args MigrateArgs) {
	db, err := sql.Open(args.DBDriver, args.DBDataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open(args.DBMigrationPath)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, args.DBDriver, instance)
	if err != nil {
		log.Fatal(err)
	}

	if args.IsDown {
		err = m.Down()
	} else {
		err = m.Up()
	}

	if err != nil {
		log.Fatal(err)
	}
}
