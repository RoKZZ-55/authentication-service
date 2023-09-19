package storage

import (
	"authentication-service/internal/model"
	"authentication-service/pkg/logger/sl"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrTokenNotValid     = errors.New("token is not valid")
	ErrUserNotFound      = errors.New("user not found")
	ErrInternal          = errors.New("unexpected server error")
	ErrWrongRefreshToken = errors.New("wrong refresh token")
)

type createAndUpdateHashParameters struct {
	guidToken    string
	refreshToken string
	collection   *mongo.Collection
	filter       primitive.D
}

// validExpRefreshToken token lifetime check
func validExpRefreshToken(expStr string) error {
	exp, err := strconv.Atoi(expStr)

	if err != nil || int64(exp) < time.Now().Unix() {
		slog.Error("refresh token is not valid by expiration time")
		return ErrTokenNotValid
	}

	return nil
}

// createAccessAndRefreshToken creating a pair of tokens,
// access (jwt-type) and refresh (arbitrary type: "expiration time/token guid/random character set")
func (s *Storage) createAccessAndRefreshToken(userAuthentication *model.UserAuthentication) error {
	// creating a guid for mapping refresh and access tokens
	userAuthentication.TokenGUID = uuid.NewString()

	// creation of a new access token
	claims := &jwt.RegisteredClaims{
		ID:        userAuthentication.TokenGUID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(s.cfg.AccessTokenLifetime))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	var err error

	userAuthentication.AccessToken, err = token.SignedString([]byte(s.cfg.AccessSecretKey))
	if err != nil {
		slog.Error("error creating access token", sl.Err(err))
		return ErrInternal
	}

	// creation of a new refresh token
	lifetime := time.Now().Add(time.Minute * time.Duration(s.cfg.RefreshTokenLifetime)).Unix()
	randUint := rand.Uint64() >> 15
	userAuthentication.RefreshToken = fmt.Sprintf("%d/%s/%d", lifetime, userAuthentication.TokenGUID, randUint)

	return nil
}

// createAndUpdateHash creating bcrypt hash refresh token and updating it in the database
func createAndUpdateHash(ctx context.Context, p *createAndUpdateHashParameters) error {
	// creating a bcrypt hash of the refresh token to store it in the database
	refreshTokenHash, err := bcrypt.GenerateFromPassword([]byte(p.refreshToken), 7)
	if err != nil {
		slog.Error("error of creating bcrypt hash refresh token", sl.Err(err))
		return ErrInternal
	}

	// updating token guid and token refresh hash in the database
	_, err = p.collection.UpdateOne(
		ctx,
		p.filter,
		bson.D{{Key: "$set", Value: bson.M{"token_guid": p.guidToken, "refresh_token_hash": refreshTokenHash}}},
	)

	if err != nil {
		slog.Error("error of updating the database document by filter", sl.Err(err))
		return ErrInternal
	}

	return nil
}

func (s *Storage) GetToken(ctx context.Context, userAuthentication *model.UserAuthentication) error {
	collection := s.db.Collection(s.cfg.Collection)
	bsonUserGUID := bson.D{{Key: "user_guid", Value: userAuthentication.GUID}}

	// check if the guid exists in the database
	if err := collection.FindOne(ctx, bsonUserGUID).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			slog.Error(
				"database document not found by user guid",
				sl.Err(err),
			)
			return ErrUserNotFound
		}

		slog.Error("error of finding a document by user guid", sl.Err(err))
		return ErrInternal
	}

	if err := s.createAccessAndRefreshToken(userAuthentication); err != nil {
		return err
	}

	parameters := &createAndUpdateHashParameters{
		guidToken:    userAuthentication.TokenGUID,
		refreshToken: userAuthentication.RefreshToken,
		collection:   collection,
		filter:       bsonUserGUID,
	}

	if err := createAndUpdateHash(ctx, parameters); err != nil {
		return err
	}

	userAuthentication.RefreshToken = base64.StdEncoding.EncodeToString([]byte(userAuthentication.RefreshToken))

	return nil
}

func (s *Storage) RefreshToken(ctx context.Context, userAuthentication *model.UserAuthentication) error {
	refreshTokenByte, err := base64.StdEncoding.DecodeString(userAuthentication.RefreshToken)
	if err != nil {
		slog.Error("error decoding refresh token from base64 to []byte", sl.Err(err))
		return ErrInternal
	}

	refreshTokenString := string(refreshTokenByte)
	refreshTokenSplit := strings.Split(refreshTokenString, "/")

	// token lifetime check
	if err := validExpRefreshToken(refreshTokenSplit[0]); err != nil {
		return err
	}

	collection := s.db.Collection(s.cfg.Collection)
	bsonTokenGUID := bson.D{{Key: "token_guid", Value: refreshTokenSplit[1]}}

	var result bson.M
	if err := collection.FindOne(ctx, bsonTokenGUID).Decode(&result); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			slog.Error(
				"database document not found by token guid",
				sl.Err(err),
			)
			return ErrUserNotFound
		}

		slog.Error("error of finding a document by token guid", sl.Err(err))
		return ErrInternal

	}

	pb, ok := result["refresh_token_hash"].(primitive.Binary)
	if !ok {
		slog.Error("data type is not primitive.Binary", slog.Any("refresh_token_hash", result["refresh_token_hash"]))
		return ErrInternal
	}

	if err = bcrypt.CompareHashAndPassword(pb.Data, []byte(refreshTokenString)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			slog.Error("refresh token does not match the hash of the token in the database", sl.Err(err))
			return ErrWrongRefreshToken
		}

		slog.Error("error comparing refresh token and its hash from the database", sl.Err(err))
		return ErrInternal
	}

	if err := s.createAccessAndRefreshToken(userAuthentication); err != nil {
		return err
	}

	parameters := &createAndUpdateHashParameters{
		guidToken:    userAuthentication.TokenGUID,
		refreshToken: userAuthentication.RefreshToken,
		collection:   collection,
		filter:       bsonTokenGUID,
	}

	if err := createAndUpdateHash(ctx, parameters); err != nil {
		return err
	}

	userAuthentication.RefreshToken = base64.StdEncoding.EncodeToString([]byte(userAuthentication.RefreshToken))

	return nil
}
