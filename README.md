# Marsadyn

**Built using OpenCode with Xiaomi MiMo-V2.5.**

## Overview

Marsadyn is a production-grade distributed observability platform that supports:

- **Metrics**: Time-series data collection and visualization
- **Logs**: Centralized log aggregation and search
- **Traces**: Distributed tracing and service maps
- **Alerting**: Rule-based alert evaluation and incident management
- **High-throughput ingestion**: Kafka-based telemetry pipeline

## Architecture

```
Applications → Collectors → Kafka → Ingestion Workers → Storage Engine → Query Engine → Dashboards + Alerting
```

## Tech Stack

| Component | Technology |
|-----------|------------|
| Backend | Go |
| Frontend | Next.js, TypeScript, Tailwind CSS |
| Message Queue | Apache Kafka |
| Time-Series DB | ClickHouse |
| Metadata DB | PostgreSQL |
| Cache | Redis |
| Containerization | Docker, Docker Compose |

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+
- Node.js 18+

### Clone and Run

```bash
git clone https://github.com/your-org/marsadyn.git
cd marsadyn

# Start infrastructure
docker compose -f deploy/docker-compose.yml up -d

# Run backend
cd backend
go run cmd/api/main.go

# Run frontend
cd frontend
npm install
npm run dev
```

### Access Services

| Service | URL |
|---------|-----|
| Frontend | http://localhost:3000 |
| API | http://localhost:8080 |
| Kafka UI | http://localhost:9092 |
| ClickHouse | http://localhost:8123 |

## Project Structure

```
marsadyn/
├── backend/           # Go backend services
│   ├── cmd/          # Service entry points
│   ├── internal/     # Private packages
│   └── migrations/   # Database migrations
├── frontend/         # Next.js web console
├── deploy/           # Docker and deployment
├── docs/             # Documentation
└── examples/         # Sample applications
```

## API Endpoints

### Ingestion

```bash
# Metrics
curl -X POST http://localhost:8080/api/v1/ingest/metrics \
  -H "Content-Type: application/json" \
  -d '{"tenantId":"tenant-001","service":"my-api","name":"requests_total","value":1}'

# Logs
curl -X POST http://localhost:8080/api/v1/ingest/logs \
  -H "Content-Type: application/json" \
  -d '{"tenantId":"tenant-001","service":"my-api","level":"info","message":"Request processed"}'

# Traces
curl -X POST http://localhost:8080/api/v1/ingest/traces \
  -H "Content-Type: application/json" \
  -d '{"tenantId":"tenant-001","service":"my-api","spanId":"span-1","name":"http.request"}'
```

### Querying

```bash
# Query metrics
curl "http://localhost:8080/api/v1/query/metrics?start=2026-01-01T00:00:00Z&end=2026-01-02T00:00:00Z"

# Query logs
curl "http://localhost:8080/api/v1/query/logs?service=my-api&level=error"

# Query traces
curl "http://localhost:8080/api/v1/query/traces?service=my-api&minDuration=100ms"
```

## Development

### Backend

```bash
cd backend
go fmt ./...
go vet ./...
go test ./...
```

### Frontend

```bash
cd frontend
npm run lint
npm run build
```

## Configuration

Environment variables:

```bash
# Database
DATABASE_URL=postgres://localhost:5432/marsadyn
CLICKHOUSE_URL=http://localhost:8123

# Kafka
KAFKA_BROKERS=localhost:9092

# Redis
REDIS_URL=redis://localhost:6379

# Server
API_PORT=8080
```

## Roadmap

- [ ] Phase 0: Repository setup and planning
- [ ] Phase 1: Project foundation
- [ ] Phase 2: Core data model
- [ ] Phase 3: Telemetry event contracts
- [ ] Phase 4: Kafka ingestion pipeline
- [ ] Phase 5: Storage engine
- [ ] Phase 6: Query engine
- [ ] Phase 7: Alerting engine
- [ ] Phase 8: Retention and data lifecycle
- [ ] Phase 9: Dashboard frontend foundation
- [ ] Phase 10: Dashboard features
- [ ] Phase 11: Sample applications
- [ ] Phase 12: Security and multi-tenancy
- [ ] Phase 13: Documentation and polish

## License

MIT License - see [LICENSE](LICENSE)

## Built With

- **OpenCode** - Open-source coding agent
- **Xiaomi MiMo-V2.5** - AI model for code generation
