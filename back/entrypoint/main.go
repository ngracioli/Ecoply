package main

import (
	"ecoply/internal/config"
	"ecoply/internal/database"
	"ecoply/internal/domain/handlers"
	"ecoply/internal/domain/services"
	"ecoply/internal/domain/tasks"
	"ecoply/internal/domain/validation"
	"ecoply/internal/mlog"
	"ecoply/internal/server"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gorm.io/gorm"
)

func main() {
	mlog.CreateServerLogger()
	defer mlog.CloseLogFiles()

	validation.RegisterCustomValidators()

	var cfg *config.Config = loadEnvironment()

	setAppTimezone(cfg)

	db := database.New(cfg)

	serverContext := buildServerContext(cfg, db)

	tasks.RunBackgroundTasks(serverContext)

	server.NewAndRun(serverContext)
}

func buildServerContext(cfg *config.Config, db *gorm.DB) *server.ServerContext {
	services := server.ServerServices{
		AuthService:      services.NewAuthService(cfg, db),
		UserService:      services.NewUserService(db),
		OfferService:     services.NewOfferService(db),
		PurchaseService:  services.NewPurchaseService(db),
		UserTypeService:  services.NewUserTypeService(db),
		ContractService:  services.NewContractService(db),
		AnalyticsService: services.NewAnalyticsService(db),
		CceeService:      services.NewCceeService(),
		BrasilApiService: services.NewBrasilApiService(),
	}

	handlers := server.ServerHandlers{
		AuthHandlers:      handlers.NewAuthHandler(services.AuthService),
		CnpjHandlers:      handlers.NewCnpjHandler(),
		OfferHandlers:     handlers.NewOfferHandler(services.OfferService),
		PurchaseHandlers:  handlers.NewPurchaseHandlers(services.PurchaseService),
		ContractHandlers:  handlers.NewContractHandlers(services.ContractService),
		AnalyticsHandlers: handlers.NewAnalyticsHandler(services.AnalyticsService),
		CceeHandlers:      handlers.NewCceeHandler(services.CceeService),
		BrasilApiHandlers: handlers.NewBrasilApiHandler(services.BrasilApiService),
	}

	return &server.ServerContext{
		Cfg:      cfg,
		Handlers: handlers,
		Services: services,
	}
}

func setAppTimezone(cfg *config.Config) {
	loc, err := time.LoadLocation(cfg.DBTimezone)
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	time.Local = loc
}

func loadEnvironment() *config.Config {
	var envFilePath string
	var mode string = os.Getenv("APP_ENV")

	if mode != "prod" {
		_, filename, _, _ := runtime.Caller(0)
		var projectRoot string = filename

		// Going directories up to reach project root, where the .env file should be placed
		projectRoot = filepath.Dir(filepath.Dir(projectRoot))

		envFilePath = filepath.Join(projectRoot, ".env")
	}

	config, err := config.Load(envFilePath)
	if err != nil {
		panic(err)
	}

	return config
}
