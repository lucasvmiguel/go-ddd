package main

import (
	"go-ddd/internal/db"
	handlerfactory "go-ddd/internal/factory/handler"
	repositoryfactory "go-ddd/internal/factory/repository"
	servicefactory "go-ddd/internal/factory/service"

	"log"

	httpserver "go-ddd/internal/http/server"
)

type config struct {
	DBDataSource   string
	DBDriver       string
	HTTPServerPort string
}

func main() {
	cfg := config{
		DBDriver:       "sqlite3",
		DBDataSource:   "./database.db",
		HTTPServerPort: ":3000",
	}

	dbConn, err := db.Connect(db.ConnectArgs{
		Driver:     cfg.DBDriver,
		DataSource: cfg.DBDataSource,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	repositories, err := repositoryfactory.New(repositoryfactory.NewArgs{
		DBConn: dbConn,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	services, err := servicefactory.New(servicefactory.NewArgs{
		ProductRepository: repositories.Product,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers, err := handlerfactory.New(handlerfactory.NewArgs{
		CreateProductService: services.CreateProductService,
		GetProductService:    services.GetProductService,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	httpServer := httpserver.New(httpserver.NewArgs{Port: cfg.HTTPServerPort})

	httpServer.Post("/products", handlers.CreateProductHandler.ServeHTTP)
	httpServer.Get("/products/{id}", handlers.GetProductHandler.ServeHTTP)

	httpServer.Listen()
}
