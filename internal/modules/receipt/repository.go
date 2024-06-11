package receipt

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (r *Repository) RevenueStore(ctx context.Context, db *sql.DB) ([]RevenueResponse, error) {
	var result []RevenueResponse

	err := queries.Raw(GET_RECEIPT).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) RevenueRoom(ctx context.Context, db *sql.DB, store string) ([]RevenueRoomResponse, error) {
	var result []RevenueRoomResponse

	err := queries.Raw(fmt.Sprintf(SELECT_RECEIPT_BY_STORE, store)).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) BillDetail(ctx context.Context, db *sql.DB, store string, room string) ([]ReceiptDetailResponse, error) {
	var result []ReceiptDetailResponse

	err := queries.Raw(fmt.Sprintf(BILL_DETAIL, store, room)).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
