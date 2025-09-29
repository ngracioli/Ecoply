package repository

// import (
// 	"ecoply/internal/domain/models"

// 	"gorm.io/gorm"
// )

// type OfferRepository interface {
// 	GetOfferByUuid(id string) (*models.Offer, error)
// 	GetUserOffers(userUuid string) ([]*models.Offer, error)
// 	CreateOffer(param OfferCreateParams) error
// 	UpdateOffer(offer *models.Offer) error
// 	DeleteOffer(id string) error
// }

// type offerRepository struct {
// 	db *gorm.DB
// }

// func NewOfferRepository(db *gorm.DB) OfferRepository {
// 	return &offerRepository{db: db}
// }

// func (r *offerRepository) GetOfferByID(id string) (*models.Offer, error) {
// 	var offer models.Offer
// 	if err := r.db.First(&offer, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &offer, nil
// }

// func (r *offerRepository) GetUserOffers(userID string) ([]*models.Offer, error) {
// 	var offers []*models.Offer
// 	if err := r.db.Where("user_id = ?", userID).Find(&offers).Error; err != nil {
// 		return nil, err
// 	}
// 	return offers, nil
// }

// func (r *offerRepository) CreateOffer(offer *models.Offer) error {
// 	return r.db.Create(offer).Error
// }

// func (r *offerRepository) UpdateOffer(offer *models.Offer) error {
// 	return r.db.Save(offer).Error
// }

// func (r *offerRepository) DeleteOffer(id string) error {
// 	return r.db.Delete(&models.Offer{}, id).Error
// }
