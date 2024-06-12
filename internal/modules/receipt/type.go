package receipt

type RevenueResponse struct {
	Revenue    float64 `boil:"revenue" json:"revenue"`
	StoreCode  string  `boil:"store_code" json:"store_code"`
	RevenueTmp float64 `boil:"revenue_tmp" json:"revenue_tmp"`
}

type RevenueRoomResponse struct {
	RoomCode   string  `boil:"room_code" json:"room_code"`
	Revenue    float64 `boil:"revenue" json:"revenue"`
	RevenueTmp float64 `boil:"revenue_tmp" json:"revenue_tmp"`
}

type ReceiptDetailResponse struct {
	ID       string  `boil:"id" json:"id"`
	Name     string  `boil:"name" json:"name"`
	Quantity int64   `boil:"quantity" json:"quantity"`
	Price    float64 `boil:"price" json:"price"`
	Total    float64 `boil:"total" json:"total"`
	Category string  `boil:"category_id" json:"category_id"`
}

type ReceiptDetailParam struct {
	Store string `form:"store" json:"store" validate:"omitempty,max=255"`
	Room  string `form:"room" json:"room" validate:"omitempty,max=255"`
}
