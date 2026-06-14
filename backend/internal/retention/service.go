package retention

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/marsadyn/marsadyn/internal/database"
)

type RetentionService struct {
	policyRepo database.RetentionPolicyRepository
	storage    RetentionStorage
}

type RetentionStorage interface {
	DeleteMetricsBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
	DeleteLogsBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
	DeleteTracesBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
	GetMetricsCountBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
	GetLogsCountBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
	GetTracesCountBefore(ctx context.Context, tenantID string, cutoff time.Time) (int64, error)
}

func NewRetentionService(policyRepo database.RetentionPolicyRepository, storage RetentionStorage) *RetentionService {
	return &RetentionService{
		policyRepo: policyRepo,
		storage:    storage,
	}
}

func (s *RetentionService) CreatePolicy(ctx context.Context, policy *database.RetentionPolicy) error {
	policy.ID = uuid.New()
	policy.CreatedAt = time.Now()
	policy.UpdatedAt = time.Now()
	
	return s.policyRepo.Create(policy)
}

func (s *RetentionService) GetPolicies(ctx context.Context, tenantID uuid.UUID) ([]database.RetentionPolicy, error) {
	return s.policyRepo.GetByTenantID(tenantID)
}

func (s *RetentionService) GetPolicyByID(ctx context.Context, id uuid.UUID) (*database.RetentionPolicy, error) {
	return s.policyRepo.GetByID(id)
}

func (s *RetentionService) UpdatePolicy(ctx context.Context, policy *database.RetentionPolicy) error {
	policy.UpdatedAt = time.Now()
	return s.policyRepo.Update(policy)
}

func (s *RetentionService) DeletePolicy(ctx context.Context, id uuid.UUID) error {
	return s.policyRepo.Delete(id)
}

func (s *RetentionService) ApplyRetention(ctx context.Context, tenantID uuid.UUID) error {
	policies, err := s.policyRepo.GetByTenantID(tenantID)
	if err != nil {
		return fmt.Errorf("failed to get retention policies: %w", err)
	}

	for _, policy := range policies {
		if !policy.IsEnabled {
			continue
		}

		cutoff := time.Now().AddDate(0, 0, -policy.RetentionDays)
		
		var deletedCount int64
		var err error

		switch policy.Type {
		case "metrics":
			deletedCount, err = s.storage.DeleteMetricsBefore(ctx, tenantID.String(), cutoff)
		case "logs":
			deletedCount, err = s.storage.DeleteLogsBefore(ctx, tenantID.String(), cutoff)
		case "traces":
			deletedCount, err = s.storage.DeleteTracesBefore(ctx, tenantID.String(), cutoff)
		default:
			log.Printf("Unknown retention type: %s", policy.Type)
			continue
		}

		if err != nil {
			log.Printf("Failed to apply retention for %s: %v", policy.Type, err)
			continue
		}

		log.Printf("Applied retention policy %s: deleted %d records", policy.Name, deletedCount)
	}

	return nil
}

type RetentionSimulation struct {
	PolicyID    uuid.UUID `json:"policyId"`
	PolicyName  string    `json:"policyName"`
	Type        string    `json:"type"`
	RetentionDays int     `json:"retentionDays"`
	AffectedRecords int64 `json:"affectedRecords"`
	CutoffDate  time.Time `json:"cutoffDate"`
}

func (s *RetentionService) SimulateRetention(ctx context.Context, tenantID uuid.UUID, policyID uuid.UUID) (*RetentionSimulation, error) {
	policy, err := s.policyRepo.GetByID(policyID)
	if err != nil {
		return nil, err
	}

	cutoff := time.Now().AddDate(0, 0, -policy.RetentionDays)
	
	var count int64
	switch policy.Type {
	case "metrics":
		count, err = s.storage.GetMetricsCountBefore(ctx, tenantID.String(), cutoff)
	case "logs":
		count, err = s.storage.GetLogsCountBefore(ctx, tenantID.String(), cutoff)
	case "traces":
		count, err = s.storage.GetTracesCountBefore(ctx, tenantID.String(), cutoff)
	}

	if err != nil {
		return nil, err
	}

	return &RetentionSimulation{
		PolicyID:       policy.ID,
		PolicyName:     policy.Name,
		Type:           policy.Type,
		RetentionDays:  policy.RetentionDays,
		AffectedRecords: count,
		CutoffDate:     cutoff,
	}, nil
}
