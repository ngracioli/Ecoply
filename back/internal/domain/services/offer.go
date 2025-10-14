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

	startTime, endTime, err := parseAndValidateOfferPeriodFromRequest(request.PeriodStart, request.PeriodEnd)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, err)
	}

	if err := s.db.Preload("Agent", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, submarket_id")
	}).Find(user).Error; err != nil {
		mlog.Log("Failed to preload agent: " + err.Error())
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInternal)
	}

	params := &repository.OfferCreateParams{
		Uuid:                 NewUuidV7String(),
		PricePerMwh:          request.PricePerMwh,
		InitialQuantityMwh:   request.QuantityMwh,
		RemainingQuantityMwh: request.QuantityMwh,
		Description:          request.Description,
		PeriodStart:          *startTime,
		PeriodEnd:            *endTime,
		Status:               models.OfferStatusFresh,
		EnergyTypeId:         energyType.ID,
		SellerId:             user.ID,
		SubmarketId:          user.Agent.SubmarketId,
	}

	offer, err := s.offerRepo.Create(params)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if err := s.db.Preload("Submarket", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(offer).Error; err != nil {
		mlog.Log("Failed to preload submarket: " + err.Error())
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var response *resources.Offer = &resources.Offer{
		Uuid:                 offer.Uuid,
		PricePerMwh:          offer.PricePerMwh,
		InitialQuantityMwh:   offer.InitialQuantityMwh,
		RemainingQuantityMwh: offer.RemainingQuantityMwh,
		Description:          offer.Description,
		PeriodStart:          offer.PeriodStart,
		PeriodEnd:            offer.PeriodEnd,
		Status:               offer.Status,
		EnergyType:           energyType.Type,
		Submarket:            offer.Submarket.Name,
		SellerUuid:           user.Uuid,
		CreatedAt:            offer.CreatedAt,
	}

	return response, nil
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

func parseOfferPeriodFromRequest(startPeriod string, endPeriod string) (*time.Time, *time.Time, error) {
	const layout = "2006-01-02"

	startTime, err := time.ParseInLocation(layout, startPeriod, time.Local)
	if err != nil {
		return nil, nil, ErrInvalidPeriodStart
	}

	endTime, err := time.ParseInLocation(layout, endPeriod, time.Local)
	if err != nil {
		return nil, nil, ErrInvalidPeriodEnd
	}

	return &startTime, &endTime, nil
}

func validateOfferPeriodFromRequest(startPeriod time.Time, endPeriod time.Time) error {
	now := time.Now()
	nowDateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	if startPeriod.After(endPeriod) || startPeriod.Before(nowDateOnly) {
		return ErrInvalidPeriod
	}

	return nil
}

func parseAndValidateOfferPeriodFromRequest(startPeriod string, endPeriod string) (*time.Time, *time.Time, error) {
	startTime, endTime, err := parseOfferPeriodFromRequest(startPeriod, endPeriod)
	if err != nil {
		return nil, nil, err
	}

	if err := validateOfferPeriodFromRequest(*startTime, *endTime); err != nil {
		return nil, nil, err
	}

	return startTime, endTime, nil
}

func makeOfferResourceFromModel(offer *models.Offer) *resources.Offer {
	var response resources.Offer = resources.Offer{
		Uuid:                 offer.Uuid,
		PricePerMwh:          offer.PricePerMwh,
		InitialQuantityMwh:   offer.InitialQuantityMwh,
		RemainingQuantityMwh: offer.RemainingQuantityMwh,
		Description:          offer.Description,
		PeriodStart:          offer.PeriodStart,
		PeriodEnd:            offer.PeriodEnd,
		Status:               offer.Status,
		EnergyType:           offer.EnergyType.Type,
		Submarket:            offer.Submarket.Name,
		SellerUuid:           offer.Seller.Uuid,
		CreatedAt:            offer.CreatedAt,
	}

	return &response
}
