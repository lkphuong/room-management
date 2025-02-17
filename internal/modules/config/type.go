package config

type ConfigDatabaseResponse struct {
	Username string `boil:"username" json:"username"`
	Password string `boil:"password" json:"password"`
	Host     string `boil:"host" json:"host"`
	Port     string `boil:"port" json:"port"`
	Database string `boil:"database" json:"database"`
}

type ConfigStoreResponse struct {
	ConfigKey   string `boil:"config_key" json:"config_key"`
	ConfigValue string `boil:"config_value" json:"config_value"`
}