package config

// AppConfig struct
type AppConfig struct {
	Environment string           `mapstructure:"ENVIRONMENT"`
	Postgresql  PostgresqlConfig `mapstructure:"POSTGRESQL"`
}

// PostgresqlConfig struct
type PostgresqlConfig struct {
	ConnectionURL string `mapstructure:"CONNECTION_URL"`
}
