package product

import (
	entity "go-ddd/internal/entity/product"
)

type CreateProductService struct {
	repository ProductCreator
}

type ProductCreator interface {
	Create(product entity.Product) (*entity.Product, error)
}

func NewCreateProductService(repository ProductCreator) *CreateProductService {
	return &CreateProductService{repository}
}

func (s *CreateProductService) Create(product entity.Product) (*entity.Product, error) {
	p, err := s.repository.Create(product)
	if err != nil {
		return nil, err
	}

	return p, nil
}
