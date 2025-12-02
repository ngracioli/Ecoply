package services

import (
	"ecoply/internal/domain/merr"
	"ecoply/internal/domain/models"
	"ecoply/internal/domain/resources"
	"ecoply/internal/domain/utils"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type AnalyticsService interface {
	Platform() (*resources.PlatformAnalytics, *merr.ResponseError)
	User(user *models.User) (*resources.UserAnalytics, *merr.ResponseError)
}

type analyticsService struct {
	db *gorm.DB
}

func NewAnalyticsService(db *gorm.DB) AnalyticsService {
	return &analyticsService{
		db: db,
	}
}

func (s *analyticsService) Platform() (*resources.PlatformAnalytics, *merr.ResponseError) {
	var platformAnalytics resources.PlatformAnalytics
	var err error

	err = s.db.Model(&models.Purchase{}).
		Where("status = ?", models.PurchaseStatusCompleted).
		Count(&platformAnalytics.SuccesfulPurchases).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Offer{}).
		Where("status IN ?", []string{models.OfferStatusOpen, models.OfferStatusFresh}).
		Count(&platformAnalytics.ActiveOffers).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Purchase{}).
		Where("status = ?", models.PurchaseStatusCompleted).
		Select("COALESCE(SUM(quantity_mwh * price_per_mwh), 0) AS total").
		Scan(&platformAnalytics.MoneyTransacted).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Purchase{}).
		Where("status = ?", models.PurchaseStatusCompleted).
		Select("COALESCE(SUM(quantity_mwh), 0) AS total").
		Scan(&platformAnalytics.EnergyTransacted).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	return &platformAnalytics, nil
}

func (s *analyticsService) User(user *models.User) (*resources.UserAnalytics, *merr.ResponseError) {
	buyerInfo, err := s.makeBuyerInfo(user)
	if err != nil {
		return nil, err
	}

	supplierInfo, err := s.makeSupplierInfo(user)
	if err != nil {
		return nil, err
	}

	response := resources.UserAnalytics{
		BuyerInfo:    buyerInfo,
		SupplierInfo: supplierInfo,
	}

	return &response, nil
}

func (s *analyticsService) makeSupplierInfo(user *models.User) (*resources.SupplierInfo, *merr.ResponseError) {
	var moneyEarned float64
	var purchasesCount int64
	var activeOffers int64
	var almostExpiringOffers int64
	var userPriceAvg float64
	var platformPriceAvg float64
	var err error

	err = s.db.Preload("UserType").Find(user).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	if user.UserType.Type == models.UserTypeBuyer {
		return nil, nil
	}

	err = s.db.
		Model(&models.Purchase{}).
		Joins("JOIN offers ON offers.id = purchases.offer_id").
		Where("offers.seller_id = ?", user.ID).
		Where("purchases.status = ?", models.PurchaseStatusCompleted).
		Select("COALESCE(SUM(purchases.quantity_mwh * purchases.price_per_mwh), 0) AS total").
		Scan(&moneyEarned).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.
		Model(&models.Purchase{}).
		Joins("JOIN offers ON offers.id = purchases.offer_id").
		Where("offers.seller_id = ?", user.ID).
		Where("purchases.status = ?", models.PurchaseStatusCompleted).
		Count(&purchasesCount).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Offer{}).Where("seller_id = ?", user.ID).Where("status IN ?", []string{models.OfferStatusOpen, models.OfferStatusFresh}).Count(&activeOffers).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var almostExpiringThreshold time.Time = utils.NowInLocal().Add(2 * 24 * time.Hour)
	err = s.db.Model(&models.Offer{}).Where("seller_id = ?", user.ID).
		Where("status = ?", models.OfferStatusOpen).
		Where("period_end < ?", almostExpiringThreshold).
		Count(&almostExpiringOffers).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Offer{}).Where("seller_id = ?", user.ID).Select("COALESCE(AVG(price_per_mwh), 0) AS avg_price").Scan(&userPriceAvg).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Offer{}).Where("seller_id != ?", user.ID).Select("COALESCE(AVG(price_per_mwh), 0) AS avg_price").Scan(&platformPriceAvg).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var supplierInfo resources.SupplierInfo = resources.SupplierInfo{
		MoneyEarned:          moneyEarned,
		PurchasesCount:       purchasesCount,
		ActiveOffers:         activeOffers,
		AlmostExpiringOffers: almostExpiringOffers,
		UserPriceAvg:         userPriceAvg,
		PlatformPriceAvg:     platformPriceAvg,
	}

	return &supplierInfo, nil
}

func (s *analyticsService) makeBuyerInfo(user *models.User) (*resources.BuyerInfo, *merr.ResponseError) {
	var purchasesCount int64
	var energyTransacted float64
	var advantageOfferUuid string
	var err error

	err = s.db.Preload("Agent").Preload("Agent.Submarket").Find(user).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Purchase{}).
		Where("buyer_id = ?", user.ID).
		Where("status = ?", models.PurchaseStatusCompleted).
		Count(&purchasesCount).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Purchase{}).
		Where("buyer_id = ?", user.ID).
		Where("status = ?", models.PurchaseStatusCompleted).
		Select("COALESCE(SUM(quantity_mwh), 0) AS total").
		Scan(&energyTransacted).Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	err = s.db.Model(&models.Offer{}).
		Select("uuid").
		Where("submarket_id = ?", user.Agent.SubmarketId).
		Where("seller_id != ?", user.ID).
		Where("status IN (?)", []string{
			models.OfferStatusFresh,
			models.OfferStatusOpen,
		}).
		Order("price_per_mwh ASC").
		Limit(1).
		Scan(&advantageOfferUuid).
		Error
	if err != nil {
		return nil, merr.NewResponseError(http.StatusInternalServerError, ErrInternal)
	}

	var buyerInfo resources.BuyerInfo = resources.BuyerInfo{
		PurchasesCount:     purchasesCount,
		EnergyTransacted:   energyTransacted,
		AdvantageOfferUuid: advantageOfferUuid,
	}

	return &buyerInfo, nil
}
