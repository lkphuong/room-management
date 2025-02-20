package reservation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (sr *Repository) ReservationAll(ctx context.Context, db *sql.DB, keyWork string) ([]ReservationResponse, error) {
	var result []ReservationResponse

	err := queries.Raw(fmt.Sprintf(SELECT_RESERVATION, keyWork, keyWork)).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
