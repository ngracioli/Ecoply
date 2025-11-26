package requests

type CreatePurchase struct {
	QuantityMwh   float64 `json:"quantity_mwh" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=pix card billet"`
}

type ListPurchase struct {
	Page          int    `form:"page" binding:"required,min=1"`
	PageSize      int    `form:"page_size" binding:"required,min=1,max=100"`
	Status        string `form:"status" binding:"omitempty"`
	PaymentMethod string `form:"payment_method" binding:"omitempty"`
	OrderPrice    string `form:"order_price" binding:"omitempty"`
	OrderQuantity string `form:"order_quantity" binding:"omitempty"`
}

type ListSold struct {
	Page          int    `form:"page" binding:"required,min=1"`
	PageSize      int    `form:"page_size" binding:"required,min=1,max=100"`
	Status        string `form:"status" binding:"omitempty"`
	PaymentMethod string `form:"payment_method" binding:"omitempty"`
	OrderPrice    string `form:"order_price" binding:"omitempty"`
	OrderQuantity string `form:"order_quantity" binding:"omitempty"`
}

type ListPurchasesFromOffer struct{}
