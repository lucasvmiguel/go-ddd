package product

import (
	"database/sql"
	"encoding/json"
	entity "go-ddd/internal/entity/product"
	internalhttp "go-ddd/internal/http"
	"net/http"
)

type GetProductHandlerRequest struct {
	ID int `json:"id"`
}

type GetProductHandlerResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetProductHandler struct {
	service ProductGetter
}

type ProductGetter interface {
	GetByID(id int) (*entity.Product, error)
}

func NewGetProductHandler(service ProductGetter) http.Handler {
	return &GetProductHandler{service}
}

func (h *GetProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := internalhttp.GetIntParam(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payload := GetProductHandlerRequest{
		ID: id,
	}

	err = h.validateGetProduct(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := h.service.GetByID(payload.ID)
	if err == sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(GetProductHandlerResponse{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	internalhttp.RespondJSON(w, http.StatusOK, body)
}

func (h *GetProductHandler) validateGetProduct(payload GetProductHandlerRequest) error {
	return nil
}
