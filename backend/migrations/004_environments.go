package migrations

import (
	"database/sql"
)

var EnvironmentMigration = Migration{
	Version: 4,
	Name:    "create_environments_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS environments (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(100) NOT NULL,
				description TEXT,
				color VARCHAR(7) DEFAULT '#3B82F6',
				is_production BOOLEAN DEFAULT false,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				UNIQUE(tenant_id, name)
			);
			
			CREATE INDEX idx_environments_tenant_id ON environments(tenant_id);
			CREATE INDEX idx_environments_name ON environments(name);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS environments;`)
		return err
	},
}
