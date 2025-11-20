package tasks

import (
	"context"
	"ecoply/internal/background"
	"ecoply/internal/domain/services"
	"ecoply/internal/server"
	"time"
)

func RunBackgroundTasks(s *server.ServerContext) {
	updateOfferStatusToExpired(s.Services.OfferService)
}

func updateOfferStatusToExpired(service services.OfferService) {
	var ctx context.Context = context.Background()

	background.StartPeriodicTask(ctx, time.Duration(time.Minute*10), func() error {
		return service.UpdateExpiredOffers()
	})
}
