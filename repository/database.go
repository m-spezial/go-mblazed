package repository

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/repository/postgres"
	"errors"
)

func NewRepository(config *config.DatabaseConfig) (logic.Repository, error) {
	switch config.DatabaseType {
	case "postgres":
		return postgres.NewRepository(config)
	default:
		return nil, errors.New("database type not supported")
	}
}
