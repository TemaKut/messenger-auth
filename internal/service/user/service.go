package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	userstorage "github.com/TemaKut/messenger-auth/internal/storage/user"
)

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Register(ctx context.Context, params userdto.RegisterParams) (userdto.User, error) {
	user, err := s.storage.UserCreate(ctx, userstorage.UserCreateParams{
		Name:         params.Name,
		LastName:     params.LastName,
		Email:        params.Email,
		PasswordHash: s.hashUserPassword(params.Password),
	})
	if err != nil {
		return userdto.User{}, fmt.Errorf("error register user. %w", err)
	}

	return encodeUser(user), nil
}

func (s *Service) hashUserPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}
