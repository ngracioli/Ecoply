package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/resources"
	"ecoply/internal/mlog"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

type OfferService interface {
	GetByUuid(uuid string) (*resources.Offer, *merr.ResponseError)
	BelongingToUser(userId uint) ([]*resources.Offer, *merr.ResponseError)
	Create(user *models.User, request *requests.CreateOffer) (*resources.Offer, *merr.ResponseError)
	Update(offer *models.Offer) *merr.ResponseError
	Delete(id string) *merr.ResponseError
}

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

func (s *offerService) Create(user *models.User, request *requests.CreateOffer) (*resources.Offer, *merr.ResponseError) {
	var energyType *models.EnergyType
	if err := s.db.Where("type = ?", request.EnergyType).First(&energyType).Error; err != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidEnergyType)
	}

	if err := validateCreateRequest(request); err != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, err)
	}

	if err := s.db.Preload("Agent", func(db *gorm.DB) *gorm.DB { return db.Select("id, submarket_id") }).
		Find(user).Error; err != nil {
		mlog.Log("Failed to preload agent: " + err.Error())
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInternal)
	}

	parsedStartPeriod, _ := parseDate(request.PeriodStart)
	parsedEndPeriod, _ := parseDate(request.PeriodEnd)

	params := &repository.OfferCreateParams{
		Uuid:                 NewUuidV7String(),
		PricePerMwh:          request.PricePerMwh,
		InitialQuantityMwh:   request.QuantityMwh,
		RemainingQuantityMwh: request.QuantityMwh,
		Description:          request.Description,
		PeriodStart:          parsedStartPeriod,
		PeriodEnd:            parsedEndPeriod,
		Status:               models.OfferStatusFresh,
		EnergyTypeId:         energyType.ID,
		SellerId:             user.ID,
		SubmarketId:          user.Agent.SubmarketId,
	}

	offer, err := s.offerRepo.Create(params)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if err := s.db.
		Preload("EnergyType", func(db *gorm.DB) *gorm.DB { return db.Select("id", "type") }).
		Preload("Submarket", func(db *gorm.DB) *gorm.DB { return db.Select("id", "name") }).
		Preload("Seller", func(db *gorm.DB) *gorm.DB { return db.Select("id", "uuid") }).
		First(offer, offer.ID).Error; err != nil {
		mlog.Log("Failed to preload offer relations: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return makeOfferResourceFromModel(offer), nil
}

func (s *offerService) Update(offer *models.Offer) *merr.ResponseError {
	err := s.offerRepo.Update(offer)
	if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return nil
}

func (s *offerService) Delete(id string) *merr.ResponseError {
	err := s.offerRepo.Delete(id)
	if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return nil
}

func (s *offerService) GetByUuid(uuid string) (*resources.Offer, *merr.ResponseError) {
	offer, err := s.offerRepo.GetByUuid(strings.ToLower(uuid))
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrOfferNotFound)
	}

	response := makeOfferResourceFromModel(offer)

	return response, nil
}

func (s *offerService) BelongingToUser(userId uint) ([]*resources.Offer, *merr.ResponseError) {
	offers, err := s.offerRepo.GetBySellerId(userId)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	response := make([]*resources.Offer, len(offers))
	for i, offer := range offers {
		response[i] = makeOfferResourceFromModel(offer)
	}

	return response, nil
}

func parseDate(date string) (time.Time, error) {
	layout := "2006-01-02"
	return time.ParseInLocation(layout, date, time.Local)
}

func validatePeriodFromRequest(startPeriod string, endPeriod string) error {
	nowInLoc := time.Now().In(time.Local)
	now := time.Date(
        nowInLoc.Year(),
        nowInLoc.Month(),
        nowInLoc.Day(),
        0, 0, 0, 0,
        time.Local,
    )

	parsedStartPeriod, err := parseDate(startPeriod)
	if err != nil {
		return ErrInvalidPeriod
	}

	parsedEndPeriod, err := parseDate(endPeriod)
	if err != nil {
		return ErrInvalidPeriod
	}

	if parsedStartPeriod.After(parsedEndPeriod) || parsedStartPeriod.Before(now) {
		return ErrInvalidPeriod
	}

	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}

func validateQuantity(quantity float64) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	return nil
}

func validateCreateRequest(request *requests.CreateOffer) error {
	if err := validatePrice(request.PricePerMwh); err != nil {
		return err
	}

	if err := validateQuantity(request.QuantityMwh); err != nil {
		return err
	}

	if err := validatePeriodFromRequest(request.PeriodStart, request.PeriodEnd); err != nil {
		return err
	}

	return nil
}

func makeOfferResourceFromModel(offer *models.Offer) *resources.Offer {
	var response resources.Offer = resources.Offer{
		Uuid:                 offer.Uuid,
		PricePerMwh:          offer.PricePerMwh,
		InitialQuantityMwh:   offer.InitialQuantityMwh,
		RemainingQuantityMwh: offer.RemainingQuantityMwh,
		Description:          offer.Description,
		PeriodStart:          offer.PeriodStart.Format("2006-01-02"),
		PeriodEnd:            offer.PeriodEnd.Format("2006-01-02"),
		Status:               offer.Status,
		EnergyType:           offer.EnergyType.Type,
		Submarket:            offer.Submarket.Name,
		SellerUuid:           offer.Seller.Uuid,
		CreatedAt:            offer.CreatedAt,
	}

	return &response
}
