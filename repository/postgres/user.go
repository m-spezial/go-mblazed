package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/log"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type userRepository struct {
	Config *config.DatabaseConfig
	Conn   *gorm.DB
}

func NewUserRepository(config *config.DatabaseConfig) (logic.UserRepository, error) {

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

	return &userRepository{
		Config: config,
		Conn:   conn,
	}, nil

}

func (u userRepository) GetByID(id string) (*models.DBUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByEmail(email string) (*models.DBUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Create(user *models.DBUser) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Update(user *models.DBUser) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetAll() ([]*models.DBUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByFilter(filter string) ([]*models.DBUser, error) {
	//TODO implement me
	panic("implement me")
}
