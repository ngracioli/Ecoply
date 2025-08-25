package main

import (
	"ecoply/internal/config"
	"ecoply/internal/server"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	loadEnvironment()

	var cfg *config.Config = config.GetConfig()
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBDatabase, cfg.DBPort, cfg.DBTimezone,
	)
	con, err := gorm.Open(postgres.Open(dsn))
	fmt.Printf("%v\n", err)

	db, err := con.DB()
	fmt.Printf("%v\n", err)

	var queries []string = []string{
		"CREATE TABLE IF NOT EXISTS users (id BIGINT GENERATED ALWAYS AS IDENTITY, name VARCHAR(255) NOT NULL)",
		"INSERT INTO users (name) VALUES ('BANANA')",
		"SELECT * FROM users",
	}

	for _, q := range queries {
		rows, err := db.Query(q)
		if err != nil {
			fmt.Printf("Query error: %v\n", err)
			continue
		}
		defer rows.Close()

		if rows.Next() {
			var user struct {
				Id   uint64
				Name string
			}
			err := rows.Scan(&user.Id, &user.Name)
			if err != nil {
				fmt.Printf("Scan error: %v\n", err)
			} else {
				fmt.Printf("User: %+v\n", user)
			}
		}
	}

	os.Exit(0)

	server.NewAndRun()
}

func loadEnvironment() {
	var envFiles []string = make([]string, 0, 3)
	mode := os.Getenv("APP_ENV")

	if mode != "" && mode != "prod" {
		_, filename, _, _ := runtime.Caller(0)
		var projectRoot string = filename

		// Going directories up to reach project root, where the .env file should be placed
		projectRoot = filepath.Dir(filepath.Dir(projectRoot))

		envFiles = append(envFiles, filepath.Join(projectRoot, ".env"))
	}

	if err := config.Load(envFiles...); err != nil {
		log.Fatalf("%v\n", err)
	}
}
