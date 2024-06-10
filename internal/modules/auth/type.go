package auth

type LoginParam struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type LoginResponse struct {
	ID   string `boil:"id" json:"id"`
	Code string `boil:"code" json:"code"`
	Name string `boil:"name" json:"name"`
}

type JwtTokenResponse struct {
	Token string `json:"token"`
}
