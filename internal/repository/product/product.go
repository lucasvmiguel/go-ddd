package product

import (
	"database/sql"
	entity "go-ddd/internal/entity/product"

	"github.com/blockloop/scan"
)

const (
	tableName      = "products"
	insertQuery    = "INSERT INTO " + tableName + " (title, description) VALUES ($1, $2);"
	selectOneQuery = "SELECT * FROM " + tableName + " WHERE id = $1;"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) (*Repository, error) {
	return &Repository{db}, nil
}

func (r *Repository) Create(product entity.Product) (*entity.Product, error) {
	res, err := r.db.Exec(insertQuery, product.Title, product.Description)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(selectOneQuery, id)
	if err != nil {
		return nil, err
	}

	p := &entity.Product{}
	err = scan.Row(p, rows)
	if err != nil {
		return nil, err
	}

	return p, nil
}
