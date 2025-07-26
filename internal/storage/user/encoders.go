package userstorage

import usermodels "github.com/TemaKut/messenger-auth/internal/models/user"

func encodeUser(userDbo UserDbo) *usermodels.User {
	return usermodels.NewUserFromDb(
		userDbo.Id,
		userDbo.Data.Name,
		userDbo.Data.LastName,
		userDbo.Email,
		userDbo.Data.PasswordHash,
	)
}
