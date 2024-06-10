package room

type RoomResponse struct {
	StoreCode string `boil:"store_code" json:"store_code"`
	StoreName string `boil:"store_name" json:"store_name"`
	RoomCode  string `boil:"room_code" json:"room_code"`
	Start     string `boil:"start" json:"start"`
	Opened    string `boil:"opened" json:"opened"`
}

type StatusResponse struct {
	Total    int64          `json:"total"`
	Active   int64          `json:"active"`
	InActive int64          `json:"inactive"`
	Data     []RoomResponse `json:"data"`
}
