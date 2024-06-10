package room

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/queries"
)

type Repository struct{}

func (rr *Repository) GetRooms(ctx context.Context, db *sql.DB) ([]RoomResponse, error) {
	var result []RoomResponse

	err := queries.Raw(SELECT_ROOMS).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rr *Repository) GetRoomsByStore(ctx context.Context, db *sql.DB, store string) ([]RoomResponse, error) {
	var result []RoomResponse

	err := queries.Raw(fmt.Sprintf(SELECT_ROOMS_BY_STORE, store)).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
