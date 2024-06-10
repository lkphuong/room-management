package room

import (
	"context"
	"database/sql"

	"github.com/lkphuong/room-management/internal/utils"
)

var (
	repository Repository
)

func init() {
	repository = Repository{}
}

type Service struct{}

func (s *Service) GetRoomByStores(ctx context.Context, db *sql.DB, store string) *utils.Response {

	var rooms []RoomResponse

	rooms, err := repository.GetRoomsByStore(ctx, db, store)

	utils.FailOnError(err, "Failed to get rooms")

	var counter = 0
	var roomsResponse []RoomResponse
	for _, room := range rooms {
		if len(room.Start) > 0 {
			counter++
		}
		var roomResponse = RoomResponse{
			StoreCode: room.StoreCode,
			RoomCode:  room.RoomCode,
			StoreName: room.StoreName,
		}

		start, err := utils.FormatDateString(room.Start)
		utils.FailOnError(err, "Failed to convert start date")

		roomResponse.Start = start

		opened := utils.CalculateTime(start)
		roomResponse.Opened = opened

		roomsResponse = append(roomsResponse, roomResponse)
	}

	response := StatusResponse{
		Total:    int64(len(rooms)),
		Active:   int64(counter),
		InActive: int64(len(rooms) - counter),
		Data:     roomsResponse,
	}

	return utils.NewResponse(response, "")
}
