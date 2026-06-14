.PHONY: help build up down logs test lint

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build all Docker services
	docker compose -f deploy/docker-compose.yml build

up: ## Start all services
	docker compose -f deploy/docker-compose.yml up -d

down: ## Stop all services
	docker compose -f deploy/docker-compose.yml down

logs: ## View service logs
	docker compose -f deploy/docker-compose.yml logs -f

test: ## Run all tests
	cd backend && go test ./...
	cd frontend && npm run lint

lint: ## Run linters
	cd backend && go fmt ./...
	cd backend && go vet ./...
	cd frontend && npm run lint

backend-run: ## Run backend API locally
	cd backend && go run cmd/api/main.go

collector-run: ## Run collector locally
	cd backend && go run cmd/collector/main.go

ingestor-run: ## Run ingestor locally
	cd backend && go run cmd/ingestor/main.go

alert-worker-run: ## Run alert worker locally
	cd backend && go run cmd/alert-worker/main.go

frontend-run: ## Run frontend locally
	cd frontend && npm run dev

clean: ## Remove build artifacts
	docker compose -f deploy/docker-compose.yml down -v
	rm -rf backend/bin frontend/.next frontend/node_modules
