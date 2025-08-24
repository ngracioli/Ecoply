package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const (
	KeyAppName     = "APP_NAME"
	KeyEnvironment = "ENVIRONMENT"
	KeyPort        = "PORT"
	KeyHost        = "HOST"
	KeyDebug       = "DEBUG"
)

type Config struct {
	AppName        string `env:"APP_NAME" envDefault:"Ecoply"`
	AppEnvironment string `env:"APP_ENV" envDefault:"dev"`
	AppDebug       bool   `env:"APP_DEBUG" envDefault:"false"`

	ServerHost string `env:"SERVER_HOST" envDefault:"localhost"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`
}

var (
	config Config
	loaded bool
)

func Load(filenames ...string) error {
	if loaded {
		return nil
	}

	if filenames != nil {
		err := godotenv.Load(filenames...)
		if err != nil {
			return err
		}
	}

	err := env.Parse(&config)
	if err != nil {
		return ErrFailedToLoadConfig
	}

	loaded = true
	return nil
}

func GetConfig() *Config {
	if !loaded {
		return nil
	}

	return &config
}

func IsLoaded() bool {
	return loaded
}
