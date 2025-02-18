package reservation

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/configs/http_code"
	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
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
			CuaHangId:    reservation[i].CuaHangId,
			Amount:       reservation[i].Amount,
			Mobile:       reservation[i].Mobile,
			Date:    	  reservation[i].Date,	
			CustomerName: utils.NullStringToString(reservation[i].CustomerName),
			Phong:        utils.NullStringToString(reservation[i].Phong),
			Email:        utils.NullStringToString(reservation[i].Email),
			StoreNote:    utils.NullStringToString(reservation[i].StoreNote),
			CustomerNote: utils.NullStringToString(reservation[i].CustomerNote),
			ApprovedBy:   utils.NullStringToString(reservation[i].ApprovedBy),
			ApprovedDate: utils.NullStringToString(reservation[i].ApprovedDate),
			PickupBy:     utils.NullStringToString(reservation[i].PickupBy),
			PickupDate:   utils.NullStringToString(reservation[i].PickupDate),
			Reason:       utils.NullStringToString(reservation[i].Reason),
		}

		result = append(result, r)
	}

	return utils.NewResponse(result, "Get all reservation success", http_code.OK)
}