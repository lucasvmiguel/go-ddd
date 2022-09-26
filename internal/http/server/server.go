package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTTPServer struct {
	router *chi.Mux
	port   string
}

type NewArgs struct {
	Port string
}

func New(args NewArgs) *HTTPServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	return &HTTPServer{
		router: r,
		port:   args.Port,
	}
}

func (s *HTTPServer) Post(pattern string, handlerFn http.HandlerFunc) {
	s.router.Post(pattern, handlerFn)
}

func (s *HTTPServer) Get(pattern string, handlerFn http.HandlerFunc) {
	s.router.Get(pattern, handlerFn)
}

func (s *HTTPServer) Listen() {
	http.ListenAndServe(s.port, s.router)
}
