package reservation

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	config "github.com/lkphuong/room-management/internal/modules/config"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
	configRepository config.Repository
)

type Service struct{}

func (s *Service) GetReservationAll(ctx context.Context, db *sql.DB, query ReservationQueryAll) *utils.Response {

	var reservation []ReservationResponse
	reservation, err := repository.ReservationAll(ctx, db, query.KeyWork)
	if utils.FailOnError(err, "Failed to get reservations") != nil {
		return utils.NewResponse(nil, "Failed to get reservations", 400)
	} 

	if reservation  == nil {
		return utils.NewResponse(nil, "Reservation detail not found", http_code.NOT_FOUND)
	}

	var result []ReservationResponseData

	for i := range reservation {
		r := ReservationResponseData{
			Id:           reservation[i].Id,
			Amount:       reservation[i].Amount,
			Mobile:       reservation[i].Mobile,
			Date:    	  reservation[i].Date,
			Hour:    	  reservation[i].Hour,
			Status:    	  reservation[i].Status,
			CustomerName: utils.NullStringToString(reservation[i].CustomerName),
		}

		result = append(result, r)
	}

	return utils.NewResponse(result, "Get all reservation success", http_code.OK)
}
