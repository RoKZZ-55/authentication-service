package handler

import (
	"authentication-service/internal/model"
	"authentication-service/pkg/http/response"
	"authentication-service/pkg/logger/sl"
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

func (h *Handler) GetToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()

		userAuthentication := &model.UserAuthentication{
			GUID: chi.URLParam(r, "guid"),
		}

		if err := validator.New().Var(userAuthentication.GUID, "uuid"); err != nil {
			slog.Error("user guid validation error", sl.Err(err))
			response.Error(w, http.StatusBadRequest, errors.New("invalid guid"))
			return
		}

		if err := h.storage.GetToken(ctx, userAuthentication); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		data := &Response{
			AccessToken:  userAuthentication.AccessToken,
			RefreshToken: userAuthentication.RefreshToken,
		}

		response.Respond(w, http.StatusOK, data)
	}
}

func (h *Handler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()

		userAuthentication := &model.UserAuthentication{
			RefreshToken: r.Header.Get("Authentication"),
		}

		if err := validator.New().Var(userAuthentication.RefreshToken, "base64"); err != nil {
			slog.Error("refresh token validation error", sl.Err(err))
			response.Error(w, http.StatusBadRequest, errors.New("incorrect token"))
			return
		}

		if err := h.storage.RefreshToken(ctx, userAuthentication); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		data := &Response{
			AccessToken:  userAuthentication.AccessToken,
			RefreshToken: userAuthentication.RefreshToken,
		}

		response.Respond(w, http.StatusOK, data)
	}
}
