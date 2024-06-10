package room

var (
	repository Repository
)

func init() {
	repository = Repository{}
}

type Service struct{}

// func (s *Service) GetRoomByStores(ctx context.Context, db *sql.DB, store string) *utils.Response {

// 	var rooms []RoomRespons
// 	var err error
// 	if store != "" {
// 		rooms, err := repository.GetRoomsByStore(ctx, db, store)

// 		utils.FailOnError(err, "Không có dữ liệu hiển thị.")
// 	} else {
// 		rooms, err := repository.GetRooms(ctx, db)
// 	}

// 	var active = int64(0)
// 	var roomsResponse []RoomResponse
// 	if len(room.start) > 0 {
// 		active++
// 	}

// 	return utils.NewResponse(rooms, "")
// }
