package store

import (
	"github.com/gin-gonic/gin"
	"github.com/lkphuong/room-management/configs/hardcode"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	service *Service
)

func GetAll(c *gin.Context) {
	r := c.Request
	ctx := r.Context()

	result := service.GetStores(ctx, db)

	utils.JSONResponse(*result, c)
}

func GetMyStores(c *gin.Context) {
	r := c.Request
	ctx := r.Context()

	user := utils.GetInfoUser(c)
	storeIDs := user.StoreIDs

	if user.Code == hardcode.OPERATOR_ACCOUNT {
		result := service.GetStores(ctx, db)

		utils.JSONResponse(*result, c)

		return
	}

	result := service.GetMyStores(ctx, db, storeIDs)

	utils.JSONResponse(*result, c)
}
