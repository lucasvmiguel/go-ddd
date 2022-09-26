package factory

import (
	"database/sql"
	productrepository "go-ddd/internal/repository/product"
)

type Repositories struct {
	Product *productrepository.Repository
}

type NewArgs struct {
	DBConn *sql.DB
}

func New(args NewArgs) (*Repositories, error) {
	productRepository, err := productrepository.New(args.DBConn)

	return &Repositories{
		Product: productRepository,
	}, err
}
