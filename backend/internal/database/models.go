package database

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Slug      string         `json:"slug" db:"slug"`
	Plan      string         `json:"plan" db:"plan"`
	Status    string         `json:"status" db:"status"`
	Config    map[string]interface{} `json:"config" db:"config"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Application struct {
	ID            uuid.UUID      `json:"id" db:"id"`
	TenantID      uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	Name          string         `json:"name" db:"name"`
	Description   *string        `json:"description,omitempty" db:"description"`
	RepositoryURL *string        `json:"repository_url,omitempty" db:"repository_url"`
	Team          *string        `json:"team,omitempty" db:"team"`
	Tags          []string       `json:"tags" db:"tags"`
	Status        string         `json:"status" db:"status"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Service struct {
	ID            uuid.UUID      `json:"id" db:"id"`
	TenantID      uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	ApplicationID *uuid.UUID     `json:"application_id,omitempty" db:"application_id"`
	Name          string         `json:"name" db:"name"`
	Description   *string        `json:"description,omitempty" db:"description"`
	Version       *string        `json:"version,omitempty" db:"version"`
	Environment   string         `json:"environment" db:"environment"`
	Status        string         `json:"status" db:"status"`
	Metadata      map[string]interface{} `json:"metadata" db:"metadata"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Environment struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	TenantID      uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	Name          string     `json:"name" db:"name"`
	Description   *string    `json:"description,omitempty" db:"description"`
	Color         string     `json:"color" db:"color"`
	IsProduction  bool       `json:"is_production" db:"is_production"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}

type MetricSeries struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	TenantID    uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	ServiceID   *uuid.UUID     `json:"service_id,omitempty" db:"service_id"`
	Name        string         `json:"name" db:"name"`
	Type        string         `json:"type" db:"type"`
	Unit        *string        `json:"unit,omitempty" db:"unit"`
	Description *string        `json:"description,omitempty" db:"description"`
	Labels      map[string]interface{} `json:"labels" db:"labels"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type LogStream struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	TenantID       uuid.UUID  `json:"tenant_id" db:"tenant_id"`
	ServiceID      *uuid.UUID `json:"service_id,omitempty" db:"service_id"`
	Name           string     `json:"name" db:"name"`
	Description    *string    `json:"description,omitempty" db:"description"`
	RetentionDays  int        `json:"retention_days" db:"retention_days"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type TraceService struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	TenantID    uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	ServiceID   *uuid.UUID     `json:"service_id,omitempty" db:"service_id"`
	Name        string         `json:"name" db:"name"`
	Description *string        `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type AlertRule struct {
	ID                    uuid.UUID      `json:"id" db:"id"`
	TenantID              uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	Name                  string         `json:"name" db:"name"`
	Description           *string        `json:"description,omitempty" db:"description"`
	Type                  string         `json:"type" db:"type"`
	Severity              string         `json:"severity" db:"severity"`
	Condition             map[string]interface{} `json:"condition" db:"condition"`
	NotificationChannels  []string       `json:"notification_channels" db:"notification_channels"`
	IsEnabled             bool           `json:"is_enabled" db:"is_enabled"`
	CooldownSeconds       int            `json:"cooldown_seconds" db:"cooldown_seconds"`
	CreatedAt             time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt             *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type AlertIncident struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	TenantID        uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	RuleID          uuid.UUID      `json:"rule_id" db:"rule_id"`
	Status          string         `json:"status" db:"status"`
	Severity        string         `json:"severity" db:"severity"`
	Message         *string        `json:"message,omitempty" db:"message"`
	Metadata        map[string]interface{} `json:"metadata" db:"metadata"`
	StartedAt       time.Time      `json:"started_at" db:"started_at"`
	ResolvedAt      *time.Time     `json:"resolved_at,omitempty" db:"resolved_at"`
	AcknowledgedAt  *time.Time     `json:"acknowledged_at,omitempty" db:"acknowledged_at"`
	AcknowledgedBy  *string        `json:"acknowledged_by,omitempty" db:"acknowledged_by"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
}

type RetentionPolicy struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	TenantID        uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	Name            string         `json:"name" db:"name"`
	Description     *string        `json:"description,omitempty" db:"description"`
	Type            string         `json:"type" db:"type"`
	RetentionDays   int            `json:"retention_days" db:"retention_days"`
	RetentionMonths *int           `json:"retention_months,omitempty" db:"retention_months"`
	Action          string         `json:"action" db:"action"`
	IsEnabled       bool           `json:"is_enabled" db:"is_enabled"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Dashboard struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	TenantID        uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	Name            string         `json:"name" db:"name"`
	Description     *string        `json:"description,omitempty" db:"description"`
	Layout          []interface{}  `json:"layout" db:"layout"`
	RefreshInterval int            `json:"refresh_interval" db:"refresh_interval"`
	IsDefault       bool           `json:"is_default" db:"is_default"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

type QueryHistory struct {
	ID              uuid.UUID      `json:"id" db:"id"`
	TenantID        uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	UserID          string         `json:"user_id" db:"user_id"`
	QueryType       string         `json:"query_type" db:"query_type"`
	QueryText       string         `json:"query_text" db:"query_text"`
	Filters         map[string]interface{} `json:"filters" db:"filters"`
	ExecutionTimeMs *int           `json:"execution_time_ms,omitempty" db:"execution_time_ms"`
	ResultCount     *int           `json:"result_count,omitempty" db:"result_count"`
	CreatedAt       time.Time      `json:"created_at" db:"created_at"`
}

type IngestionToken struct {
	ID            uuid.UUID      `json:"id" db:"id"`
	TenantID      uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	Name          string         `json:"name" db:"name"`
	TokenHash     string         `json:"token_hash" db:"token_hash"`
	TokenPrefix   string         `json:"token_prefix" db:"token_prefix"`
	Permissions   []string       `json:"permissions" db:"permissions"`
	ExpiresAt     *time.Time     `json:"expires_at,omitempty" db:"expires_at"`
	LastUsedAt    *time.Time     `json:"last_used_at,omitempty" db:"last_used_at"`
	IsActive      bool           `json:"is_active" db:"is_active"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
}

type AuditEvent struct {
	ID           uuid.UUID      `json:"id" db:"id"`
	TenantID     uuid.UUID      `json:"tenant_id" db:"tenant_id"`
	UserID       string         `json:"user_id" db:"user_id"`
	Action       string         `json:"action" db:"action"`
	ResourceType string         `json:"resource_type" db:"resource_type"`
	ResourceID   *uuid.UUID     `json:"resource_id,omitempty" db:"resource_id"`
	Details      map[string]interface{} `json:"details" db:"details"`
	IPAddress    *string        `json:"ip_address,omitempty" db:"ip_address"`
	UserAgent    *string        `json:"user_agent,omitempty" db:"user_agent"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
}

type TenantRepository interface {
	Create(tenant *Tenant) error
	GetByID(id uuid.UUID) (*Tenant, error)
	GetBySlug(slug string) (*Tenant, error)
	Update(tenant *Tenant) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]Tenant, int64, error)
}

type ApplicationRepository interface {
	Create(app *Application) error
	GetByID(id uuid.UUID) (*Application, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]Application, int64, error)
	Update(app *Application) error
	Delete(id uuid.UUID) error
}

type ServiceRepository interface {
	Create(service *Service) error
	GetByID(id uuid.UUID) (*Service, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]Service, int64, error)
	Update(service *Service) error
	Delete(id uuid.UUID) error
}

type EnvironmentRepository interface {
	Create(env *Environment) error
	GetByID(id uuid.UUID) (*Environment, error)
	GetByTenantID(tenantID uuid.UUID) ([]Environment, error)
	Update(env *Environment) error
	Delete(id uuid.UUID) error
}

type AlertRuleRepository interface {
	Create(rule *AlertRule) error
	GetByID(id uuid.UUID) (*AlertRule, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]AlertRule, int64, error)
	GetEnabledByTenantID(tenantID uuid.UUID) ([]AlertRule, error)
	Update(rule *AlertRule) error
	Delete(id uuid.UUID) error
}

type AlertIncidentRepository interface {
	Create(incident *AlertIncident) error
	GetByID(id uuid.UUID) (*AlertIncident, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]AlertIncident, int64, error)
	GetByRuleID(ruleID uuid.UUID) ([]AlertIncident, error)
	Update(incident *AlertIncident) error
}

type RetentionPolicyRepository interface {
	Create(policy *RetentionPolicy) error
	GetByID(id uuid.UUID) (*RetentionPolicy, error)
	GetByTenantID(tenantID uuid.UUID) ([]RetentionPolicy, error)
	GetByType(tenantID uuid.UUID, policyType string) (*RetentionPolicy, error)
	Update(policy *RetentionPolicy) error
	Delete(id uuid.UUID) error
}

type DashboardRepository interface {
	Create(dashboard *Dashboard) error
	GetByID(id uuid.UUID) (*Dashboard, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]Dashboard, int64, error)
	Update(dashboard *Dashboard) error
	Delete(id uuid.UUID) error
}

type QueryHistoryRepository interface {
	Create(history *QueryHistory) error
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]QueryHistory, int64, error)
	GetByUserID(userID string, limit, offset int) ([]QueryHistory, int64, error)
}

type IngestionTokenRepository interface {
	Create(token *IngestionToken) error
	GetByID(id uuid.UUID) (*IngestionToken, error)
	GetByTokenHash(hash string) (*IngestionToken, error)
	GetByTenantID(tenantID uuid.UUID) ([]IngestionToken, error)
	Update(token *IngestionToken) error
	Delete(id uuid.UUID) error
}

type AuditEventRepository interface {
	Create(event *AuditEvent) error
	GetByID(id uuid.UUID) (*AuditEvent, error)
	GetByTenantID(tenantID uuid.UUID, limit, offset int) ([]AuditEvent, int64, error)
	GetByResourceType(resourceType string, limit, offset int) ([]AuditEvent, int64, error)
}
