package repository

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/repository/postgres"
	"errors"
)

func NewUserRepository(config *config.DatabaseConfig) (logic.UserRepository, error) {
	switch config.DatabaseType {
	case "postgres":
		return postgres.NewUserRepository(config)
	default:
		return nil, errors.New("database type not supported")
	}
}
