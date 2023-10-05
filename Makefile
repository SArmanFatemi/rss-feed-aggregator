# Load the .env file into the current shell session
include .env

MIGRATION_PATH=database/migrations

all: start

test:
	@echo $(DB_URL)

build:
	@go build

start: build
	@./rssagg.exe

migration:
	@goose -dir="$(MIGRATION_PATH)" postgres $(DB_URL) up

migration-down:
	@goose -dir="$(MIGRATION_PATH)" postgres $(DB_URL) down

sqlc:
	@docker run --rm -v ${PWD}:/src -w /src sqlc/sqlc generate
