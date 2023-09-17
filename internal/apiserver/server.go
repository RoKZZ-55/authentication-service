package apiserver

import (
	"authentication-service/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router  *chi.Mux
	handler *handler.Handler
}

func New(handler *handler.Handler) *Server {
	s := &Server{
		router:  chi.NewRouter(),
		handler: handler,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
