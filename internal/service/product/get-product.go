package product

import (
	entity "go-ddd/internal/entity/product"
)

type GetProductService struct {
	repository ProductGetter
}

type ProductGetter interface {
	GetByID(id int) (*entity.Product, error)
}

func NewGetProductService(repository ProductGetter) *GetProductService {
	return &GetProductService{repository}
}

func (s *GetProductService) GetByID(id int) (*entity.Product, error) {
	p, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
