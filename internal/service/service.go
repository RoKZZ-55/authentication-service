package service

import (
	"authentication-service/internal/storage"
)

type Service struct {
	storage               storage.Storage
	authenticationService *AuthenticationService
}

func New(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
