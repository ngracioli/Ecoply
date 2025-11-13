package resources

type Purchase struct {
	Uuid          string  `json:"uuid"`
	QuantityMwh   float64 `json:"quantity_mwh"`
	PricePerMwh   float64 `json:"price_per_mwh"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
	OfferUuid     string  `json:"offer_uuid"`
	SellerUuid    string  `json:"seller_uuid"`
	CreatedAt     string  `json:"created_at"`
}
