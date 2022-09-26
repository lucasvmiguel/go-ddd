package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RespondJSON(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body != nil {
		w.Write(body)
	}
}

func GetIntParam(r *http.Request, param string) (int, error) {
	p, err := strconv.Atoi(chi.URLParam(r, param))
	if err != nil {
		return 0, err
	}

	return p, nil
}

func GetParam(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}
