package migrations

import (
	"database/sql"
)

var TenantMigration = Migration{
	Version: 1,
	Name:    "create_tenants_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS tenants (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				name VARCHAR(255) NOT NULL,
				slug VARCHAR(100) UNIQUE NOT NULL,
				plan VARCHAR(50) DEFAULT 'free',
				status VARCHAR(50) DEFAULT 'active',
				config JSONB DEFAULT '{}',
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_tenants_slug ON tenants(slug);
			CREATE INDEX idx_tenants_status ON tenants(status);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS tenants;`)
		return err
	},
}
