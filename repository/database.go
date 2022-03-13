package repository

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/repository/postgres"
	"errors"
	"gorm.io/gorm"
)

func NewConnection(config *config.DatabaseConfig) (*gorm.DB, error) {
	switch config.DatabaseType {
	case "postgres":
		return postgres.NewConnection(config)
	default:
		return nil, errors.New("database type not supported")
	}
}

func NewUserRepository(config *config.DatabaseConfig) (logic.UserRepository, error) {
	switch config.DatabaseType {
	case "postgres":
		return postgres.NewUserRepository(config)
	default:
		return nil, errors.New("database type not supported")
	}
}
