package room

type RoomResponse struct {
	StoreCode  string  `boil:"store_code" json:"store_code"`
	StoreName  string  `boil:"store_name" json:"store_name"`
	RoomCode   string  `boil:"room_code" json:"room_code"`
	Start      string  `boil:"start" json:"start"`
	Opened     string  `boil:"opened" json:"opened"`
	Revenue    float64 `boil:"revenue" json:"revenue"`
	RevenueTmp float64 `boil:"revenue_tmp" json:"revenue_tmp"`
}

type StatusResponse struct {
	Total    int64          `json:"total"`
	Active   int64          `json:"active"`
	InActive int64          `json:"inactive"`
	Data     []RoomResponse `json:"data"`
}

type RequestParam struct {
	Store string `form:"store" json:"store" validate:"omitempty,max=255"`
}
