package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/marsadyn/marsadyn/internal/telemetry"
)

type ValidationWorker struct {
	consumer        *Consumer
	metricProducer  *Producer
	logProducer     *Producer
	traceProducer   *Producer
	deadLetterProducer *Producer
}

func NewValidationWorker(
	consumer *Consumer,
	metricProducer, logProducer, traceProducer, deadLetterProducer *Producer,
) *ValidationWorker {
	return &ValidationWorker{
		consumer:           consumer,
		metricProducer:     metricProducer,
		logProducer:        logProducer,
		traceProducer:      traceProducer,
		deadLetterProducer: deadLetterProducer,
	}
}

func (w *ValidationWorker) Start(ctx context.Context) error {
	return w.consumer.Consume(ctx, w.handleMessage)
}

func (w *ValidationWorker) handleMessage(msg interface{}) error {
	var rawEvent map[string]interface{}
	if err := json.Unmarshal(msg.Value, &rawEvent); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		w.publishToDeadLetter(msg.Value, "unmarshal_error")
		return nil
	}

	eventType, ok := rawEvent["type"].(string)
	if !ok {
		log.Printf("Missing event type")
		w.publishToDeadLetter(msg.Value, "missing_type")
		return nil
	}

	switch eventType {
	case "metric":
		return w.validateMetric(rawEvent)
	case "log":
		return w.validateLog(rawEvent)
	case "trace":
		return w.validateTrace(rawEvent)
	default:
		log.Printf("Unknown event type: %s", eventType)
		w.publishToDeadLetter(msg.Value, "unknown_type")
		return nil
	}
}

func (w *ValidationWorker) validateMetric(rawEvent map[string]interface{}) error {
	var event telemetry.MetricEvent
	data, _ := json.Marshal(rawEvent)
	if err := json.Unmarshal(data, &event); err != nil {
		w.publishToDeadLetter(data, "invalid_metric_format")
		return nil
	}

	validator := telemetry.NewMetricEventValidator()
	if err := validator.Validate(&event); err != nil {
		w.publishToDeadLetter(data, "validation_error: "+err.Error())
		return nil
	}

	ctx := context.Background()
	return w.metricProducer.Publish(ctx, event.ID.String(), event)
}

func (w *ValidationWorker) validateLog(rawEvent map[string]interface{}) error {
	var event telemetry.LogEvent
	data, _ := json.Marshal(rawEvent)
	if err := json.Unmarshal(data, &event); err != nil {
		w.publishToDeadLetter(data, "invalid_log_format")
		return nil
	}

	validator := telemetry.NewLogEventValidator()
	if err := validator.Validate(&event); err != nil {
		w.publishToDeadLetter(data, "validation_error: "+err.Error())
		return nil
	}

	ctx := context.Background()
	return w.logProducer.Publish(ctx, event.ID.String(), event)
}

func (w *ValidationWorker) validateTrace(rawEvent map[string]interface{}) error {
	var span telemetry.TraceSpan
	data, _ := json.Marshal(rawEvent)
	if err := json.Unmarshal(data, &span); err != nil {
		w.publishToDeadLetter(data, "invalid_trace_format")
		return nil
	}

	validator := telemetry.NewTraceSpanValidator()
	if err := validator.Validate(&span); err != nil {
		w.publishToDeadLetter(data, "validation_error: "+err.Error())
		return nil
	}

	ctx := context.Background()
	return w.traceProducer.Publish(ctx, span.ID.String(), span)
}

func (w *ValidationWorker) publishToDeadLetter(data []byte, reason string) {
	ctx := context.Background()
	deadLetterEvent := map[string]interface{}{
		"originalData": string(data),
		"reason":       reason,
		"timestamp":    time.Now().UTC(),
	}
	if err := w.deadLetterProducer.Publish(ctx, "", deadLetterEvent); err != nil {
		log.Printf("Failed to publish to dead letter: %v", err)
	}
}
