package database

import (
	"ecoply/internal/config"
	"fmt"
	"log"
	"time"

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

func getDatabaseConfig(cfg *config.Config) *DatabaseConfig {
	return &DatabaseConfig{
		DBHost:     cfg.DBHost,
		DBUsername: cfg.DBUsername,
		DBPassword: cfg.DBPassword,
		DBDatabase: cfg.DBDatabase,
		DBPort:     cfg.DBPort,
		DBTimezone: cfg.DBTimezone,
	}
}

func mountPostgresDsn(cfg *DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBDatabase, cfg.DBPort, cfg.DBTimezone,
	)
}

func New(cfg *config.Config) *gorm.DB {
	con := Open(cfg)
	Migrate(con)
	return con
}

func Open(cfg *config.Config) *gorm.DB {
	var dbCfg *DatabaseConfig = getDatabaseConfig(cfg)
	var dsn string = mountPostgresDsn(dbCfg)
	con, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		NowFunc: func() time.Time {
			return time.Now().In(time.Local)
		},
	})
	if err != nil {
		log.Fatalf("%v: %v\n", ErrFailedToOpenConnectionPostgres, err)
	}
	return con
}
