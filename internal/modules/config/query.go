package config

const (
	SELECT_CONFIG_STORE = `
	SELECT TOP 1 config_key, config_value FROM config_tbl WHERE config_key = '%s'`
)
