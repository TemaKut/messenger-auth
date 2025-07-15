package userservice

import (
	userdto "github.com/TemaKut/messenger-auth/internal/dto/user"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
)

func encodeUser(user *usermodels.User) userdto.User {
	return userdto.User{
		Id:       user.Id(),
		Name:     user.Name(),
		LastName: user.LastName(),
	}
}
