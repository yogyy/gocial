include .env
MIGRATION_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create 
migration: 
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: 
migrate-up: 
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_URL) up

.PHONY: 
migrate-down: 
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_URL) down $(filter-out $@,$(MAKECMDGOALS))
