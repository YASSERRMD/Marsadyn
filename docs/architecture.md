# Marsadyn Architecture

## Overview

Marsadyn is a production-grade distributed observability platform that supports metrics, logs, traces, alerting, and high-throughput ingestion.

## System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Applications                           │
│                    (Sample Apps, SDKs)                       │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                     Collectors                              │
│              (HTTP Ingestion Endpoints)                     │
│    POST /api/v1/ingest/metrics                             │
│    POST /api/v1/ingest/logs                                │
│    POST /api/v1/ingest/traces                              │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    Kafka Cluster                            │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐          │
│  │marsadyn.    │ │marsadyn.    │ │marsadyn.    │          │
│  │metrics.raw  │ │logs.raw     │ │traces.raw   │          │
│  └──────┬──────┘ └──────┬──────┘ └──────┬──────┘          │
│         │               │               │                   │
│  ┌──────▼──────┐ ┌──────▼──────┐ ┌──────▼──────┐          │
│  │marsadyn.    │ │marsadyn.    │ │marsadyn.    │          │
│  │metrics.     │ │logs.        │ │traces.      │          │
│  │validated    │ │validated    │ │validated    │          │
│  └──────┬──────┘ └──────┬──────┘ └──────┬──────┘          │
│         │               │               │                   │
│  ┌──────▼───────────────▼───────────────▼──────┐          │
│  │         marsadyn.deadletter                   │          │
│  └──────────────────────────────────────────────┘          │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                  Ingestion Workers                          │
│            (Validation, Transformation)                     │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    Storage Engine                           │
│  ┌─────────────────┐  ┌─────────────────┐                  │
│  │  ClickHouse     │  │  PostgreSQL     │                  │
│  │  (Telemetry)    │  │  (Metadata)     │                  │
│  └─────────────────┘  └─────────────────┘                  │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                    Query Engine                             │
│          (Metrics, Logs, Traces Queries)                    │
└──────────────────────────┬──────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────────────┐
│                   Presentation Layer                        │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐          │
│  │ Dashboards  │ │ Alerting    │ │ API         │          │
│  │ (Next.js)   │ │ Engine      │ │ Endpoints   │          │
│  └─────────────┘ └─────────────┘ └─────────────┘          │
└─────────────────────────────────────────────────────────────┘
```

## Components

### 1. Collectors
- HTTP endpoints for telemetry ingestion
- Accept metrics, logs, and traces
- Validate incoming data
- Publish to Kafka topics

### 2. Kafka Pipeline
- **Raw Topics**: Receive unvalidated telemetry
- **Validated Topics**: Store validated telemetry
- **Dead Letter Topic**: Handle failed validations

### 3. Ingestion Workers
- Consume from Kafka topics
- Validate and transform telemetry
- Write to storage engine

### 4. Storage Engine
- **ClickHouse**: Time-series telemetry storage
- **PostgreSQL**: Metadata, users, configuration

### 5. Query Engine
- Time-range queries
- Label filtering
- Aggregation functions
- Pagination support

### 6. Alerting Engine
- Rule-based evaluation
- Threshold alerts
- Log pattern alerts
- Trace latency alerts
- Incident management

### 7. Frontend (Next.js)
- Dashboard overview
- Metrics explorer
- Logs explorer
- Trace explorer
- Alert management
- Settings

## Data Flow

1. **Ingestion**: Applications send telemetry to collectors
2. **Validation**: Collectors validate and publish to Kafka
3. **Processing**: Workers consume and process telemetry
4. **Storage**: Processed telemetry stored in ClickHouse
5. **Querying**: Query engine retrieves data for dashboards
6. **Alerting**: Alert engine evaluates rules against stored data

## Technology Stack

| Component | Technology |
|-----------|------------|
| Backend | Go |
| Frontend | Next.js, TypeScript, Tailwind CSS |
| Message Queue | Apache Kafka |
| Time-Series DB | ClickHouse |
| Metadata DB | PostgreSQL |
| Cache | Redis |
| Containerization | Docker, Docker Compose |

## Scalability

- **Horizontal scaling**: Multiple collector instances
- **Kafka partitioning**: Parallel processing
- **ClickHouse sharding**: Distributed storage
- **Worker scaling**: Independent ingestion workers

## Reliability

- **Dead letter queue**: Failed message handling
- **Retry mechanisms**: Automatic retry with backoff
- **Health checks**: Service monitoring
- **Circuit breaking**: Fault tolerance
