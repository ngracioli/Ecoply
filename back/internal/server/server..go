package server

import (
	"ecoply/internal/config"
	"net"
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

	var engine gin.Engine = *gin.Default(
		registerRoutes,
	)

	return &Server{
		Engine: &engine,
		Host:   cfg.ServerHost,
		Port:   cfg.ServerPort,
	}
}

func NewAndRun() *Server {
	var server *Server = New()
	server.Run()
	return server
}

func (s *Server) Run() {
	var address string = net.JoinHostPort(
		s.Host,
		strconv.FormatUint(uint64(s.Port), 10),
	)

	s.Engine.Run(address)
}
