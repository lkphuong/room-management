package store

import (
	"context"
	"database/sql"

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
