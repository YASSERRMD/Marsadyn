package migrations

import (
	"database/sql"
)

var IngestionTokenMigration = Migration{
	Version: 11,
	Name:    "create_ingestion_tokens_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS ingestion_tokens (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				name VARCHAR(255) NOT NULL,
				token_hash VARCHAR(255) NOT NULL,
				token_prefix VARCHAR(10) NOT NULL,
				permissions JSONB DEFAULT '[]',
				expires_at TIMESTAMP WITH TIME ZONE,
				last_used_at TIMESTAMP WITH TIME ZONE,
				is_active BOOLEAN DEFAULT true,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
				updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
			
			CREATE INDEX idx_ingestion_tokens_tenant_id ON ingestion_tokens(tenant_id);
			CREATE INDEX idx_ingestion_tokens_token_hash ON ingestion_tokens(token_hash);
			CREATE INDEX idx_ingestion_tokens_is_active ON ingestion_tokens(is_active);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS ingestion_tokens;`)
		return err
	},
}
