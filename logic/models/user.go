package models

import (
	"time"
)

type DBUser struct {
	DBModel

	Username string
	Lastname string
	Email    string `gorm:"unique"`

	LastLogin   time.Time
	FailedLogin int
	Active      bool
	Password    []byte
}
