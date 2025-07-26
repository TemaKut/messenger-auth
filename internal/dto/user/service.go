package userdto

import "time"

type RegisterParams struct {
	Name     string
	LastName string
	Email    string
	Password string
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
