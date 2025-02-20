package config

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

type Repository struct{}

func (cs *Repository) ConfigStoreDetail(ctx context.Context, db *sql.DB, storeID string) (*ConfigDatabaseResponse, error) {
	var configValue ConfigStoreResponse

	err := queries.Raw(fmt.Sprintf(SELECT_CONFIG_STORE,storeID)).Bind(ctx, db, &configValue)

	if err != nil {
		return nil, fmt.Errorf("error executing query for storeID '%s': %v", storeID, err)
	}

	if configValue.ConfigValue == "" {
		return nil, fmt.Errorf("no config found for storeID: %s", storeID)
	}


	// Unmarshal chuỗi JSON thành struct ConfigDatabaseResponse
	var result ConfigDatabaseResponse
	err = json.Unmarshal([]byte(configValue.ConfigValue), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal JSON: %v", err)
	}

	return &result, nil
}