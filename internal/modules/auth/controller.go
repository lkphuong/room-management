package auth

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/lkphuong/room-management/configs/http_code"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	service *Service
)

func Login(c *gin.Context) {

	r := c.Request
	ctx := r.Context()

	var param LoginParam
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&param)

	if err != nil {
		utils.SadResp(err, http_code.BAD_REQUEST, c)

		return
	}

	result := service.Login(ctx, db, param)

	utils.JSONResponse(*result, c)
}
