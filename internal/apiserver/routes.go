package apiserver

import "github.com/go-chi/chi/v5/middleware"

// configureRouter contains routes advising methods, as well as middleware
func (s *Server) configureRouter() {
	// server middleware
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.URLFormat)
	s.router.Use(s.LogRequest())

	// authentication routes
	s.router.Get("/authentication-service/get-tokens/{uuid}", s.handler.GetToken())
	s.router.Get("/authentication-service/refresh-tokens", s.handler.RefreshToken())
}
