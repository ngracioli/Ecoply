package main

import (
	"ecoply/internal/config"
	"ecoply/internal/server"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	loadEnvironment()
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
