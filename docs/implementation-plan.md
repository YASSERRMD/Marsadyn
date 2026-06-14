# Marsadyn Implementation Plan

## Project Overview

Marsadyn is a production-grade distributed observability platform supporting metrics, logs, traces, alerting, and high-throughput ingestion.

## Phase 0: Repository Inspection and Planning
**Status**: In Progress

| Task | Status | Commit |
|------|--------|--------|
| Inspect repository | Done | - |
| Document architecture | Done | phase 0: add architecture overview |
| Create implementation plan | In Progress | - |
| Add git workflow docs | Pending | - |
| Add README with attribution | Pending | - |

## Phase 1: Project Foundation
**Status**: Pending

| Task | Description |
|------|-------------|
| Initialize Go module | `go mod init github.com/marsadyn/marsadyn` |
| Add backend structure | cmd/, internal/, migrations/ |
| Add service commands | api, collector, ingestor, alert-worker |
| Initialize Next.js | Frontend app with TypeScript |
| Add Dockerfiles | Backend and frontend |
| Add Docker Compose | PostgreSQL, Redis, Kafka, ClickHouse |
| Add Makefile | Developer commands |
| Add health endpoint | Backend health check |

## Phase 2: Core Data Model
**Status**: Pending

| Entity | Purpose |
|--------|---------|
| Tenant | Multi-tenancy support |
| Application | Application grouping |
| Service | Service metadata |
| Environment | Environment separation |
| MetricSeries | Metric metadata |
| LogStream | Log stream metadata |
| TraceService | Trace service metadata |
| AlertRule | Alert configuration |
| AlertIncident | Alert incidents |
| RetentionPolicy | Data retention |
| Dashboard | Dashboard config |
| QueryHistory | Query audit |
| IngestionToken | API authentication |
| AuditEvent | Audit logging |

## Phase 3: Telemetry Event Contracts
**Status**: Pending

| Contract | Format |
|----------|--------|
| Metric Event | JSON with tenantId, service, name, value, labels |
| Log Event | JSON with tenantId, service, level, message |
| Trace Span | JSON with tenantId, service, spanId, parentSpanId |

## Phase 4: Kafka Ingestion Pipeline
**Status**: Pending

| Topic | Purpose |
|-------|---------|
| marsadyn.metrics.raw | Raw metrics |
| marsadyn.logs.raw | Raw logs |
| marsadyn.traces.raw | Raw traces |
| marsadyn.metrics.validated | Validated metrics |
| marsadyn.logs.validated | Validated logs |
| marsadyn.traces.validated | Validated traces |
| marsadyn.deadletter | Failed validations |

## Phase 5: Storage Engine
**Status**: Pending

| Component | Technology |
|-----------|------------|
| Telemetry Storage | ClickHouse |
| Metadata Storage | PostgreSQL |
| Batch Writer | Go routine with buffer |
| Retry Strategy | Exponential backoff |

## Phase 6: Query Engine
**Status**: Pending

| Endpoint | Description |
|----------|-------------|
| GET /api/v1/query/metrics | Query metrics |
| GET /api/v1/query/logs | Query logs |
| GET /api/v1/query/traces | Query traces |
| POST /api/v1/query/execute | Execute custom query |

## Phase 7: Alerting Engine
**Status**: Pending

| Component | Description |
|-----------|-------------|
| Alert Rules | Threshold, log pattern, trace latency |
| Evaluator | Scheduled rule evaluation |
| State Cache | Redis for alert state |
| Incidents | Incident creation and resolution |

## Phase 8: Retention and Data Lifecycle
**Status**: Pending

| Worker | Purpose |
|--------|---------|
| Metrics Retention | Archive/delete old metrics |
| Logs Retention | Archive/delete old logs |
| Traces Retention | Archive/delete old traces |

## Phase 9: Dashboard Frontend Foundation
**Status**: Pending

| Page | Description |
|------|-------------|
| /dashboard | Overview dashboard |
| /metrics | Metrics explorer |
| /logs | Logs explorer |
| /traces | Trace explorer |
| /alerts | Alert rules |
| /incidents | Active incidents |
| /retention | Retention policies |

## Phase 10: Dashboard Features
**Status**: Pending

| Feature | Description |
|---------|-------------|
| Metrics Explorer | Query and visualize metrics |
| Logs Explorer | Search and filter logs |
| Trace Explorer | View trace details |
| Alert Management | Create and manage alerts |

## Phase 11: Sample Applications
**Status**: Pending

| Component | Description |
|-----------|-------------|
| Sample App | Go application with instrumentation |
| Telemetry Generator | Generate test data |
| Error Simulation | Simulate errors and latency |

## Phase 12: Security and Multi-Tenancy
**Status**: Pending

| Feature | Description |
|---------|-------------|
| Token Validation | Ingestion token auth |
| Tenant Resolution | Middleware for tenant context |
| Rate Limiting | API rate limiting |
| Audit Logging | Track API operations |

## Phase 13: Documentation and Polish
**Status**: Pending

| Task | Description |
|------|-------------|
| README update | Project overview and setup |
| Attribution | OpenCode + Xiaomi MiMo-V2.5 |
| Architecture diagrams | Visual documentation |
| API examples | curl examples |
| Roadmap | Future plans |

## Timeline

| Phase | Duration | Focus |
|-------|----------|-------|
| 0 | 1 hour | Planning |
| 1 | 2 hours | Foundation |
| 2 | 2 hours | Data Model |
| 3 | 1 hour | Contracts |
| 4 | 3 hours | Kafka Pipeline |
| 5 | 2 hours | Storage |
| 6 | 2 hours | Query Engine |
| 7 | 3 hours | Alerting |
| 8 | 2 hours | Retention |
| 9 | 3 hours | Frontend Base |
| 10 | 3 hours | Frontend Features |
| 11 | 2 hours | Sample Apps |
| 12 | 2 hours | Security |
| 13 | 1 hour | Documentation |

**Total**: ~30 hours
