package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ConnectArgs struct {
	Driver     string
	DataSource string
}

func Connect(args ConnectArgs) (*sql.DB, error) {
	db, err := sql.Open(args.Driver, args.DataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
