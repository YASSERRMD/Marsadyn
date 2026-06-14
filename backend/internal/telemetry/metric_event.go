package telemetry

import (
	"time"

	"github.com/google/uuid"
)

type MetricEvent struct {
	ID          uuid.UUID         `json:"id"`
	TenantID    string            `json:"tenantId"`
	Service     string            `json:"service"`
	Environment string            `json:"environment"`
	Name        string            `json:"name"`
	Value       float64           `json:"value"`
	Timestamp   time.Time         `json:"timestamp"`
	Labels      map[string]string `json:"labels"`
	Unit        string            `json:"unit,omitempty"`
	Type        MetricType        `json:"type"`
}

type MetricType string

const (
	MetricTypeGauge   MetricType = "gauge"
	MetricTypeCounter MetricType = "counter"
	MetricTypeHistogram MetricType = "histogram"
)

type MetricEventValidator struct{}

func NewMetricEventValidator() *MetricEventValidator {
	return &MetricEventValidator{}
}

func (v *MetricEventValidator) Validate(event *MetricEvent) error {
	if event.TenantID == "" {
		return ErrMissingTenantID
	}
	if event.Service == "" {
		return ErrMissingService
	}
	if event.Name == "" {
		return ErrMissingMetricName
	}
	if event.Timestamp.IsZero() {
		return ErrMissingTimestamp
	}
	if event.Labels == nil {
		event.Labels = make(map[string]string)
	}
	return nil
}
