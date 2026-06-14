package migrations

import (
	"database/sql"
)

var TelemetryMetadataMigration = Migration{
	Version: 5,
	Name:    "create_telemetry_metadata_tables",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS metric_series (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				service_id UUID REFERENCES services(id) ON DELETE SET NULL,
				name VARCHAR(255) NOT NULL,
				type VARCHAR(50) NOT NULL,
				unit VARCHAR(50),
				description TEXT,
				labels JSONB DEFAULT '{}',
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_metric_series_tenant_id ON metric_series(tenant_id);
			CREATE INDEX idx_metric_series_service_id ON metric_series(service_id);
			CREATE INDEX idx_metric_series_name ON metric_series(name);
			CREATE INDEX idx_metric_series_type ON metric_series(type);
			
			CREATE TABLE IF NOT EXISTS log_streams (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				service_id UUID REFERENCES services(id) ON DELETE SET NULL,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				retention_days INTEGER DEFAULT 30,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_log_streams_tenant_id ON log_streams(tenant_id);
			CREATE INDEX idx_log_streams_service_id ON log_streams(service_id);
			CREATE INDEX idx_log_streams_name ON log_streams(name);
			
			CREATE TABLE IF NOT EXISTS trace_services (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				service_id UUID REFERENCES services(id) ON DELETE SET NULL,
				name VARCHAR(255) NOT NULL,
				description TEXT,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				deleted_at TIMESTAMP WITH TIME ZONE
			);
			
			CREATE INDEX idx_trace_services_tenant_id ON trace_services(tenant_id);
			CREATE INDEX idx_trace_services_service_id ON trace_services(service_id);
			CREATE INDEX idx_trace_services_name ON trace_services(name);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			DROP TABLE IF EXISTS trace_services;
			DROP TABLE IF EXISTS log_streams;
			DROP TABLE IF EXISTS metric_series;
		`)
		return err
	},
}
