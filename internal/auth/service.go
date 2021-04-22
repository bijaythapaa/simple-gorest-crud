package auth

import (
	"context"
	"time"

	"github.com/dbijay/simple-gorest-crud/internal/entity"
	"github.com/dbijay/simple-gorest-crud/internal/errors"
	"github.com/dbijay/simple-gorest-crud/pkg/log"
	"github.com/dgrijalva/jwt-go"
)

// Service Encapsulates the authentication Logic.
type Service interface {
	// authenticates the user with username and password
	// returns JWT token, if authentication successed, else an error is occured.
	Login(ctx context.Context, username, password string) (string, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns user ID.
	GetID() string
	// GetName returns user name.
	GetName() string
}

type service struct {
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// Login authenticates user and generates a JWT token if authentication successed.
// Else, an error is returned.
func (s service) Login(ctx context.Context, username, password string) (string, error) {
	if identity := s.authenticate(ctx, username, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// NewService creates a new authentication service.
func NewService(signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{signingKey: signingKey, tokenExpiration: tokenExpiration, logger: logger}
}

// authenticate authenticates a user using username and password.
// If username and password are correct, an identity is returned. Else, nil is returned.
func (s service) authenticate(ctx context.Context, username, password string) Identity {
	logger := s.logger.With(ctx, "user", username)

	// TODO: the following authentication logic is only for demo purpose.
	if username == "demo" && password == "pass" {
		logger.Infof("authentication successful for user: %v", username)
		return entity.User{ID: "100", Name: "demo"}
	}

	logger.Infof("authentication failed for user: %v", username)
	return nil
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   identity.GetID(),
		"name": identity.GetName(),
		"exp":  time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
