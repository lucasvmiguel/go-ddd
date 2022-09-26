package product

import (
	"encoding/json"
	entity "go-ddd/internal/entity/product"
	internalhttp "go-ddd/internal/http"
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

	err = h.validateCreateProduct(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.service.Create(entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(CreateProductHandlerResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	internalhttp.RespondJSON(w, http.StatusCreated, body)
}

func (h *CreateProductHandler) validateCreateProduct(payload CreateProductHandlerRequest) error {
	return nil
}
