package migrations

import (
	"database/sql"
)

var AlertRuleMigration = Migration{
	Version: 6,
	Name:    "create_alert_rules_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS alert_rules (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				type VARCHAR(50) NOT NULL,
				severity VARCHAR(50) NOT NULL DEFAULT 'warning',
				condition JSONB NOT NULL,
				notification_channels JSONB DEFAULT '[]',
				is_enabled BOOLEAN DEFAULT true,
				cooldown_seconds INTEGER DEFAULT 300,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_alert_rules_tenant_id ON alert_rules(tenant_id);
			CREATE INDEX idx_alert_rules_type ON alert_rules(type);
			CREATE INDEX idx_alert_rules_severity ON alert_rules(severity);
			CREATE INDEX idx_alert_rules_is_enabled ON alert_rules(is_enabled);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS alert_rules;`)
		return err
	},
}
