package handler

import "net/http"

func (h *Handler) GetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *Handler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
