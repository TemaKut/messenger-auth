package userservice

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type authTokenType = string

const (
	authTokenTypeAccess  authTokenType = "access"
	authTokenTypeRefresh authTokenType = "refresh"
)

type authToken struct {
	claims struct {
		jwt.RegisteredClaims

		Type authTokenType `json:"type"`
	}
}

func newAuthToken() *authToken {
	return &authToken{}
}

func (a *authToken) setSubject(sub string) *authToken {
	a.claims.Subject = sub

	return a
}

func (a *authToken) setExpiredAt(exp time.Time) *authToken {
	a.claims.ExpiresAt = jwt.NewNumericDate(exp)

	return a
}

func (a *authToken) setType(t authTokenType) *authToken {
	a.claims.Type = t

	return a
}

func (a *authToken) build(secret string) (string, error) {
	if a.claims.Type == "" {
		return "", fmt.Errorf("error token has no type")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = a.claims

	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("error sign token. %w", err)
	}

	return tokenStr, nil
}
