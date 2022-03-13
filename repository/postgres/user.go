package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/models"
	"errors"
	"gorm.io/gorm"
)

type userRepository struct {
	Config *config.DatabaseConfig
	Conn   *gorm.DB
}

func NewUserRepository(config *config.DatabaseConfig) (logic.UserRepository, error) {

	conn, err := NewConnection(config)

	if err != nil {
		return nil, err
	}

	return &userRepository{
		Config: config,
		Conn:   conn,
	}, nil

}

func (u userRepository) Migrate(currentVersion string) (string, error) {
	switch currentVersion {
	case "":
		return "1.0.0", u.Conn.AutoMigrate(&models.DBUser{})
	case "1.0.0":
		return "1.0.0", nil
	}

	return "", errors.New("unsupported version")
}

// GetByID returns the user by his ID from the database
func (u userRepository) GetByID(id string) (*models.DBUser, error) {
	var user models.DBUser

	if err := u.Conn.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetByEmail returns the user by his email from the database
func (u userRepository) GetByEmail(email string) (*models.DBUser, error) {
	var user models.DBUser

	if err := u.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) Create(user *models.DBUser) error {
	if err := u.Conn.Create(user).Error; err != nil {
		return err
	}

	return nil
}

// Update saves the complete struct with all information. It's important to make sure no field is empty.
func (u userRepository) Update(user *models.DBUser) error {
	err := u.Conn.Where("id = ?", user.ID).Save(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) SoftDelete(id string) error {
	return u.Conn.Where("id = ?", id).Delete(&models.DBUser{}).Error
}

func (u userRepository) HardDelete(id string) error {
	return u.Conn.Where("id = ?", id).Unscoped().Delete(&models.DBUser{}).Error
}

// GetAll returns all users from the database. Its not recommended to use this function because of performance problems
// if you have a lot of users.
func (u userRepository) GetAll() ([]*models.DBUser, error) {
	var users []*models.DBUser

	if err := u.Conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) GetByFilter(filter string, vals ...interface{}) ([]*models.DBUser, error) {
	var users []*models.DBUser

	if err := u.Conn.Where(filter, vals).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepository) GetByUsername(username string) (*models.DBUser, error) {
	var user models.DBUser

	if err := u.Conn.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// ExistsUsername checks if a username exists in the database and returns true if it does
func (u userRepository) ExistsUsername(username string) (bool, error) {
	var exists bool
	err := u.Conn.Model(models.DBUser{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u userRepository) GetByPagination(page *models.Pagination) ([]*models.DBUser, error) {

	if page.Page == 0 {
		page.Page = 1
	}

	if page.Limit == 0 {
		page.Limit = 10
	}

	offset := (page.Page - 1) * page.Limit

	var users []*models.DBUser
	err := u.Conn.Offset(offset).Limit(page.Limit).Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
