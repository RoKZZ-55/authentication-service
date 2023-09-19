package storage

import (
	"authentication-service/config"
	"authentication-service/internal/model"
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gotest.tools/v3/assert"
)

func Test_validExpRefreshToken(t *testing.T) {
	type args struct {
		expStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "valid",
			args:    args{expStr: fmt.Sprint(time.Now().Add(time.Minute).Unix())},
			wantErr: nil,
		},
		{
			name:    "atoi error",
			args:    args{expStr: "string"},
			wantErr: ErrTokenNotValid,
		},
		{
			name:    "token expired",
			args:    args{expStr: fmt.Sprint(time.Now().Add(-time.Minute).Unix())},
			wantErr: ErrTokenNotValid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validExpRefreshToken(tt.args.expStr)
			assert.Equal(t, tt.wantErr, got)
		})
	}
}

func TestStorage_createAccessAndRefreshToken(t *testing.T) {
	type args struct {
		userAuthentication *model.UserAuthentication
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr error
	}{
		{
			name: "access and refresh token successfully created",
			s: New(new(mongo.Database), &config.Config{
				TokenPair: config.TokenPair{
					AccessSecretKey:      "key123",
					AccessTokenLifetime:  30,
					RefreshTokenLifetime: 525600,
				},
			}),
			args:    args{userAuthentication: new(model.UserAuthentication)},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.createAccessAndRefreshToken(tt.args.userAuthentication)
			assert.Equal(t, tt.wantErr, got)
		})
	}
}
