package userservice

import (
	"context"
	"fmt"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
)

type userCredentialsEmailAuthenticator struct {
	storage  Storage
	email    string
	password string
}

func newUserCredentialsEmailAuthenticator(
	storage Storage,
	email string,
	password string,
) *userCredentialsEmailAuthenticator {
	return &userCredentialsEmailAuthenticator{
		storage:  storage,
		email:    email,
		password: password,
	}
}

func (u *userCredentialsEmailAuthenticator) authentify(ctx context.Context) (*usermodels.User, error) {
	user, err := u.storage.UserByEmail(ctx, u.email)
	if err != nil {
		return nil, fmt.Errorf("error fetch user by email %s. %w", u.email, err)
	}

	if !user.ComparePassword(u.password) {
		return nil, fmt.Errorf("invalid password") // TODO specific error
	}

	return user, nil
}
