package config

var defaultConfig Config

func init() {
	defaultConfig.Logger.Level = LoggerLevelDebug
	defaultConfig.Server.Grpc.Addr = ":8001"

	defaultConfig.Storage.Postgres.ConnStr = "postgres://root:root@localhost:5432/postgres?sslmode=disable"
}
