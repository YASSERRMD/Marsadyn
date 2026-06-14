package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
)

type Migration struct {
	Version int
	Name    string
	Up      func(tx *sql.Tx) error
	Down    func(tx *sql.Tx) error
}

type Migrator struct {
	db         *sql.DB
	migrations []Migration
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{
		db:         db,
		migrations: make([]Migration, 0),
	}
}

func (m *Migrator) Register(migration Migration) {
	m.migrations = append(m.migrations, migration)
}

func (m *Migrator) Migrate() error {
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version < m.migrations[j].Version
	})

	for _, migration := range m.migrations {
		if err := m.runMigration(migration); err != nil {
			return fmt.Errorf("migration %d failed: %w", migration.Version, err)
		}
	}
	return nil
}

func (m *Migrator) runMigration(migration Migration) error {
	log.Printf("Running migration %d: %s", migration.Version, migration.Name)

	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := migration.Up(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (m *Migrator) Rollback(targetVersion int) error {
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version > m.migrations[j].Version
	})

	for _, migration := range m.migrations {
		if migration.Version <= targetVersion {
			break
		}
		log.Printf("Rolling back migration %d: %s", migration.Version, migration.Name)

		tx, err := m.db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()

		if err := migration.Down(tx); err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}
	return nil
}
