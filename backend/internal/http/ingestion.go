package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/kafka"
	"github.com/marsadyn/marsadyn/internal/telemetry"
)

type IngestionHandler struct {
	metricProducer *kafka.Producer
	logProducer    *kafka.Producer
	traceProducer  *kafka.Producer
}

func NewIngestionHandler(metricProducer, logProducer, traceProducer *kafka.Producer) *IngestionHandler {
	return &IngestionHandler{
		metricProducer: metricProducer,
		logProducer:    logProducer,
		traceProducer:  traceProducer,
	}
}

func (h *IngestionHandler) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event telemetry.MetricEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validator := telemetry.NewMetricEventValidator()
	if err := validator.Validate(&event); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	event.ID = uuid.New()

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.metricProducer.Publish(ctx, event.ID.String(), event); err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "accepted",
		"type":    "metrics",
		"eventId": event.ID.String(),
	})
}

func (h *IngestionHandler) HandleLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var event telemetry.LogEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validator := telemetry.NewLogEventValidator()
	if err := validator.Validate(&event); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	event.ID = uuid.New()

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.logProducer.Publish(ctx, event.ID.String(), event); err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "accepted",
		"type":    "logs",
		"eventId": event.ID.String(),
	})
}

func (h *IngestionHandler) HandleTraces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var span telemetry.TraceSpan
	if err := json.Unmarshal(body, &span); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validator := telemetry.NewTraceSpanValidator()
	if err := validator.Validate(&span); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	span.ID = uuid.New()

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.traceProducer.Publish(ctx, span.ID.String(), span); err != nil {
		http.Error(w, "Failed to publish event", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "accepted",
		"type":    "traces",
		"eventId": span.ID.String(),
	})
}
