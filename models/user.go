package models

import (
	"time"
)

type DBUser struct {
	DBModel

	Username string `gorm:"index,unique"`
	Email    string `gorm:"index"`

	LastLogin   time.Time
	FailedLogin int
	Active      bool
	Password    []byte
}
