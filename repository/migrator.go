package repository

import (
	"code.m-spezial.de/M-Spezial/go-mblazed/config"
	"code.m-spezial.de/M-Spezial/go-mblazed/logic"
	"code.m-spezial.de/M-Spezial/go-mblazed/repository/postgres"
	"errors"
	"reflect"
)

func newMigrationRepository(config *config.DatabaseConfig) (logic.MigrationRepository, error) {
	switch config.DatabaseType {
	case "postgres":
		return postgres.NewMigrationRepository(config)
	default:
		return nil, errors.New("database type not supported")
	}
}

type Migrator struct {
	MigrationRepository logic.MigrationRepository
	Repositories        []logic.Repository
}

func NewMigrator(config *config.DatabaseConfig) (*Migrator, error) {
	migrationRepo, err := newMigrationRepository(config)

	if err != nil {
		return nil, err
	}

	_, err = migrationRepo.Migrate("")

	if err != nil {
		return nil, err
	}

	return &Migrator{
		MigrationRepository: migrationRepo,
		Repositories:        make([]logic.Repository, 0),
	}, nil
}

func (m *Migrator) AddRepository(migrator logic.Repository) {
	m.Repositories = append(m.Repositories, migrator)
}

func (m *Migrator) Migrate() error {
	for _, migrator := range m.Repositories {
		typeData := reflect.TypeOf(migrator)
		name := typeData.Elem().PkgPath() + "." + typeData.Elem().Name()

		version, err := m.MigrationRepository.GetMigratorVersion(name)

		if err != nil {
			return err
		}

		newVersion, err := migrator.Migrate(version)

		if err != nil {
			return err
		}

		if version == newVersion {
			continue
		}

		err = m.MigrationRepository.UpdateMigratorVersion(name, newVersion)

		if err != nil {
			return err
		}
	}

	return nil
}
