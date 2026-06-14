package migrations

import (
	"database/sql"
)

var ServiceMigration = Migration{
	Version: 3,
	Name:    "create_services_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS services (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				application_id UUID REFERENCES applications(id) ON DELETE SET NULL,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				version VARCHAR(100),
				environment VARCHAR(100) NOT NULL,
				status VARCHAR(50) DEFAULT 'active',
				metadata JSONB DEFAULT '{}',
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_services_tenant_id ON services(tenant_id);
			CREATE INDEX idx_services_application_id ON services(application_id);
			CREATE INDEX idx_services_name ON services(name);
			CREATE INDEX idx_services_environment ON services(environment);
			CREATE INDEX idx_services_status ON services(status);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS services;`)
		return err
	},
}
