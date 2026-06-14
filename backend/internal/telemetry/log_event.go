package telemetry

import (
	"time"

	"github.com/google/uuid"
)

type LogEvent struct {
	ID          uuid.UUID         `json:"id"`
	TenantID    string            `json:"tenantId"`
	Service     string            `json:"service"`
	Environment string            `json:"environment"`
	Level       LogLevel          `json:"level"`
	Message     string            `json:"message"`
	Timestamp   time.Time         `json:"timestamp"`
	Labels      map[string]string `json:"labels"`
	Fields      map[string]interface{} `json:"fields,omitempty"`
	TraceID     string            `json:"traceId,omitempty"`
	SpanID      string            `json:"spanId,omitempty"`
}

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
)

type LogEventValidator struct{}

func NewLogEventValidator() *LogEventValidator {
	return &LogEventValidator{}
}

func (v *LogEventValidator) Validate(event *LogEvent) error {
	if event.TenantID == "" {
		return ErrMissingTenantID
	}
	if event.Service == "" {
		return ErrMissingService
	}
	if event.Message == "" {
		return ErrMissingLogMessage
	}
	if event.Timestamp.IsZero() {
		return ErrMissingTimestamp
	}
	if event.Level == "" {
		event.Level = LogLevelInfo
	}
	if event.Labels == nil {
		event.Labels = make(map[string]string)
	}
	if event.Fields == nil {
		event.Fields = make(map[string]interface{})
	}
	return nil
}
