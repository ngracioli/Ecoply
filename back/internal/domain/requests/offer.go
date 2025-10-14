package requests

type GetOffer struct {
	Uuid string `form:"uuid" binding:"required"`
}

type CreateOffer struct {
	PricePerMwh float64 `json:"price_per_mwh" binding:"required"`
	QuantityMwh float64 `json:"quantity_mwh" binding:"required"`
	PeriodStart string  `json:"period_start" binding:"required,datetime=2006-01-02"`
	PeriodEnd   string  `json:"period_end" binding:"required,datetime=2006-01-02"`
	Description string  `json:"description" binding:"required"`
	EnergyType  string  `json:"energy_type" binding:"required"`
}

type UpdateOffer struct{}
