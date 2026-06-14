package migrations

import (
	"database/sql"
)

var RetentionPolicyMigration = Migration{
	Version: 8,
	Name:    "create_retention_policies_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS retention_policies (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				type VARCHAR(50) NOT NULL,
				retention_days INTEGER NOT NULL,
				retention_months INTEGER,
				action VARCHAR(50) DEFAULT 'delete',
				is_enabled BOOLEAN DEFAULT true,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_retention_policies_tenant_id ON retention_policies(tenant_id);
			CREATE INDEX idx_retention_policies_type ON retention_policies(type);
			CREATE INDEX idx_retention_policies_is_enabled ON retention_policies(is_enabled);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS retention_policies;`)
		return err
	},
}
