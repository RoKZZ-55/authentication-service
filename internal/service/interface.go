package service

import "authentication-service/internal/model"

type AuthenticationService interface {
	GetTokens(authentication *model.Token) error
	RefreshTokens(authentication *model.Token) error
}
