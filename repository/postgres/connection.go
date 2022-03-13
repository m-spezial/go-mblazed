package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewConnection(config *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Database,
		config.Password)

	gormLogger, err := log.NewGormLogger(
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	if err != nil {
		return nil, err
	}

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		return nil, err
	}

	if config.Debug {
		conn.Debug()
	}

	return conn, nil
}
