package requests

type CreatePurchase struct {
	QuantityMwh float64 `json:"quantity_mwh" binding:"required,gt=0"`
}
