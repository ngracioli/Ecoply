package main

import (
	"ecoply/internal/config"
	"ecoply/internal/database"
	"ecoply/internal/domain/services"
	"ecoply/internal/domain/validation"
	"ecoply/internal/mlog"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gorm.io/gorm"
)

func main() {
	var cfg *config.Config = loadEnvironment()

	loc, err := time.LoadLocation(cfg.DBTimezone)
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}

	time.Local = loc

	validation.RegisterCustomValidators()

	var db *gorm.DB = database.New()

	services.InitServices(db)

	offerUpdateCh := launchExpiredOffersUpdaterJob()

	mlog.CreateServerLogger()
	defer mlog.CloseLogFiles()

	offerUpdateCh <- true
	close(offerUpdateCh)
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

	if err := config.Load(envFilePath); err != nil {
		log.Fatalf("%v\n", err)
	}

	return config.GetConfig()
}

// Must be called after services initialization
func launchExpiredOffersUpdaterJob() chan bool {
	offerService := services.Offer

	ticker := time.NewTicker(5 * time.Minute)
	closeChannel := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				err := offerService.UpdateExpiredOffers()
				if err != nil {
					mlog.Log("Error updating expired offers: " + err.Error())
				}
			case <-closeChannel:
				ticker.Stop()
				fmt.Printf("Killed\n")
				return
			}
		}
	}()

	return closeChannel
}
