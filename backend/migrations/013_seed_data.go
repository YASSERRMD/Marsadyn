package migrations

import (
	"database/sql"
)

var SeedDataMigration = Migration{
	Version: 13,
	Name:    "seed_initial_data",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			INSERT INTO tenants (id, name, slug, plan, status) VALUES
				('550e8400-e29b-41d4-a716-446655440001', 'Default Organization', 'default', 'enterprise', 'active')
			ON CONFLICT (slug) DO NOTHING;
			
			INSERT INTO environments (tenant_id, name, description, color, is_production) VALUES
				('550e8400-e29b-41d4-a716-446655440001', 'production', 'Production environment', '#22C55E', true),
				('550e8400-e29b-41d4-a716-446655440001', 'staging', 'Staging environment', '#EAB308', false),
				('550e8400-e29b-41d4-a716-446655440001', 'development', 'Development environment', '#3B82F6', false)
			ON CONFLICT DO NOTHING;
			
			INSERT INTO alert_rules (tenant_id, name, description, type, severity, condition, notification_channels, is_enabled) VALUES
				('550e8400-e29b-41d4-a716-446655440001', 'High Error Rate', 'Alert when error rate exceeds 5%', 'threshold', 'critical', '{"metric":"http_requests_total","filter":{"status":"5xx"},"threshold":0.05,"operator":"greater_than"}', '["email","slack"]', true),
				('550e8400-e29b-41d4-a716-446655440001', 'High Latency', 'Alert when p99 latency exceeds 1s', 'threshold', 'warning', '{"metric":"http_request_duration_ms","filter":{},"percentile":99,"threshold":1000,"operator":"greater_than"}', '["email"]', true)
			ON CONFLICT DO NOTHING;
			
			INSERT INTO retention_policies (tenant_id, name, description, type, retention_days, action, is_enabled) VALUES
				('550e8400-e29b-41d4-a716-446655440001', 'Default Metrics Retention', 'Keep metrics for 90 days', 'metrics', 90, 'delete', true),
				('550e8400-e29b-41d4-a716-446655440001', 'Default Logs Retention', 'Keep logs for 30 days', 'logs', 30, 'delete', true),
				('550e8400-e29b-41d4-a716-446655440001', 'Default Traces Retention', 'Keep traces for 14 days', 'traces', 14, 'delete', true)
			ON CONFLICT DO NOTHING;
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			DELETE FROM retention_policies WHERE tenant_id = '550e8400-e29b-41d4-a716-446655440001';
			DELETE FROM alert_rules WHERE tenant_id = '550e8400-e29b-41d4-a716-446655440001';
			DELETE FROM environments WHERE tenant_id = '550e8400-e29b-41d4-a716-446655440001';
			DELETE FROM tenants WHERE id = '550e8400-e29b-41d4-a716-446655440001';
		`)
		return err
	},
}
