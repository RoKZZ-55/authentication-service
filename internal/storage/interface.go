package storage

import (
	"authentication-service/internal/model"
)

type AuthenticationStorage interface {
	GetToken(authentication *model.Token) error
	RefreshToken(authentication *model.Token) error
}
