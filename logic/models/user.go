package models

import (
	"gorm.io/gorm"
	"time"
)

type DBUser struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Username string
	Lastname string
	Email    string `gorm:"unique"`

	LastLogin   time.Time
	FailedLogin int
	Active      bool
	Password    []byte
}
