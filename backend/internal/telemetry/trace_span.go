package telemetry

import (
	"time"

	"github.com/google/uuid"
)

type TraceSpan struct {
	ID            uuid.UUID         `json:"id"`
	TenantID      string            `json:"tenantId"`
	Service       string            `json:"service"`
	Environment   string            `json:"environment"`
	TraceID       string            `json:"traceId"`
	SpanID        string            `json:"spanId"`
	ParentSpanID  string            `json:"parentSpanId,omitempty"`
	Name          string            `json:"name"`
	Operation     string            `json:"operation"`
	StartTime     time.Time         `json:"startTime"`
	EndTime       time.Time         `json:"endTime"`
	DurationMs    float64           `json:"durationMs"`
	Status        SpanStatus        `json:"status"`
	Labels        map[string]string `json:"labels"`
	Events        []SpanEvent       `json:"events,omitempty"`
 Links         []SpanLink        `json:"links,omitempty"`
}

type SpanStatus string

const (
	SpanStatusOK         SpanStatus = "ok"
	SpanStatusError      SpanStatus = "error"
	SpanStatusUnset      SpanStatus = "unset"
)

type SpanEvent struct {
	Name       string                 `json:"name"`
	Timestamp  time.Time              `json:"timestamp"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

type SpanLink struct {
	TraceID    string            `json:"traceId"`
	SpanID     string            `json:"spanId"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type TraceSpanValidator struct{}

func NewTraceSpanValidator() *TraceSpanValidator {
	return &TraceSpanValidator{}
}

func (v *TraceSpanValidator) Validate(span *TraceSpan) error {
	if span.TenantID == "" {
		return ErrMissingTenantID
	}
	if span.Service == "" {
		return ErrMissingService
	}
	if span.TraceID == "" {
		return ErrMissingTraceID
	}
	if span.SpanID == "" {
		return ErrMissingSpanID
	}
	if span.Name == "" {
		return ErrMissingSpanName
	}
	if span.StartTime.IsZero() {
		return ErrMissingTimestamp
	}
	if span.EndTime.IsZero() {
		return ErrMissingTimestamp
	}
	if span.DurationMs == 0 {
		span.DurationMs = float64(span.EndTime.Sub(span.StartTime).Milliseconds())
	}
	if span.Labels == nil {
		span.Labels = make(map[string]string)
	}
	return nil
}
