package migrations

import (
	"database/sql"
)

var AlertIncidentMigration = Migration{
	Version: 7,
	Name:    "create_alert_incidents_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS alert_incidents (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				rule_id UUID NOT NULL REFERENCES alert_rules(id) ON DELETE CASCADE,
				status VARCHAR(50) NOT NULL DEFAULT 'firing',
				severity VARCHAR(50) NOT NULL,
				message TEXT,
				metadata JSONB DEFAULT '{}',
				started_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				resolved_at TIMESTAMP WITH TIME ZONE,
				acknowledged_at TIMESTAMP WITH TIME ZONE,
				acknowledged_by VARCHAR(255),
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
			
			CREATE INDEX idx_alert_incidents_tenant_id ON alert_incidents(tenant_id);
			CREATE INDEX idx_alert_incidents_rule_id ON alert_incidents(rule_id);
			CREATE INDEX idx_alert_incidents_status ON alert_incidents(status);
			CREATE INDEX idx_alert_incidents_severity ON alert_incidents(severity);
			CREATE INDEX idx_alert_incidents_started_at ON alert_incidents(started_at);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS alert_incidents;`)
		return err
	},
}
