package apiserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

type Server struct {
	router *chi.Mux
	log    *slog.Logger
}

func New(log *slog.Logger) *Server {
	s := &Server{
		router: chi.NewRouter(),
		log:    log,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
