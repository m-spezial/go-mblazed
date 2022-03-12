package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	Config *config.DatabaseConfig
	Conn   *gorm.DB
}

func NewRepository(config *config.DatabaseConfig) (logic.Repository, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Database,
		config.Password)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if config.Debug {
		conn.Debug()
	}

	return &repository{
		Config: config,
		Conn:   conn,
	}, nil

}
