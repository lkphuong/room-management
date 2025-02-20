package auth

type LoginParam struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type LoginResponse struct {
	ID      string `boil:"id" json:"id"`
	Code    string `boil:"code" json:"code"`
	Name    string `boil:"name" json:"name"`
	StoreID string `boil:"store_id" json:"store_id"`
}

type JwtPayload struct {
	ID       string   `json:"id"`
	Code     string   `json:"code"`
	Name     string   `json:"name"`
	StoreIDs []string `json:"store_ids"`
}

type JwtTokenResponse struct {
	Token string `json:"token"`
}
