package alerting

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/database"
)

type AlertEvaluator struct {
	ruleRepo     database.AlertRuleRepository
	incidentRepo database.AlertIncidentRepository
	queryFunc    func(ctx context.Context, metricName string, filters map[string]string, start, end time.Time) (float64, error)
	stateCache   map[uuid.UUID]AlertState
	mu           sync.RWMutex
}

type AlertState struct {
	RuleID    uuid.UUID
	Status    string
	LastCheck time.Time
	Value     float64
}

func NewAlertEvaluator(
	ruleRepo database.AlertRuleRepository,
	incidentRepo database.AlertIncidentRepository,
	queryFunc func(ctx context.Context, metricName string, filters map[string]string, start, end time.Time) (float64, error),
) *AlertEvaluator {
	return &AlertEvaluator{
		ruleRepo:     ruleRepo,
		incidentRepo: incidentRepo,
		queryFunc:    queryFunc,
		stateCache:   make(map[uuid.UUID]AlertState),
	}
}

func (e *AlertEvaluator) EvaluateRules(ctx context.Context, tenantID uuid.UUID) error {
	rules, err := e.ruleRepo.GetEnabledByTenantID(tenantID)
	if err != nil {
		return fmt.Errorf("failed to get alert rules: %w", err)
	}

	for _, rule := range rules {
		if err := e.evaluateRule(ctx, &rule); err != nil {
			log.Printf("Failed to evaluate rule %s: %v", rule.ID, err)
			continue
		}
	}

	return nil
}

func (e *AlertEvaluator) evaluateRule(ctx context.Context, rule *database.AlertRule) error {
	e.mu.RLock()
	state, exists := e.stateCache[rule.ID]
	e.mu.RUnlock()

	if exists && time.Since(state.LastCheck) < time.Duration(rule.CooldownSeconds)*time.Second {
		return nil
	}

	var value float64
	var err error

	switch rule.Type {
	case "threshold":
		value, err = e.evaluateThresholdRule(ctx, rule)
	case "log_pattern":
		value, err = e.evaluateLogPatternRule(ctx, rule)
	case "trace_latency":
		value, err = e.evaluateTraceLatencyRule(ctx, rule)
	default:
		return fmt.Errorf("unknown rule type: %s", rule.Type)
	}

	if err != nil {
		return err
	}

	condition := rule.Condition
	threshold := condition["threshold"].(float64)
	operator := condition["operator"].(string)

	triggered := false
	switch operator {
	case "greater_than":
		triggered = value > threshold
	case "less_than":
		triggered = value < threshold
	case "equals":
		triggered = value == threshold
	case "not_equals":
		triggered = value != threshold
	}

	e.mu.Lock()
	if triggered {
		if !exists || state.Status != "firing" {
			e.stateCache[rule.ID] = AlertState{
				RuleID:    rule.ID,
				Status:    "firing",
				LastCheck: time.Now(),
				Value:     value,
			}
			go e.createIncident(rule, value)
		}
	} else {
		if exists && state.Status == "firing" {
			e.stateCache[rule.ID] = AlertState{
				RuleID:    rule.ID,
				Status:    "resolved",
				LastCheck: time.Now(),
				Value:     value,
			}
			go e.resolveIncidents(rule.ID)
		}
	}
	e.mu.Unlock()

	return nil
}

func (e *AlertEvaluator) evaluateThresholdRule(ctx context.Context, rule *database.AlertRule) (float64, error) {
	condition := rule.Condition
	metricName := condition["metric"].(string)
	filters := make(map[string]string)
	if f, ok := condition["filter"].(map[string]interface{}); ok {
		for k, v := range f {
			filters[k] = fmt.Sprintf("%v", v)
		}
	}

	end := time.Now()
	start := end.Add(-5 * time.Minute)

	return e.queryFunc(ctx, metricName, filters, start, end)
}

func (e *AlertEvaluator) evaluateLogPatternRule(ctx context.Context, rule *database.AlertRule) (float64, error) {
	condition := rule.Condition
	pattern := condition["pattern"].(string)
	filters := make(map[string]string)
	if f, ok := condition["filter"].(map[string]interface{}); ok {
		for k, v := range f {
			filters[k] = fmt.Sprintf("%v", v)
		}
	}

	end := time.Now()
	start := end.Add(-5 * time.Minute)

	return e.queryFunc(ctx, "log_count:"+pattern, filters, start, end)
}

func (e *AlertEvaluator) evaluateTraceLatencyRule(ctx context.Context, rule *database.AlertRule) (float64, error) {
	condition := rule.Condition
	service := condition["service"].(string)
	percentile := condition["percentile"].(float64)

	end := time.Now()
	start := end.Add(-5 * time.Minute)

	return e.queryFunc(ctx, "trace_latency_p"+fmt.Sprintf("%.0f", percentile), map[string]string{"service": service}, start, end)
}

func (e *AlertEvaluator) createIncident(rule *database.AlertRule, value float64) {
	incident := &database.AlertIncident{
		ID:         uuid.New(),
		TenantID:   rule.TenantID,
		RuleID:     rule.ID,
		Status:     "firing",
		Severity:   rule.Severity,
		Metadata:   map[string]interface{}{"value": value},
		StartedAt:  time.Now(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := e.incidentRepo.Create(incident); err != nil {
		log.Printf("Failed to create incident for rule %s: %v", rule.ID, err)
	}
}

func (e *AlertEvaluator) resolveIncidents(ruleID uuid.UUID) {
	incidents, err := e.incidentRepo.GetByRuleID(ruleID)
	if err != nil {
		log.Printf("Failed to get incidents for rule %s: %v", ruleID, err)
		return
	}

	now := time.Now()
	for _, incident := range incidents {
		if incident.Status == "firing" {
			incident.Status = "resolved"
			incident.ResolvedAt = &now
			incident.UpdatedAt = now
			if err := e.incidentRepo.Update(&incident); err != nil {
				log.Printf("Failed to resolve incident %s: %v", incident.ID, err)
			}
		}
	}
}
