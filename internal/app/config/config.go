package config

type Config struct {
	Logger struct {
		Level LoggerLevel
	}
	Server struct {
		Grpc struct {
			Addr string
		}
	}
	Storage struct {
		Postgres struct {
			ConnStr string
		}
	}
}

func NewConfig() *Config { // TODO default + parse from env
	return &defaultConfig
}
