package model

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	UUID         string `json:"uuid"`
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty"`
	jwt.RegisteredClaims
}
