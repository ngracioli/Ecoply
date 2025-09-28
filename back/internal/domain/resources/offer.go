package resources

import (
	"time"
)

type Offer struct {
	Uuid                 string    `json:"uuid"`
	PricePerMwh          float64   `json:"price_per_mwh"`
	InitialQuantityMwh   float64   `json:"initial_quantity_mwh"`
	RemainingQuantityMwh float64   `json:"remaining_quantity_mwh"`
	Description          string    `json:"description"`
	PeriodStart          time.Time `json:"period_start"`
	PeriodEnd            time.Time `json:"period_end"`
	Status               uint8     `json:"status"`
	EnergyType           string    `json:"energy_type"`
	Submarket            string    `json:"submarket"`
	SellerUuid           string    `json:"seller_agent_uuid"`
	CreatedAt            time.Time `json:"created_at"`
}
