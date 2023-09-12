package handler

import (
	"authentication-service/internal/service"
)

type Handler struct {
	service               service.Service
	AuthenticationHandler *AuthenticationHandler
}

func New(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
