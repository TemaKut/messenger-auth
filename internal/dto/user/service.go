package userdto

import "time"

type RegisterParams struct {
	Name     string `validate:"required"`
	LastName string
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=40"`
}

type UserAuthorizeParams struct {
	Credentials UserAuthorizeCredentials
}

type UserAuthorizeCredentials struct {
	Email *UserAuthorizeEmailCredential
}

type UserAuthorizeEmailCredential struct {
	Email    string
	Password string
}

type UserAuthorizeResult struct {
	User       User
	AuthParams AuthParams
}

type AuthParams struct {
	AccessToken  AuthToken
	RefreshToken AuthToken
}

type AuthToken struct {
	Token     string
	ExpiredAt time.Time
}
