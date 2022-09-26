package product

import (
	"encoding/json"
	entity "go-ddd/internal/entity/product"
	"net/http"
)

type CreateProductHandlerRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateProductHandlerResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateProductHandler struct {
	service ProductCreator
}

type ProductCreator interface {
	Create(product entity.Product) (*entity.Product, error)
}

func NewCreateProductHandler(service ProductCreator) http.Handler {
	return &CreateProductHandler{service}
}

func (h *CreateProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	payload := CreateProductHandlerRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.validate(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.Create(entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(CreateProductHandlerResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *CreateProductHandler) validate(payload CreateProductHandlerRequest) error {
	return nil
}
