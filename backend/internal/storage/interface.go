package storage

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/telemetry"
)

type Storage interface {
	WriteMetrics(ctx context.Context, metrics []telemetry.MetricEvent) error
	WriteLogs(ctx context.Context, logs []telemetry.LogEvent) error
	WriteTraces(ctx context.Context, traces []telemetry.TraceSpan) error
	
	QueryMetrics(ctx context.Context, query MetricQuery) ([]MetricResult, error)
	QueryLogs(ctx context.Context, query LogQuery) ([]LogResult, error)
	QueryTraces(ctx context.Context, query TraceQuery) ([]TraceResult, error)
	
	GetMetricsSummary(ctx context.Context, tenantID string, start, end time.Time) (*MetricsSummary, error)
	GetLogsSummary(ctx context.Context, tenantID string, start, end time.Time) (*LogsSummary, error)
	GetTracesSummary(ctx context.Context, tenantID string, start, end time.Time) (*TracesSummary, error)
	
	Close() error
}

type MetricQuery struct {
	TenantID    string            `json:"tenantId"`
	Service     string            `json:"service,omitempty"`
	Environment string            `json:"environment,omitempty"`
	Name        string            `json:"name,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Start       time.Time         `json:"start"`
	End         time.Time         `json:"end"`
	Aggregation string            `json:"aggregation,omitempty"`
	Interval    string            `json:"interval,omitempty"`
	Limit       int               `json:"limit,omitempty"`
}

type LogQuery struct {
	TenantID    string            `json:"tenantId"`
	Service     string            `json:"service,omitempty"`
	Environment string            `json:"environment,omitempty"`
	Level       string            `json:"level,omitempty"`
	Search      string            `json:"search,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Start       time.Time         `json:"start"`
	End         time.Time         `json:"end"`
	Limit       int               `json:"limit,omitempty"`
	Offset      int               `json:"offset,omitempty"`
}

type TraceQuery struct {
	TenantID    string            `json:"tenantId"`
	Service     string            `json:"service,omitempty"`
	Environment string            `json:"environment,omitempty"`
	TraceID     string            `json:"traceId,omitempty"`
	MinDuration float64           `json:"minDuration,omitempty"`
	MaxDuration float64           `json:"maxDuration,omitempty"`
	Status      string            `json:"status,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Start       time.Time         `json:"start"`
	End         time.Time         `json:"end"`
	Limit       int               `json:"limit,omitempty"`
	Offset      int               `json:"offset,omitempty"`
}

type MetricResult struct {
	ID        uuid.UUID         `json:"id"`
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Timestamp time.Time         `json:"timestamp"`
	Labels    map[string]string `json:"labels"`
}

type LogResult struct {
	ID        uuid.UUID              `json:"id"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Timestamp time.Time              `json:"timestamp"`
	Labels    map[string]string      `json:"labels"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

type TraceResult struct {
	ID         uuid.UUID         `json:"id"`
	TraceID    string            `json:"traceId"`
	SpanID     string            `json:"spanId"`
	Name       string            `json:"name"`
	StartTime  time.Time         `json:"startTime"`
	EndTime    time.Time         `json:"endTime"`
	DurationMs float64           `json:"durationMs"`
	Status     string            `json:"status"`
	Labels     map[string]string `json:"labels"`
}

type MetricsSummary struct {
	TotalSeries  int64            `json:"totalSeries"`
	TotalSamples int64            `json:"totalSamples"`
	Services     []ServiceCount   `json:"services"`
	TopMetrics   []MetricCount    `json:"topMetrics"`
}

type LogsSummary struct {
	TotalLogs  int64            `json:"totalLogs"`
	ByLevel    map[string]int64 `json:"byLevel"`
	Services   []ServiceCount   `json:"services"`
	TopErrors  []LogCount       `json:"topErrors"`
}

type TracesSummary struct {
	TotalTraces   int64          `json:"totalTraces"`
	TotalSpans    int64          `json:"totalSpans"`
	AvgDurationMs float64        `json:"avgDurationMs"`
	Services      []ServiceCount `json:"services"`
	ErrorRate     float64        `json:"errorRate"`
}

type ServiceCount struct {
	Service string `json:"service"`
	Count   int64  `json:"count"`
}

type MetricCount struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type LogCount struct {
	Message string `json:"message"`
	Count   int64  `json:"count"`
}
