package services

import (
	"ecoply/internal/domain/repository"

	"gorm.io/gorm"
)

type OfferService interface{}

type offerService struct {
	offerRepo repository.OfferRepository
	db        *gorm.DB
}

func NewOfferService(db *gorm.DB) OfferService {
	return &offerService{
		offerRepo: repository.NewOfferRepository(db),
		db:        db,
	}
}
