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

var commands map[string]byte = map[string]byte{
	"migrate": 1,
	"serve":   1,
}

func main() {
	command := getCommand()

	loadEnvironment()

	mlog.CreateServerLogger()
	defer mlog.CloseLogFiles()

	runCommand(command)
}

func runCommand(command string) {
	switch command {
	case "migrate":
		runMigrations()
	case "serve":
		runServer()
	default:
	}
}

func runMigrations() {
	database.New()
}

func runServer() {
	database.New()
	validation.RegisterCustomValidators()
	server.NewAndRun()
}

func getCommand() string {
	if len(os.Args) <= 1 {
		return ""
	}

	if commandExists(os.Args[1]) {
		return os.Args[1]
	}

	return ""
}

func commandExists(command string) bool {
	_, exists := commands[command]
	return exists
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
