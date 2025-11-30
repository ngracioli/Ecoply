package server

import (
	"ecoply/internal/config"
	"ecoply/internal/domain/handlers"
	"ecoply/internal/domain/services"
	"ecoply/internal/mlog"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Host   string
	Port   uint16
}

type ServerServices struct {
	services.AuthService
	services.UserService
	services.OfferService
	services.PurchaseService
	services.UserTypeService
	services.ContractService
	services.AnalyticsService
	services.CceeService
}

type ServerHandlers struct {
	handlers.AuthHandlers
	handlers.CnpjHandlers
	handlers.OfferHandlers
	handlers.PurchaseHandlers
	handlers.ContractHandlers
	handlers.AnalyticsHandlers
	handlers.CceeHandlers
}

type ServerContext struct {
	Cfg      *config.Config
	Services ServerServices
	Handlers ServerHandlers
}

func New(s *ServerContext) *Server {
	var engine *gin.Engine = gin.Default()

	registerRoutes(engine, s)

	mlog.LogGinRoutes(engine)

	return &Server{
		Engine: engine,
		Host:   s.Cfg.ServerHost,
		Port:   s.Cfg.ServerPort,
	}
}

func NewAndRun(s *ServerContext) *Server {
	var server *Server = New(s)
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
