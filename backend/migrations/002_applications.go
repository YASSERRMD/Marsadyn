package migrations

import (
	"database/sql"
)

var ApplicationMigration = Migration{
	Version: 2,
	Name:    "create_applications_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS applications (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				repository_url VARCHAR(500),
				team VARCHAR(100),
				tags JSONB DEFAULT '[]',
				status VARCHAR(50) DEFAULT 'active',
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_applications_tenant_id ON applications(tenant_id);
			CREATE INDEX idx_applications_name ON applications(name);
			CREATE INDEX idx_applications_status ON applications(status);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS applications;`)
		return err
	},
}
