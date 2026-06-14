package telemetry

import (
	"testing"
	"time"
)

func TestMetricEventValidator(t *testing.T) {
	validator := NewMetricEventValidator()

	tests := []struct {
		name    string
		event   MetricEvent
		wantErr error
	}{
		{
			name: "valid metric event",
			event: MetricEvent{
				TenantID:    "tenant-001",
				Service:     "payment-api",
				Environment: "prod",
				Name:        "http_request_duration_ms",
				Value:       123.4,
				Timestamp:   time.Now(),
				Labels:      map[string]string{"method": "GET"},
			},
			wantErr: nil,
		},
		{
			name: "missing tenantId",
			event: MetricEvent{
				Service:     "payment-api",
				Name:        "http_request_duration_ms",
				Value:       123.4,
				Timestamp:   time.Now(),
			},
			wantErr: ErrMissingTenantID,
		},
		{
			name: "missing service",
			event: MetricEvent{
				TenantID:  "tenant-001",
				Name:      "http_request_duration_ms",
				Value:     123.4,
				Timestamp: time.Now(),
			},
			wantErr: ErrMissingService,
		},
		{
			name: "missing metric name",
			event: MetricEvent{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				Value:     123.4,
				Timestamp: time.Now(),
			},
			wantErr: ErrMissingMetricName,
		},
		{
			name: "missing timestamp",
			event: MetricEvent{
				TenantID: "tenant-001",
				Service:  "payment-api",
				Name:     "http_request_duration_ms",
				Value:    123.4,
			},
			wantErr: ErrMissingTimestamp,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(&tt.event)
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogEventValidator(t *testing.T) {
	validator := NewLogEventValidator()

	tests := []struct {
		name    string
		event   LogEvent
		wantErr error
	}{
		{
			name: "valid log event",
			event: LogEvent{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				Level:     LogLevelInfo,
				Message:   "Payment processed",
				Timestamp: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "missing tenantId",
			event: LogEvent{
				Service:   "payment-api",
				Level:     LogLevelInfo,
				Message:   "Payment processed",
				Timestamp: time.Now(),
			},
			wantErr: ErrMissingTenantID,
		},
		{
			name: "missing service",
			event: LogEvent{
				TenantID:  "tenant-001",
				Level:     LogLevelInfo,
				Message:   "Payment processed",
				Timestamp: time.Now(),
			},
			wantErr: ErrMissingService,
		},
		{
			name: "missing message",
			event: LogEvent{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				Level:     LogLevelInfo,
				Timestamp: time.Now(),
			},
			wantErr: ErrMissingLogMessage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(&tt.event)
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTraceSpanValidator(t *testing.T) {
	validator := NewTraceSpanValidator()

	tests := []struct {
		name    string
		span    TraceSpan
		wantErr error
	}{
		{
			name: "valid trace span",
			span: TraceSpan{
				TenantID:    "tenant-001",
				Service:     "payment-api",
				TraceID:     "abc123",
				SpanID:      "span123",
				Name:        "process-payment",
				StartTime:   time.Now(),
				EndTime:     time.Now().Add(100 * time.Millisecond),
				DurationMs:  100,
				Status:      SpanStatusOK,
			},
			wantErr: nil,
		},
		{
			name: "missing tenantId",
			span: TraceSpan{
				Service:   "payment-api",
				TraceID:   "abc123",
				SpanID:    "span123",
				Name:      "process-payment",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(100 * time.Millisecond),
			},
			wantErr: ErrMissingTenantID,
		},
		{
			name: "missing traceId",
			span: TraceSpan{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				SpanID:    "span123",
				Name:      "process-payment",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(100 * time.Millisecond),
			},
			wantErr: ErrMissingTraceID,
		},
		{
			name: "missing spanId",
			span: TraceSpan{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				TraceID:   "abc123",
				Name:      "process-payment",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(100 * time.Millisecond),
			},
			wantErr: ErrMissingSpanID,
		},
		{
			name: "missing span name",
			span: TraceSpan{
				TenantID:  "tenant-001",
				Service:   "payment-api",
				TraceID:   "abc123",
				SpanID:    "span123",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(100 * time.Millisecond),
			},
			wantErr: ErrMissingSpanName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(&tt.span)
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
