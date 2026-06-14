package migrations

import (
	"database/sql"
)

var DashboardMigration = Migration{
	Version: 9,
	Name:    "create_dashboards_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS dashboards (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				layout JSONB DEFAULT '[]',
				refresh_interval INTEGER DEFAULT 30,
				is_default BOOLEAN DEFAULT false,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_dashboards_tenant_id ON dashboards(tenant_id);
			CREATE INDEX idx_dashboards_name ON dashboards(name);
			CREATE INDEX idx_dashboards_is_default ON dashboards(is_default);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS dashboards;`)
		return err
	},
}
