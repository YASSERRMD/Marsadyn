package query

import (
	"context"
	"time"

	"github.com/marsadyn/marsadyn/internal/storage"
)

type QueryService struct {
	storage storage.Storage
}

func NewQueryService(storage storage.Storage) *QueryService {
	return &QueryService{
		storage: storage,
	}
}

func (s *QueryService) QueryMetrics(ctx context.Context, tenantID string, query storage.MetricQuery) ([]storage.MetricResult, error) {
	query.TenantID = tenantID
	
	if query.Start.IsZero() {
		query.Start = time.Now().Add(-24 * time.Hour)
	}
	if query.End.IsZero() {
		query.End = time.Now()
	}
	if query.Limit == 0 {
		query.Limit = 1000
	}
	
	return s.storage.QueryMetrics(ctx, query)
}

func (s *QueryService) QueryLogs(ctx context.Context, tenantID string, query storage.LogQuery) ([]storage.LogResult, error) {
	query.TenantID = tenantID
	
	if query.Start.IsZero() {
		query.Start = time.Now().Add(-24 * time.Hour)
	}
	if query.End.IsZero() {
		query.End = time.Now()
	}
	if query.Limit == 0 {
		query.Limit = 1000
	}
	
	return s.storage.QueryLogs(ctx, query)
}

func (s *QueryService) QueryTraces(ctx context.Context, tenantID string, query storage.TraceQuery) ([]storage.TraceResult, error) {
	query.TenantID = tenantID
	
	if query.Start.IsZero() {
		query.Start = time.Now().Add(-24 * time.Hour)
	}
	if query.End.IsZero() {
		query.End = time.Now()
	}
	if query.Limit == 0 {
		query.Limit = 1000
	}
	
	return s.storage.QueryTraces(ctx, query)
}

func (s *QueryService) GetMetricsSummary(ctx context.Context, tenantID string, start, end time.Time) (*storage.MetricsSummary, error) {
	if start.IsZero() {
		start = time.Now().Add(-24 * time.Hour)
	}
	if end.IsZero() {
		end = time.Now()
	}
	
	return s.storage.GetMetricsSummary(ctx, tenantID, start, end)
}

func (s *QueryService) GetLogsSummary(ctx context.Context, tenantID string, start, end time.Time) (*storage.LogsSummary, error) {
	if start.IsZero() {
		start = time.Now().Add(-24 * time.Hour)
	}
	if end.IsZero() {
		end = time.Now()
	}
	
	return s.storage.GetLogsSummary(ctx, tenantID, start, end)
}

func (s *QueryService) GetTracesSummary(ctx context.Context, tenantID string, start, end time.Time) (*storage.TracesSummary, error) {
	if start.IsZero() {
		start = time.Now().Add(-24 * time.Hour)
	}
	if end.IsZero() {
		end = time.Now()
	}
	
	return s.storage.GetTracesSummary(ctx, tenantID, start, end)
}
