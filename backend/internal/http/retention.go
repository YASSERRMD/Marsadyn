package http

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/database"
	"github.com/marsadyn/marsadyn/internal/retention"
)

type RetentionHandler struct {
	retentionService *retention.RetentionService
}

func NewRetentionHandler(retentionService *retention.RetentionService) *RetentionHandler {
	return &RetentionHandler{
		retentionService: retentionService,
	}
}

func (h *RetentionHandler) HandleCreatePolicy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var policy database.RetentionPolicy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.retentionService.CreatePolicy(r.Context(), &policy); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(policy)
}

func (h *RetentionHandler) HandleGetPolicies(w http.ResponseWriter, r *http.Request) {
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

	policies, err := h.retentionService.GetPolicies(r.Context(), tenantUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(policies)
}

func (h *RetentionHandler) HandleSimulateRetention(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tenantID := r.URL.Query().Get("tenantId")
	policyID := r.URL.Query().Get("policyId")
	
	if tenantID == "" || policyID == "" {
		http.Error(w, "tenantId and policyId are required", http.StatusBadRequest)
		return
	}

	tenantUUID, err := uuid.Parse(tenantID)
	if err != nil {
		http.Error(w, "Invalid tenantId", http.StatusBadRequest)
		return
	}

	policyUUID, err := uuid.Parse(policyID)
	if err != nil {
		http.Error(w, "Invalid policyId", http.StatusBadRequest)
		return
	}

	simulation, err := h.retentionService.SimulateRetention(r.Context(), tenantUUID, policyUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(simulation)
}
