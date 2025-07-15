package userstorage

type UserDbo struct {
	Id       string
	Name     string
	LastName string
	Data     UserDboData
}

type UserDboData struct {
	PasswordHash string `json:"password_hash"`
}
