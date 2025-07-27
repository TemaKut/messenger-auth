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
	Services struct {
		User struct {
			AuthTokenSecret              string
			AccessTokenLifetimeDuration  string
			RefreshTokenLifetimeDuration string
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
