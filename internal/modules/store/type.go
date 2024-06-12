package store

type StoreResponse struct {
	StoreID   string `boil:"store_id" json:"store_id"`
	StoreName string `boil:"store_name" json:"store_name"`
	RoomCount int64  `boil:"room_count" json:"room_count"`
}

type AllStoreResponse struct {
	StoreID    string  `boil:"store_id" json:"store_id"`
	StoreName  string  `boil:"store_name" json:"store_name"`
	RoomCount  int64   `boil:"room_count" json:"room_count"`
	Active     int64   `boil:"active" json:"active"`
	InActive   int64   `boil:"inactive" json:"inactive"`
	Revenue    float64 `boil:"revenue" json:"revenue"`
	RevenueTmp float64 `boil:"revenue_tmp" json:"revenue_tmp"`
}
