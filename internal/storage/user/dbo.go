package userstorage

import (
	"encoding/json"
	"fmt"
)

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

func (u *UserDboData) Scan(src any) error {
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}

	srcBytes, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("src must be []byte")
	}

	if err := json.Unmarshal(srcBytes, u); err != nil {
		return fmt.Errorf("error unmarshalling user data. %w", err)
	}

	return nil
}
