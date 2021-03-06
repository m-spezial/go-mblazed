package models

import (
	nanoid "github.com/matoous/go-nanoid"
	"gorm.io/gorm"
	"time"
)

type DBModel struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate will set a nano id rather than numeric ID.
func (base *DBModel) BeforeCreate(db *gorm.DB) error {
	id, err := nanoid.ID(21)

	if err != nil {
		return err
	}

	base.ID = id
	return nil
}

type Migrator struct {
	ID      string `gorm:"primaryKey"`
	Name    string `gorm:"unique"`
	Version string
}

// BeforeCreate will set a nano id rather than numeric ID.
func (base *Migrator) BeforeCreate(db *gorm.DB) error {
	id, err := nanoid.ID(21)

	if err != nil {
		return err
	}

	base.ID = id
	return nil
}
