.PHONY: build run clean

BINARY_NAME={{BINARY_NAME}}

build:
	@echo "🏗️  Compilando binario..."
	@go build -o bin/$(BINARY_NAME) ./cmd
	@echo "✅ ¡Listo! Ejecutable creado como './bin/$(BINARY_NAME)'"

# Desarrollo con hot-reload
dev:
	@echo "🚀 Iniciando servidor en modo desarrollo..."
	@pnpm dlx nodemon -e go -x "go run cmd/main.go" --signal SIGTERM --ignore '**/*_test.go'

# Ejecutar sin hot-reload
run:
	@echo "▶️  Ejecutando aplicación..."
	@go run cmd/main.go

include .env
export

DB_URL="user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=$(DB_SSLMODE)"
MIGRATIONS_DIR="internal/storage/migrations"

migrate-status:
	@echo "📊 Estado de Migraciones:"
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) status

# Uso: make migrate-create name=create_users_table
migrate-create:
	@echo "🛠️ Creando migración '$(name)'..."
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql
# ejecuta la creación de las tablas
migrate-up:
	@echo "⬆️ Ejecutando migraciones pendientes (Up)..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

migrate-down:
	@echo "⬇️ Deshaciendo la última migración (Down)..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

migrate-reset:
	@echo "🚨 ATENCIÓN: Deshaciendo TODAS las migraciones (Reset)..."
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) reset
