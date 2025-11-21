package resources

import "time"

type Contract struct {
	Supplier ContractSupplier `json:"supplier"`
	Offer    ContractOffer    `json:"offer"`
	Buyer    ContractBuyer    `json:"buyer"`
}

type ContractSupplier struct {
	Uuid          string `json:"uuid"`
	Cnpj          string `json:"cnpj"`
	CceeCode      string `json:"ccee_code"`
	SubmarketName string `json:"submarket_name"`
	CompanyName   string `json:"company_name"`
}

type ContractOffer struct {
	Uuid                  string    `json:"uuid"`
	PricePerMwh           float64   `json:"price_per_mwh"`
	InitialQuantityMwh    float64   `json:"initial_quantity_mwh"`
	ContractedQuantityMwh float64   `json:"contracted_quantity_mwh"`
	Description           string    `json:"description"`
	PeriodStart           string    `json:"period_start"`
	PeriodEnd             string    `json:"period_end"`
	EnergyType            string    `json:"energy_type"`
	Submarket             string    `json:"submarket"`
	CreatedAt             time.Time `json:"created_at"`
}

type ContractBuyer struct {
	Uuid          string `json:"uuid"`
	Cnpj          string `json:"cnpj"`
	CceeCode      string `json:"ccee_code"`
	SubmarketName string `json:"submarket_name"`
	CompanyName   string `json:"company_name"`
}
