package main

import (
	"ecoply/internal/config"
	"ecoply/internal/database"
	"ecoply/internal/domain/validation"
	"ecoply/internal/mlog"
	"ecoply/internal/server"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	loadEnvironment()

	validation.RegisterCustomValidators()

	database.New()

	mlog.CreateServerLogger()
	defer mlog.CloseLogFiles()

	server.NewAndRun()
}

func loadEnvironment() {
	var envFilePath string
	var mode string = os.Getenv("APP_ENV")

	if mode != "prod" {
		_, filename, _, _ := runtime.Caller(0)
		var projectRoot string = filename

		// Going directories up to reach project root, where the .env file should be placed
		projectRoot = filepath.Dir(filepath.Dir(projectRoot))

		envFilePath = filepath.Join(projectRoot, ".env")
	}

	if err := config.Load(envFilePath); err != nil {
		log.Fatalf("%v\n", err)
	}
}
