package userstorage

import usermodels "github.com/TemaKut/messenger-auth/internal/models/user"

func decodeUser(user *usermodels.User) UserDbo {
	return UserDbo{
		Id:    user.Id(),
		Email: user.Email(),
		Data: UserDboData{
			Name:         user.Name(),
			LastName:     user.LastName(),
			PasswordHash: user.PasswordHash(),
		},
	}
}
