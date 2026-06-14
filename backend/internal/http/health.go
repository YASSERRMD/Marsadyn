package http

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
	Uptime    string `json:"uptime"`
}

type HealthHandler struct {
	startTime time.Time
	service   string
}

func NewHealthHandler(service string) *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		service:   service,
	}
}

func (h *HealthHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(h.startTime)

	response := HealthResponse{
		Status:    "healthy",
		Service:   h.service,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   "0.1.0",
		GoVersion: runtime.Version(),
		Uptime:    uptime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *HealthHandler) HandleReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ready",
	})
}
