package postgres

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/models"
	"gorm.io/gorm"
)

type migrationRepository struct {
	Config *config.DatabaseConfig
	Conn   *gorm.DB
}

func NewMigrationRepository(config *config.DatabaseConfig) (logic.MigrationRepository, error) {

	conn, err := NewConnection(config)

	if err != nil {
		return nil, err
	}

	return &migrationRepository{
		Config: config,
		Conn:   conn,
	}, nil

}

func (m migrationRepository) Migrate(currentVersion string) (string, error) {
	return "", m.Conn.AutoMigrate(&models.Migrator{})
}

func (m migrationRepository) GetMigratorVersion(name string) (string, error) {
	var migrator models.Migrator
	err := m.Conn.Where("name = ?", name).First(&migrator).Error
	if err != nil {
		return "", err
	}
	return migrator.Version, nil
}

func (m migrationRepository) UpdateMigratorVersion(name string, version string) error {
	var migrator models.Migrator
	err := m.Conn.Where("name = ?", name).First(&migrator).Error
	if err != nil {
		return err
	}
	migrator.Version = version
	return m.Conn.Save(&migrator).Error
}
