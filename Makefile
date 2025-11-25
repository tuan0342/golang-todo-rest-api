APP_NAME=todo-go
BIN=bin/server

# Default Go command
GO_CMD=go

# Docker settings
DOCKER_COMPOSE=docker compose
DOCKER_IMAGE=todo-go-image

# Env variables for local run
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=123456
export DB_NAME=mydb
export DB_PORT=5432
export PORT=9090


# =====================================================
# Local development
# =====================================================

.PHONY: run
run:
	@echo ">> Running app locally..."
	$(GO_CMD) run -mod=vendor main.go

.PHONY: build
build:
	@echo ">> Building binary locally..."
	$(GO_CMD) build -mod=vendor -o $(BIN) main.go


# =====================================================
# Modules & vendor
# =====================================================

.PHONY: tidy
tidy:
	@echo ">> Running go mod tidy"
	$(GO_CMD) mod tidy

.PHONY: vendor
vendor:
	@echo ">> Sync vendor folder"
	$(GO_CMD) mod vendor


# =====================================================
# Docker commands
# =====================================================

.PHONY: dev
dev:
	@echo ">> Starting backend and postgres with Docker Compose..."
	$(DOCKER_COMPOSE) up -d --build

.PHONY: down
down:
	@echo ">> Stopping containers..."
	$(DOCKER_COMPOSE) down

.PHONY: down-v
down-v:
	@echo ">> Stopping containers and deleting volumes..."
	$(DOCKER_COMPOSE) down -v

.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f backend

.PHONY: shell
shell:
	@echo ">> Entering backend container..."
	docker exec -it todo-go sh


# =====================================================
# Code tools
# =====================================================

.PHONY: fmt
fmt:
	@echo ">> Formatting code..."
	$(GO_CMD) fmt ./...

.PHONY: test
test:
	@echo ">> Running tests..."
	$(GO_CMD) test ./... -v

.PHONY: clean
clean:
	@echo ">> Cleaning build artifacts..."
	rm -rf $(BIN)
