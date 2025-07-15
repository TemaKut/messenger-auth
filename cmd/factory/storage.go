package factory

import (
	userservice "github.com/TemaKut/messenger-auth/internal/service/user"
	userstorage "github.com/TemaKut/messenger-auth/internal/storage/user"
	"github.com/google/wire"
)

var StorageSet = wire.NewSet(
	ProvideUserStorage,
	wire.Bind(new(userservice.Storage), new(*userstorage.Storage)),
)

//type UserStoragePostgresDb *sql.DB
//
//func ProvideUserStoragePostgresDb(cfg *config.Config) (*UserStoragePostgresDb, error) {
//	return nil, nil
//}

func ProvideUserStorage(
// TODO db UserStoragePostgresDb,
) *userstorage.Storage {
	return userstorage.NewStorage()
}
