package factory

import (
	producthandler "go-ddd/internal/http/handler/product"
	productservice "go-ddd/internal/service/product"
	"net/http"
)

type Handlers struct {
	CreateProductHandler http.Handler
}

type NewArgs struct {
	CreateProductService productservice.ProductCreator
}

func New(args NewArgs) (*Handlers, error) {
	createProductHandler := producthandler.NewCreateProductHandler(args.CreateProductService)

	return &Handlers{
		CreateProductHandler: createProductHandler,
	}, nil
}
