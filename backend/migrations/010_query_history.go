package migrations

import (
	"database/sql"
)

var QueryHistoryMigration = Migration{
	Version: 10,
	Name:    "create_query_history_table",
	Up: func(tx *sql.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS query_history (
				id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
				tenant_id UUID NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
				user_id VARCHAR(255) NOT NULL,
				query_type VARCHAR(50) NOT NULL,
				query_text TEXT NOT NULL,
				filters JSONB DEFAULT '{}',
				execution_time_ms INTEGER,
				result_count INTEGER,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
			);
			
			CREATE INDEX idx_query_history_tenant_id ON query_history(tenant_id);
			CREATE INDEX idx_query_history_user_id ON query_history(user_id);
			CREATE INDEX idx_query_history_query_type ON query_history(query_type);
			CREATE INDEX idx_query_history_created_at ON query_history(created_at);
		`)
		return err
	},
	Down: func(tx *sql.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS query_history;`)
		return err
	},
}
