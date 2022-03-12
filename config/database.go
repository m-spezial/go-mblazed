package config

import "code.m-spezial.de/M-Spezial/go-mblazed/config/env"

type DatabaseConfig struct {
	Debug        bool   `env:"MBLAZED_DB_DEBUG,debug=false"`
	DatabaseType string `env:"MBLAZED_DB_TYPE,required"`
	Host         string `env:"MBLAZED_DB_HOST,required"`
	Port         int    `env:"MBLAZED_DB_PORT,required"`
	Username     string `env:"MBLAZED_DB_USER,required"`
	Password     string `env:"MBLAZED_DB_PASS"`
	Database     string `env:"MBLAZED_DB_DATABASE,required"`
}

func NewDatabaseConfig(dbType string, host string, port int, database string, username string, password string) *DatabaseConfig {
	return &DatabaseConfig{
		Debug:        false,
		DatabaseType: dbType,
		Host:         host,
		Port:         port,
		Username:     username,
		Password:     password,
		Database:     database,
	}
}

// NewDatabaseConfigFromEnv create a new DatabaseConfig and loads the values from the environment variables
func NewDatabaseConfigFromEnv() (*DatabaseConfig, error) {
	config := &DatabaseConfig{}

	err := config.LoadConfigFromEnv()

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (dc *DatabaseConfig) LoadConfigFromEnv() error {
	if err := env.ReadValues(dc); err != nil {
		return err
	}
	return nil
}
