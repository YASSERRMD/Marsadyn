package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/marsadyn/marsadyn/internal/telemetry"
)

type ClickHouseStorage struct {
	db *sql.DB
}

func NewClickHouseStorage(dsn string) (*ClickHouseStorage, error) {
	db, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping clickhouse: %w", err)
	}

	return &ClickHouseStorage{db: db}, nil
}

func (s *ClickHouseStorage) Initialize() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS metrics (
			id UUID,
			tenant_id String,
			service String,
			environment String,
			name String,
			value Float64,
			timestamp DateTime64(3),
			labels Map(String, String),
			unit String,
			type String
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(timestamp)
		ORDER BY (tenant_id, service, name, timestamp)`,
		
		`CREATE TABLE IF NOT EXISTS logs (
			id UUID,
			tenant_id String,
			service String,
			environment String,
			level String,
			message String,
			timestamp DateTime64(3),
			labels Map(String, String),
			fields String,
			trace_id String,
			span_id String
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(timestamp)
		ORDER BY (tenant_id, service, level, timestamp)`,
		
		`CREATE TABLE IF NOT EXISTS traces (
			id UUID,
			tenant_id String,
			service String,
			environment String,
			trace_id String,
			span_id String,
			parent_span_id String,
			name String,
			operation String,
			start_time DateTime64(3),
			end_time DateTime64(3),
			duration_ms Float64,
			status String,
			labels Map(String, String)
		) ENGINE = MergeTree()
		PARTITION BY toYYYYMM(start_time)
		ORDER BY (tenant_id, service, trace_id, start_time)`,
	}

	for _, query := range queries {
		if _, err := s.db.Exec(query); err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	}

	return nil
}

func (s *ClickHouseStorage) WriteMetrics(ctx context.Context, metrics []telemetry.MetricEvent) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO metrics (id, tenant_id, service, environment, name, value, timestamp, labels, unit, type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, m := range metrics {
		if _, err := stmt.ExecContext(ctx,
			m.ID, m.TenantID, m.Service, m.Environment, m.Name,
			m.Value, m.Timestamp, m.Labels, m.Unit, string(m.Type),
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *ClickHouseStorage) WriteLogs(ctx context.Context, logs []telemetry.LogEvent) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO logs (id, tenant_id, service, environment, level, message, timestamp, labels, fields, trace_id, span_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, l := range logs {
		if _, err := stmt.ExecContext(ctx,
			l.ID, l.TenantID, l.Service, l.Environment, string(l.Level),
			l.Message, l.Timestamp, l.Labels, l.Fields, l.TraceID, l.SpanID,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *ClickHouseStorage) WriteTraces(ctx context.Context, traces []telemetry.TraceSpan) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO traces (id, tenant_id, service, environment, trace_id, span_id, parent_span_id, name, operation, start_time, end_time, duration_ms, status, labels)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, t := range traces {
		if _, err := stmt.ExecContext(ctx,
			t.ID, t.TenantID, t.Service, t.Environment, t.TraceID,
			t.SpanID, t.ParentSpanID, t.Name, t.Operation,
			t.StartTime, t.EndTime, t.DurationMs, string(t.Status), t.Labels,
		); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *ClickHouseStorage) QueryMetrics(ctx context.Context, query MetricQuery) ([]MetricResult, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{query.TenantID}

	if query.Service != "" {
		where = append(where, "service = ?")
		args = append(args, query.Service)
	}
	if query.Environment != "" {
		where = append(where, "environment = ?")
		args = append(args, query.Environment)
	}
	if query.Name != "" {
		where = append(where, "name = ?")
		args = append(args, query.Name)
	}
	if !query.Start.IsZero() {
		where = append(where, "timestamp >= ?")
		args = append(args, query.Start)
	}
	if !query.End.IsZero() {
		where = append(where, "timestamp <= ?")
		args = append(args, query.End)
	}

	queryStr := fmt.Sprintf(`
		SELECT id, name, value, timestamp, labels
		FROM metrics
		WHERE %s
		ORDER BY timestamp DESC
	`, strings.Join(where, " AND "))

	if query.Limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", query.Limit)
	}

	rows, err := s.db.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []MetricResult
	for rows.Next() {
		var r MetricResult
		if err := rows.Scan(&r.ID, &r.Name, &r.Value, &r.Timestamp, &r.Labels); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func (s *ClickHouseStorage) QueryLogs(ctx context.Context, query LogQuery) ([]LogResult, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{query.TenantID}

	if query.Service != "" {
		where = append(where, "service = ?")
		args = append(args, query.Service)
	}
	if query.Environment != "" {
		where = append(where, "environment = ?")
		args = append(args, query.Environment)
	}
	if query.Level != "" {
		where = append(where, "level = ?")
		args = append(args, query.Level)
	}
	if query.Search != "" {
		where = append(where, "message LIKE ?")
		args = append(args, "%"+query.Search+"%")
	}
	if !query.Start.IsZero() {
		where = append(where, "timestamp >= ?")
		args = append(args, query.Start)
	}
	if !query.End.IsZero() {
		where = append(where, "timestamp <= ?")
		args = append(args, query.End)
	}

	queryStr := fmt.Sprintf(`
		SELECT id, level, message, timestamp, labels
		FROM logs
		WHERE %s
		ORDER BY timestamp DESC
	`, strings.Join(where, " AND "))

	if query.Limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", query.Limit)
	}
	if query.Offset > 0 {
		queryStr += fmt.Sprintf(" OFFSET %d", query.Offset)
	}

	rows, err := s.db.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []LogResult
	for rows.Next() {
		var r LogResult
		if err := rows.Scan(&r.ID, &r.Level, &r.Message, &r.Timestamp, &r.Labels); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func (s *ClickHouseStorage) QueryTraces(ctx context.Context, query TraceQuery) ([]TraceResult, error) {
	where := []string{"tenant_id = ?"}
	args := []interface{}{query.TenantID}

	if query.Service != "" {
		where = append(where, "service = ?")
		args = append(args, query.Service)
	}
	if query.Environment != "" {
		where = append(where, "environment = ?")
		args = append(args, query.Environment)
	}
	if query.TraceID != "" {
		where = append(where, "trace_id = ?")
		args = append(args, query.TraceID)
	}
	if query.MinDuration > 0 {
		where = append(where, "duration_ms >= ?")
		args = append(args, query.MinDuration)
	}
	if query.MaxDuration > 0 {
		where = append(where, "duration_ms <= ?")
		args = append(args, query.MaxDuration)
	}
	if query.Status != "" {
		where = append(where, "status = ?")
		args = append(args, query.Status)
	}
	if !query.Start.IsZero() {
		where = append(where, "start_time >= ?")
		args = append(args, query.Start)
	}
	if !query.End.IsZero() {
		where = append(where, "start_time <= ?")
		args = append(args, query.End)
	}

	queryStr := fmt.Sprintf(`
		SELECT id, trace_id, span_id, name, start_time, end_time, duration_ms, status, labels
		FROM traces
		WHERE %s
		ORDER BY start_time DESC
	`, strings.Join(where, " AND "))

	if query.Limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", query.Limit)
	}
	if query.Offset > 0 {
		queryStr += fmt.Sprintf(" OFFSET %d", query.Offset)
	}

	rows, err := s.db.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []TraceResult
	for rows.Next() {
		var r TraceResult
		if err := rows.Scan(&r.ID, &r.TraceID, &r.SpanID, &r.Name, &r.StartTime, &r.EndTime, &r.DurationMs, &r.Status, &r.Labels); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}

func (s *ClickHouseStorage) GetMetricsSummary(ctx context.Context, tenantID string, start, end time.Time) (*MetricsSummary, error) {
	summary := &MetricsSummary{}

	err := s.db.QueryRowContext(ctx, `
		SELECT COUNT(DISTINCT name) as total_series, COUNT(*) as total_samples
		FROM metrics
		WHERE tenant_id = ? AND timestamp >= ? AND timestamp <= ?
	`, tenantID, start, end).Scan(&summary.TotalSeries, &summary.TotalSamples)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, `
		SELECT service, COUNT(*) as count
		FROM metrics
		WHERE tenant_id = ? AND timestamp >= ? AND timestamp <= ?
		GROUP BY service
		ORDER BY count DESC
		LIMIT 10
	`, tenantID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sc ServiceCount
		if err := rows.Scan(&sc.Service, &sc.Count); err != nil {
			return nil, err
		}
		summary.Services = append(summary.Services, sc)
	}

	return summary, nil
}

func (s *ClickHouseStorage) GetLogsSummary(ctx context.Context, tenantID string, start, end time.Time) (*LogsSummary, error) {
	summary := &LogsSummary{
		ByLevel: make(map[string]int64),
	}

	err := s.db.QueryRowContext(ctx, `
		SELECT COUNT(*) as total_logs
		FROM logs
		WHERE tenant_id = ? AND timestamp >= ? AND timestamp <= ?
	`, tenantID, start, end).Scan(&summary.TotalLogs)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, `
		SELECT level, COUNT(*) as count
		FROM logs
		WHERE tenant_id = ? AND timestamp >= ? AND timestamp <= ?
		GROUP BY level
	`, tenantID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var level string
		var count int64
		if err := rows.Scan(&level, &count); err != nil {
			return nil, err
		}
		summary.ByLevel[level] = count
	}

	return summary, nil
}

func (s *ClickHouseStorage) GetTracesSummary(ctx context.Context, tenantID string, start, end time.Time) (*TracesSummary, error) {
	summary := &TracesSummary{}

	err := s.db.QueryRowContext(ctx, `
		SELECT COUNT(DISTINCT trace_id) as total_traces, COUNT(*) as total_spans, AVG(duration_ms) as avg_duration
		FROM traces
		WHERE tenant_id = ? AND start_time >= ? AND start_time <= ?
	`, tenantID, start, end).Scan(&summary.TotalTraces, &summary.TotalSpans, &summary.AvgDurationMs)
	if err != nil {
		return nil, err
	}

	var errorCount int64
	err = s.db.QueryRowContext(ctx, `
		SELECT COUNT(*) as error_count
		FROM traces
		WHERE tenant_id = ? AND start_time >= ? AND start_time <= ? AND status = 'error'
	`, tenantID, start, end).Scan(&errorCount)
	if err != nil {
		return nil, err
	}

	if summary.TotalSpans > 0 {
		summary.ErrorRate = float64(errorCount) / float64(summary.TotalSpans) * 100
	}

	return summary, nil
}

func (s *ClickHouseStorage) Close() error {
	return s.db.Close()
}
