package migrations

import (
	"database/sql"
)

var AuditEventMigration = Migration{
	Version: 12,
	Name:    "create_audit_events_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS audit_events (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				user_id VARCHAR(255) NOT NULL,
				action VARCHAR(100) NOT NULL,
				resource_type VARCHAR(100) NOT NULL,
				resource_id UUID,
				details JSONB DEFAULT '{}',
				ip_address INET,
				user_agent TEXT,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
			
			CREATE INDEX idx_audit_events_tenant_id ON audit_events(tenant_id);
			CREATE INDEX idx_audit_events_user_id ON audit_events(user_id);
			CREATE INDEX idx_audit_events_action ON audit_events(action);
			CREATE INDEX idx_audit_events_resource_type ON audit_events(resource_type);
			CREATE INDEX idx_audit_events_created_at ON audit_events(created_at);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS audit_events;`)
		return err
	},
}
