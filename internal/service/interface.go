package service

import "authentication-service/internal/model"

type AuthenticationService interface {
	GetToken(authentication *model.Token) error
	RefreshToken(authentication *model.Token) error
}
