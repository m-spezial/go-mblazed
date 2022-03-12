package logic

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/models"
)

// UserRepository is an interface for storing persistent data of the User
type UserRepository interface {
	GetByID(id string) (*models.DBUser, error)
	GetByEmail(email string) (*models.DBUser, error)

	Create(user *models.DBUser) error
	Update(user *models.DBUser) error
	Delete(id string) error

	GetAll() ([]*models.DBUser, error)
	GetByFilter(filter string) ([]*models.DBUser, error)
}
