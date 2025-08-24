package server

import (
	"ecoply/internal/config"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Host   string
	Port   uint16
}

func New() *Server {
	var cfg *config.Config = config.GetConfig()

	var engine *gin.Engine = gin.Default(
		registerRoutes,
	)

	storeEngineRoutesInFile(engine)

	return &Server{
		Engine: engine,
		Host:   cfg.ServerHost,
		Port:   cfg.ServerPort,
	}
}

func NewAndRun() *Server {
	var server *Server = New()
	server.Run()
	return server
}

func storeEngineRoutesInFile(engine *gin.Engine) {
	if config.IsProduction() {
		return
	}

	ensureLogDirExists()

	file, err := os.OpenFile("log/routes.log", os.O_WRONLY|os.O_CREATE, 0644)
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
			log.Fatalf("Failed to check log dir stat: %v\n", err)
		}
	}
}

func (s *Server) Run() {
	var address string = net.JoinHostPort(
		s.Host,
		strconv.FormatUint(uint64(s.Port), 10),
	)

	s.Engine.Run(address)
}
