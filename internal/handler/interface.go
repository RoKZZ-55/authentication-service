package handler

import "net/http"

type AuthenticationHandler interface {
	GetToken() http.HandlerFunc
	RefreshToken() http.HandlerFunc
}
