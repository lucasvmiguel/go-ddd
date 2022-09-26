package factory

import (
	producthandler "go-ddd/internal/http/handler/product"
	productservice "go-ddd/internal/service/product"
	"net/http"
)

type Handlers struct {
	CreateProductHandler http.Handler
	GetProductHandler    http.Handler
}

type NewArgs struct {
	CreateProductService productservice.ProductCreator
	GetProductService    productservice.ProductGetter
}

func New(args NewArgs) (*Handlers, error) {
	createProductHandler := producthandler.NewCreateProductHandler(args.CreateProductService)
	getProductHandler := producthandler.NewGetProductHandler(args.GetProductService)

	return &Handlers{
		CreateProductHandler: createProductHandler,
		GetProductHandler:    getProductHandler,
	}, nil
}
