package main

import (
	"ecoply/internal/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

func main() {
	loadEnvironment()
	fields := reflect.ValueOf(*config.GetConfig())
	fieldTypes := fields.Type()

	var totalFields int = fields.NumField()
	for i := 0; i < totalFields; i++ {
		fmt.Printf(
			"%v=%v - %v\n",
			fieldTypes.Field(i).Name,
			fields.Field(i).Interface(),
			fieldTypes.Field(i).Tag,
		)
	}
}

func loadEnvironment() {
	var envFiles []string = make([]string, 0, 3)
	mode := os.Getenv("APP_ENV")

	if mode != "prod" {
		_, filename, _, _ := runtime.Caller(0)
		var projectRoot string = filename
		for filepath.Base(projectRoot) != "back" {
			projectRoot = filepath.Dir(projectRoot)
		}
		envFiles = append(envFiles, filepath.Join(projectRoot, ".env"))
	}

	if err := config.Load(envFiles...); err != nil {
		log.Fatalf("%v\n", err)
	}
}
