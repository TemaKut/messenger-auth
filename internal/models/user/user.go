package usermodels

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
)

type User struct {
	id           string
	name         string
	lastName     string
	email        string
	passwordHash string
}

func NewUser(
	name string,
	lastName string,
	email string,
	password string,
) *User {
	user := &User{
		id:       uuid.Must(uuid.NewV7()).String(),
		name:     name,
		lastName: lastName,
		email:    email,
	}

	user.passwordHash = user.hashPassword(password)

	return user
}

func NewUserFromDb(
	id string,
	name string,
	lastName string,
	email string,
	passwordHash string,
) *User {
	return &User{
		id:           id,
		name:         name,
		lastName:     lastName,
		email:        email,
		passwordHash: passwordHash,
	}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) Id() string {
	return u.id
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) ComparePassword(password string) bool {
	return u.passwordHash == u.hashPassword(password)
}

func (u *User) hashPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum(nil))
}
