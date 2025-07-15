package config

var defaultConfig Config

func init() {
	defaultConfig.Logger.Level = LoggerLevelDebug
	defaultConfig.Server.Grpc.Addr = ":8001"
}
