package database

import (
	"ecoply/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DBHost     string
	DBUsername string
	DBPassword string
	DBDatabase string
	DBPort     uint16
	DBTimezone string
}

var dbConfig *DatabaseConfig = nil
var Con *gorm.DB = nil

func getDatabaseConfig() *DatabaseConfig {
	if dbConfig == nil {
		var cfg *config.Config = config.GetConfig()
		dbConfig = &DatabaseConfig{
			DBHost:     cfg.DBHost,
			DBUsername: cfg.DBUsername,
			DBPassword: cfg.DBPassword,
			DBDatabase: cfg.DBDatabase,
			DBPort:     cfg.DBPort,
			DBTimezone: cfg.DBTimezone,
		}
	}
	return dbConfig
}

func mountPostgresDsn(cfg *DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBDatabase, cfg.DBPort, cfg.DBTimezone,
	)
}

func New() *gorm.DB {
	Con = Open()
	Migrate(Con)
	return Con
}

func Open() *gorm.DB {
	var cfg *DatabaseConfig = getDatabaseConfig()
	var dsn string = mountPostgresDsn(cfg)
	con, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("%v: %v\n", ErrFailedToOpenConnectionPostgres, err)
	}
	return con
}
