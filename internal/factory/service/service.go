package factory

import (
	productrepository "go-ddd/internal/repository/product"
	productservice "go-ddd/internal/service/product"
)

type Services struct {
	CreateProductService productservice.ProductCreator
}

type NewArgs struct {
	ProductRepository *productrepository.Repository
}

func New(args NewArgs) (*Services, error) {
	createProductService := productservice.NewCreateProductService(args.ProductRepository)

	return &Services{
		CreateProductService: createProductService,
	}, nil
}
