package resources

type PlatformAnalytics struct {
	SuccesfulPurchases int64   `json:"sucessful_purchases"`
	ActiveOffers       int64   `json:"active_offers"`
	MoneyTransacted    float64 `json:"money_transacted"`
	EnergyTransacted   float64 `json:"energy_transacted"`
}

type UserAnalytics struct {
	*BuyerInfo    `json:"buyer_info"`
	*SupplierInfo `json:"supplier_info,omitempty"`
}

type SupplierInfo struct {
	MoneyEarned          float64 `json:"money_earned"`
	PurchasesCount       int64   `json:"purchases_count"`
	ActiveOffers         int64   `json:"active_offers"`
	AlmostExpiringOffers int64   `json:"almost_expiring_offers"`
	UserPriceAvg         float64 `json:"user_price_avg"`
	PlatformPriceAvg     float64 `json:"platform_price_avg"`
}

type BuyerInfo struct {
	PurchasesCount     int64   `json:"purchases_count"`
	EnergyTransacted   float64 `json:"energy_transacted"`
	AdvantageOfferUuid string  `json:"advantage_offer_uuid"`
}
