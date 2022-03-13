package logic

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/models"
)

type Repository interface {
	// Migrate is called by the server to initialize the database. It returns the new Version of the Database
	Migrate(currentVersion string) (string, error)
}

type MigrationRepository interface {
	Repository

	GetMigratorVersion(name string) (string, error)
	UpdateMigratorVersion(name string, version string) error
}

// UserRepository is an interface for storing persistent data of the User
type UserRepository interface {
	Repository

	GetByID(id string) (*models.DBUser, error)
	GetByEmail(email string) (*models.DBUser, error)
	GetByUsername(username string) (*models.DBUser, error)

	ExistsUsername(username string) (bool, error)

	Create(user *models.DBUser) error
	Update(user *models.DBUser) error

	SoftDelete(id string) error
	HardDelete(id string) error

	GetAll() ([]*models.DBUser, error)
	GetByFilter(filter string, vals ...interface{}) ([]*models.DBUser, error)
	GetByPagination(page *models.Pagination) ([]*models.DBUser, error)
}
