package store

import (
	"github.com/gin-gonic/gin"
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
