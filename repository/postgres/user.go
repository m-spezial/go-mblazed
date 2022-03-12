package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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
