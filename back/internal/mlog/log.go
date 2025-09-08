package mlog

import (
	"ecoply/internal/config"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var serverLogFile *os.File

func CreateServerLogger() {
	ensureLogDirExists()

	file, err := os.OpenFile("log/server.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("Failed to open server.log: %v\n", err)
		return
	}

	serverLogFile = file
}

func getServerLogger() io.Writer {
	if serverLogFile == nil {
		CreateServerLogger()
	}

	return serverLogFile
}

func log(w io.Writer, message string) {
	if config.IsProduction() {
		return
	}

	var str string = fmt.Sprintf("%v %s\n", time.Now().Format(time.RFC3339), message)
	w.Write([]byte(str))
}

func Log(message string) {
	log(getServerLogger(), message)
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

func CloseLogFiles() {
	if serverLogFile != nil {
		serverLogFile.Close()
	}
}

func ensureLogDirExists() {
	_, err := os.Stat("log")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("log", 0766)
		} else {
			fmt.Printf("Failed to check log dir stat: %v\n", err)
		}
	}
}
