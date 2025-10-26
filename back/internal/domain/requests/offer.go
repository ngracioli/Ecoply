package requests

type CreateOffer struct {
	PricePerMwh float64 `json:"price_per_mwh" binding:"required"`
	QuantityMwh float64 `json:"quantity_mwh" binding:"required"`
	PeriodStart string  `json:"period_start" binding:"required"`
	PeriodEnd   string  `json:"period_end" binding:"required"`
	Description string  `json:"description" binding:"required"`
	EnergyType  string  `json:"energy_type" binding:"required"`
}

type ListOffers struct {
	Page        int    `form:"page" binding:"required,min=1"`
	PageSize    int    `form:"page_size" binding:"required,min=1,max=100"`
	Submarket   string `form:"submarket" binding:"omitempty"`
	EnergyType  string `form:"energy_type" binding:"omitempty"`
	PeriodStart string `form:"period_start" binding:"omitempty"`
	PeriodEnd   string `form:"period_end" binding:"omitempty"`
}

type UpdateOffer struct{}
