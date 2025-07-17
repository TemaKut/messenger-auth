package userstorage

type UserDbo struct {
	Id    string
	Email string
	Data  UserDboData
}

type UserDboData struct {
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	PasswordHash string `json:"password_hash"`
}
