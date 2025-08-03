package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Logger struct {
		Level LoggerLevel `env:"LEVEL" envDefault:"1"`
	} `envPrefix:"LOGGER_"`
	Server struct {
		Grpc struct {
			Addr string `env:"ADDR" envDefault:":8001"`
		} `envPrefix:"GRPC_"`
	} `envPrefix:"SERVER_"`
	Services struct {
		User struct {
			AuthTokenSecret              string `env:"AUTH_TOKEN_SECRET" envDefault:"fake secret"`
			AccessTokenLifetimeDuration  string `env:"ACCESS_TOKEN_LIFETIME" envDefault:"30m"`
			RefreshTokenLifetimeDuration string `env:"REFRESH_TOKEN_LIFETIME" envDefault:"24h"`
		} `envPrefix:"USER_"`
	} `envPrefix:"SERVICES_"`
	Storage struct {
		Postgres struct {
			ConnStr string `env:"CONN_STR" envDefault:"postgres://root:root@localhost:5432/postgres?sslmode=disable"`
		} `envPrefix:"POSTGRES_"`
	} `envPrefix:"STORAGE_"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("error parse config. %w", err)
	}

	return &cfg, nil
}
