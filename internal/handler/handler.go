package handler

import (
	"authentication-service/internal/storage"
)

type Handler struct {
	storage               *storage.Storage
	AuthenticationHandler *AuthenticationHandler
}

func New(storage *storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
