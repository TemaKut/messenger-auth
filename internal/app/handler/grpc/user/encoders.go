package user

import (
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	authv1 "github.com/TemaKut/messenger-service-proto/gen/go/auth"
)

func encodeUser(user userdto.User) *authv1.User {
	return &authv1.User{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
	}
}
