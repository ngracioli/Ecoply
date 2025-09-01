package mlog

import (
	"ecoply/internal/config"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Log(message string) {
	if config.IsProduction() {
		return
	}

	ensureLogDirExists()

	var fileName = "server.log"

	file, err := os.OpenFile("log/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Failed to open %s: %v\n", fileName, err)
	}

	defer file.Close()

	var str string = fmt.Sprintf("%v %s\n", time.Now().String(), message)
	file.WriteString(str)
}

func LogGinRoutes(engine *gin.Engine) {
	if config.IsProduction() {
		return
	}

	ensureLogDirExists()

	file, err := os.OpenFile("log/routes.log", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Failed to open routes.log: %v\n", err)
	}

	defer file.Close()

	routes := engine.Routes()
	for _, route := range routes {
		fmt.Fprintf(file, "%-8s %-30s %s\n", route.Method, route.Path, route.Handler)
	}
}

func ensureLogDirExists() {
	_, err := os.Stat("log")
	if err != nil {
		switch true {
		case os.IsNotExist(err):
			os.Mkdir("log", 0766)
		case false:
			fmt.Printf("Failed to check log dir stat: %v\n", err)
		}
	}
}
