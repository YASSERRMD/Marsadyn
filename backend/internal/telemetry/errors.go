package telemetry

import "errors"

var (
	ErrMissingTenantID    = errors.New("tenantId is required")
	ErrMissingService     = errors.New("service is required")
	ErrMissingMetricName  = errors.New("metric name is required")
	ErrMissingLogMessage  = errors.New("log message is required")
	ErrMissingTraceID     = errors.New("traceId is required")
	ErrMissingSpanID      = errors.New("spanId is required")
	ErrMissingSpanName    = errors.New("span name is required")
	ErrMissingTimestamp   = errors.New("timestamp is required")
	ErrInvalidLogLevel    = errors.New("invalid log level")
	ErrInvalidMetricType  = errors.New("invalid metric type")
	ErrInvalidSpanStatus  = errors.New("invalid span status")
)
