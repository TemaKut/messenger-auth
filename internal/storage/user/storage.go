package userstorage

import (
	"context"
	"database/sql"
	usermodels "github.com/TemaKut/messenger-auth/internal/models/user"
	"github.com/google/uuid"
)

type Storage struct {
	postgresDb *sql.DB
}

func NewStorage(postgresDb *sql.DB) *Storage {
	return &Storage{
		postgresDb: postgresDb,
	}
}

type UserCreateParams struct {
	Name         string
	LastName     string
	Email        string
	PasswordHash string
}

func (s *Storage) UserCreate(ctx context.Context, params UserCreateParams) (*usermodels.User, error) {
	userDbo := UserDbo{
		Id:       uuid.New().String(),
		Name:     params.Name,
		LastName: params.LastName,
		Data: UserDboData{
			PasswordHash: params.PasswordHash,
		},
	}
	// TODO запрос без ORM!
	// TODO в файле миграции создам таблицу вида (Какие-то поисковые поля среди колонок -
	// оставшиеся в колонке data для возможности расширения без миграций)

	//CREATE TABLE users(
	//	id UUID PRIMARY KEY,
	//	name TEXT,
	//	last_name TEXT,
	//	data JSONB
	//);

	return encodeUser(userDbo), nil
}
