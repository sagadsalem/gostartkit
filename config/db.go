package config

import "os"

type DatabaseConfig struct {
	ConnectionString string
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		ConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}
}
