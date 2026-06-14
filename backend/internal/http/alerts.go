package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/alerting"
	"github.com/marsadyn/marsadyn/internal/database"
)

type AlertHandler struct {
	evaluator  *alerting.AlertEvaluator
	ruleRepo   database.AlertRuleRepository
	incidentRepo database.AlertIncidentRepository
}

func NewAlertHandler(
	evaluator *alerting.AlertEvaluator,
	ruleRepo database.AlertRuleRepository,
	incidentRepo database.AlertIncidentRepository,
) *AlertHandler {
	return &AlertHandler{
		evaluator:    evaluator,
		ruleRepo:     ruleRepo,
		incidentRepo: incidentRepo,
	}
}

func (h *AlertHandler) HandleCreateRule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var rule database.AlertRule
	if err := json.NewDecoder(r.Body).Decode(&rule); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	rule.ID = uuid.New()
	rule.CreatedAt = time.Now()
	rule.UpdatedAt = time.Now()

	if err := h.ruleRepo.Create(&rule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rule)
}

func (h *AlertHandler) HandleGetRules(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tenantID := r.URL.Query().Get("tenantId")
	if tenantID == "" {
		http.Error(w, "tenantId is required", http.StatusBadRequest)
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		http.Error(w, "Invalid tenantId", http.StatusBadRequest)
		return
	}

	rules, _, err := h.ruleRepo.GetByTenantID(tenantUUID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rules)
}

func (h *AlertHandler) HandleUpdateRule(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ruleID := r.URL.Query().Get("id")
	if ruleID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	ruleUUID, err := uuid.Parse(ruleID)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updates database.AlertRule
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	existing, err := h.ruleRepo.GetByID(ruleUUID)
	if err != nil {
		http.Error(w, "Rule not found", http.StatusNotFound)
		return
	}

	if updates.Name != "" {
		existing.Name = updates.Name
	}
	if updates.Description != nil {
		existing.Description = updates.Description
	}
	if updates.Condition != nil {
		existing.Condition = updates.Condition
	}
	if updates.IsEnabled != existing.IsEnabled {
		existing.IsEnabled = updates.IsEnabled
	}
	existing.UpdatedAt = time.Now()

	if err := h.ruleRepo.Update(existing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existing)
}

func (h *AlertHandler) HandleGetIncidents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tenantID := r.URL.Query().Get("tenantId")
	if tenantID == "" {
		http.Error(w, "tenantId is required", http.StatusBadRequest)
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		http.Error(w, "Invalid tenantId", http.StatusBadRequest)
		return
	}

	incidents, _, err := h.incidentRepo.GetByTenantID(tenantUUID, 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incidents)
}

func (h *AlertHandler) HandleResolveIncident(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	incidentID := r.URL.Query().Get("id")
	if incidentID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	incidentUUID, err := uuid.Parse(incidentID)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	incident, err := h.incidentRepo.GetByID(incidentUUID)
	if err != nil {
		http.Error(w, "Incident not found", http.StatusNotFound)
		return
	}

	now := time.Now()
	incident.Status = "resolved"
	incident.ResolvedAt = &now
	incident.UpdatedAt = now

	if err := h.incidentRepo.Update(incident); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incident)
}
