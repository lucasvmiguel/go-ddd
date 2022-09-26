package factory

import (
	productrepository "go-ddd/internal/repository/product"
	productservice "go-ddd/internal/service/product"
)

type Services struct {
	CreateProductService productservice.ProductCreator
	GetProductService    productservice.ProductGetter
}

type NewArgs struct {
	ProductRepository *productrepository.Repository
}

func New(args NewArgs) (*Services, error) {
	createProductService := productservice.NewCreateProductService(args.ProductRepository)
	getProductService := productservice.NewGetProductService(args.ProductRepository)

	return &Services{
		CreateProductService: createProductService,
		GetProductService:    getProductService,
	}, nil
}
