package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName        string `env:"APP_NAME" envDefault:"Ecoply"`
	AppEnvironment string `env:"APP_ENV" envDefault:"dev"`
	AppDebug       bool   `env:"APP_DEBUG" envDefault:"false"`

	ServerHost string `env:"SERVER_HOST" envDefault:"localhost"`
	ServerPort uint16 `env:"SERVER_PORT" envDefault:"8080"`

	DBConnection string `env:"DB_CONNECTION" envDefault:"postgres"`
	DBHost       string `env:"DB_HOST" envDefault:"localhost"`
	DBPort       uint16 `env:"DB_PORT" envDefault:"5432"`
	DBDatabase   string `env:"DB_DATABASE" envDefault:"ecoply"`
	DBUsername   string `env:"DB_USERNAME" envDefault:"root"`
	DBPassword   string `env:"DB_PASSWORD" envDefault:"root"`
	DBTimezone   string `env:"DB_TIMEZONE" envDefault:"UTC"`
}

var (
	config Config
	loaded bool
)

func Load(filenames ...string) error {
	if loaded {
		return nil
	}

	if len(filenames) > 0 {
		if err := godotenv.Load(filenames...); err != nil {
			fmt.Printf("%v\n", ErrFailedToLoadEnvFile)
		}
	}

	if err := env.Parse(&config); err != nil {
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

func IsDevelopment() bool {
	return config.AppEnvironment == "dev"
}

func IsProduction() bool {
	return config.AppEnvironment == "prod"
}
