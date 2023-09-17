package storage

import (
	"authentication-service/internal/model"
	"context"
)

type AuthenticationStorage interface {
	GetToken(ctx context.Context, userAuthentication *model.UserAuthentication) error
	RefreshToken(ctx context.Context, userAuthentication *model.UserAuthentication) error
}
