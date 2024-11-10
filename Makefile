dep: 
	go mod tidy

run: 
	go run ./cmd/main.go

build: 
	go build -o main main.go

run-build: build
	./cmd/main

test:
	go test -v ./tests

init-docker:
	docker compose up -d --build

up: 
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

DB_URL=postgresql://admin:password@0.0.0.0:5432/postgres?sslmode=disable

MIGRATION_PATH=./migrations

MIGRATE_CMD=migrate -path $(MIGRATION_PATH) -database "$(DB_URL)"

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down

migrate-create:
	$ migrate create -ext sql -dir ${MIGRATION_PATH} create_users_table


	migrate $(DB_URL) -f ./migrations/001_create_users_table.down.sql



.PHONY: migrate-up migrate-down migrate-create
