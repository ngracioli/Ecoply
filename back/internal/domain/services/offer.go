package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/repository"
	"ecoply/internal/domain/requests"
	"ecoply/internal/domain/resources"
	"ecoply/internal/domain/utils"
	"ecoply/internal/mlog"
	"errors"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

type OfferService interface {
	GetByUuid(uuid string) (*resources.Offer, *merr.ResponseError)
	BelongingToUser(userId uint) ([]*resources.Offer, *merr.ResponseError)
	Create(user *models.User, request *requests.CreateOffer) (*resources.Offer, *merr.ResponseError)
	Update(user *models.User, uuid string, request *requests.UpdateOffer) *merr.ResponseError
	Delete(user *models.User, uuid string) *merr.ResponseError
	List(params *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*resources.Offer], *merr.ResponseError)
	Purchases(offerUuid string, request *requests.ListPurchasesFromOffer, user *models.User) ([]*resources.Purchase, *merr.ResponseError)
	UpdateExpiredOffers() error
}

type offerService struct {
	offerRepo      repository.OfferRepository
	submarketRepo  repository.SubmarketRepository
	userTypeRepo   repository.UserTypeRepository
	energyTypeRepo repository.EnergyTypeRepository
	db             *gorm.DB
}

func NewOfferService(db *gorm.DB) OfferService {
	return &offerService{
		offerRepo:      repository.NewOfferRepository(db),
		submarketRepo:  repository.NewSubmarketRepository(db),
		userTypeRepo:   repository.NewUserTypeRepository(db),
		energyTypeRepo: repository.NewEnergyRepository(db),
		db:             db,
	}
}

func (s *offerService) Create(user *models.User, request *requests.CreateOffer) (*resources.Offer, *merr.ResponseError) {
	var energyType *models.EnergyType
	var err error

	energyType, err = s.energyTypeRepo.GetByType(request.EnergyType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidEnergyType)
	} else if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if err = validateCreateRequest(request); err != nil {
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, err)
	}

	if err = s.db.Preload("Agent", func(db *gorm.DB) *gorm.DB { return db.Select("id, submarket_id") }).
		Find(user).Error; err != nil {
		mlog.Log("Failed to preload agent: " + err.Error())
		return nil, merr.NewResponseError(http.StatusUnprocessableEntity, ErrInternal)
	}

	parsedStartPeriod, _ := parseDate(request.PeriodStart)
	parsedEndPeriod, _ := parseDate(request.PeriodEnd)

	offer := &models.Offer{
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

	offer, err = s.offerRepo.Create(offer)
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

func (s *offerService) Update(user *models.User, uuid string, request *requests.UpdateOffer) *merr.ResponseError {
	var energyType *models.EnergyType
	var offer *models.Offer
	var err error
	var periodStart time.Time
	var periodEnd time.Time

	offer, err = s.offerRepo.GetByUuid(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
	}

	if offer.Status != models.OfferStatusFresh {
		return merr.NewResponseError(http.StatusUnprocessableEntity, ErrCannotDeleteOffer)
	}

	if offer.SellerId != user.ID {
		return merr.NewResponseError(http.StatusForbidden, ErrUserIsNotTheOfferOwner)
	}

	energyType, err = s.energyTypeRepo.GetByType(request.EnergyType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.NewResponseError(http.StatusUnprocessableEntity, ErrInvalidEnergyType)
	} else if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = validateUpdateRequest(offer, request)
	if err != nil {
		return merr.NewResponseError(http.StatusUnprocessableEntity, err)
	}

	periodStart, _ = parseDate(request.PeriodStart)
	periodEnd, _ = parseDate(request.PeriodEnd)

	offer.Description = request.Description
	offer.EnergyTypeId = energyType.ID
	offer.EnergyType = *energyType
	offer.InitialQuantityMwh = request.QuantityMwh
	offer.RemainingQuantityMwh = request.QuantityMwh
	offer.PricePerMwh = request.PricePerMwh
	offer.PeriodStart = periodStart
	offer.PeriodEnd = periodEnd

	err = s.offerRepo.Update(offer)
	if err != nil {
		return merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return nil
}

func (s *offerService) Delete(user *models.User, uuid string) *merr.ResponseError {
	var offer *models.Offer
	var err error

	offer, err = s.offerRepo.GetByUuid(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
	}

	if offer.Status != models.OfferStatusFresh {
		return merr.NewResponseError(http.StatusUnprocessableEntity, ErrCannotDeleteOffer)
	}

	if offer.SellerId != user.ID {
		return merr.NewResponseError(http.StatusForbidden, ErrUserIsNotTheOfferOwner)
	}

	err = s.offerRepo.Delete(uuid)
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

func validateCreatePeriodFromRequest(startPeriod string, endPeriod string) error {
	var nowZeroHour = utils.NowInLocalZeroHour()

	parsedStartPeriod, err := parseDate(startPeriod)
	if err != nil {
		return ErrInvalidPeriod
	}

	parsedEndPeriod, err := parseDate(endPeriod)
	if err != nil {
		return ErrInvalidPeriod
	}

	if parsedStartPeriod.After(parsedEndPeriod) || parsedStartPeriod.Before(nowZeroHour) {
		return ErrInvalidPeriod
	}

	return nil
}

func validateUpdatePeriodFromRequest(offer *models.Offer, request *requests.UpdateOffer) error {
	var err error
	var parsedStartPeriod time.Time
	var parsedEndPeriod time.Time

	parsedStartPeriod, err = parseDate(request.PeriodStart)
	if err != nil {
		return ErrInvalidPeriod
	}

	parsedEndPeriod, err = parseDate(request.PeriodEnd)
	if err != nil {
		return ErrInvalidPeriod
	}

	var nowZeroHour time.Time = utils.NowInLocalZeroHour()
	var offerStartTruncated time.Time = utils.TruncateDateToLocalZeroHour(offer.PeriodStart)
	var offerEndTruncated time.Time = utils.TruncateDateToLocalZeroHour(offer.PeriodEnd)

	if !parsedStartPeriod.Equal(offerStartTruncated) {
		if parsedStartPeriod.After(parsedEndPeriod) || parsedStartPeriod.Before(nowZeroHour) {
			return ErrInvalidPeriod
		}
	}

	if !parsedEndPeriod.Equal(offerEndTruncated) {
		if parsedEndPeriod.Before(parsedStartPeriod) || parsedEndPeriod.Before(nowZeroHour) {
			return ErrInvalidPeriod
		}
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

	if err := validateCreatePeriodFromRequest(request.PeriodStart, request.PeriodEnd); err != nil {
		return err
	}

	return nil
}

func validateUpdateRequest(offer *models.Offer, request *requests.UpdateOffer) error {
	if err := validatePrice(request.PricePerMwh); err != nil {
		return err
	}

	if err := validateQuantity(request.QuantityMwh); err != nil {
		return err
	}

	if err := validateUpdatePeriodFromRequest(offer, request); err != nil {
		return err
	}

	return nil
}

func makeOfferResourceFromModel(offer *models.Offer) *resources.Offer {
	var createdAt time.Time = utils.TruncateDateToLocal(offer.CreatedAt)
	var response resources.Offer = resources.Offer{
		Uuid:                 offer.Uuid,
		PricePerMwh:          offer.PricePerMwh,
		InitialQuantityMwh:   offer.InitialQuantityMwh,
		RemainingQuantityMwh: offer.RemainingQuantityMwh,
		Description:          offer.Description,
		PeriodStart:          offer.PeriodStart.Format(time.DateOnly),
		PeriodEnd:            offer.PeriodEnd.Format(time.DateOnly),
		Status:               offer.Status,
		EnergyType:           offer.EnergyType.Type,
		Submarket:            offer.Submarket.Name,
		SellerUuid:           offer.Seller.Uuid,
		CreatedAt:            createdAt,
	}

	return &response
}

func (s *offerService) List(request *requests.ListOffers, user *models.User) (*utils.PaginationWrapper[*resources.Offer], *merr.ResponseError) {
	var list *utils.PaginationWrapper[*models.Offer]
	var err error

	list, err = s.offerRepo.List(request, user)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var response utils.PaginationWrapper[*resources.Offer]

	response.Page = list.Page
	response.PageSize = list.PageSize
	response.HasNext = list.HasNext
	response.HasPrev = list.HasPrev
	response.Data = make([]*resources.Offer, 0, len(list.Data))

	for _, offer := range list.Data {
		response.Data = append(response.Data, makeOfferResourceFromModel(offer))
	}

	return &response, nil
}

func (s *offerService) UpdateExpiredOffers() error {
	return s.offerRepo.UpdateExpiredOffers()
}

func (s *offerService) Purchases(offerUuid string, request *requests.ListPurchasesFromOffer, user *models.User) ([]*resources.Purchase, *merr.ResponseError) {
	var offer *models.Offer
	var err error

	offer, err = s.offerRepo.GetByUuid(offerUuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, merr.NewResponseError(http.StatusNotFound, ErrOfferNotFound)
	} else if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if !offer.IsOwner(user) {
		return nil, merr.NewResponseError(http.StatusForbidden, ErrUserIsNotTheOfferOwner)
	}

	purchases, err := s.offerRepo.Purchases(offerUuid, request)
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	response := make([]*resources.Purchase, 0, len(purchases))

	for _, purchase := range purchases {
		var createdAt time.Time = utils.TruncateDateToLocal(purchase.CreatedAt)
		response = append(response, &resources.Purchase{
			Uuid:          purchase.Uuid,
			QuantityMwh:   purchase.QuantityMwh,
			PricePerMwh:   purchase.PricePerMwh,
			Status:        purchase.Status,
			PaymentMethod: purchase.PaymentMethod,
			OfferUuid:     purchase.Offer.Uuid,
			SellerUuid:    purchase.Offer.Seller.Uuid,
			BuyerName:     purchase.Buyer.Name,
			SellerName:    purchase.Offer.Seller.Name,
			CreatedAt:     createdAt.Format(time.RFC3339),
		})
	}

	return response, nil
}
