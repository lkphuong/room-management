package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lkphuong/room-management/internal/utils"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (sr *Repository) GetStores(ctx context.Context, db *sql.DB) ([]StoreResponse, error) {
	var result []StoreResponse

	err := queries.Raw(SELECT_STORES).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (sr *Repository) GetStoreByIDs(ctx context.Context, storeIDs []string, db *sql.DB) ([]StoreResponse, error) {
	var result []StoreResponse

	err := queries.Raw(fmt.Sprintf(SELECT_STORES_BY_IDS, utils.ConvertSliceToString(storeIDs))).Bind(ctx, db, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
