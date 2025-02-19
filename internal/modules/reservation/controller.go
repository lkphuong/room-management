package reservation

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

	var q ReservationQueryAll

	KeyWork := c.DefaultQuery("keywork", "")

	q.KeyWork = KeyWork

	user := utils.GetInfoUser(c)
	storeIDs := user.StoreIDs

	if user.Code == hardcode.OPERATOR_ACCOUNT {
		
		result := service.GetReservationAll(ctx, db, q)

		utils.JSONResponse(*result, c)

		return
	}

	result := service.GetMyReservation(ctx, db, q, storeIDs)
		
	utils.JSONResponse(*result, c)
}