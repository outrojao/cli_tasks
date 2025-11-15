.PHONY: all build run test clean migrate rollback status create

BINARY := main
SRC := ./cmd/cli/main.go
BIN_DIR := bin
MIGRATIONS_DIR := ./internal/database/migrations
ENV_FILE := ./configs/.env

-include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE) 2>/dev/null)

all: run

build: $(BIN_DIR)/$(BINARY)

$(BIN_DIR)/$(BINARY): $(SRC)
	@mkdir -p $(BIN_DIR)
	@go build -o $@ $<

run: build
	@./$(BIN_DIR)/$(BINARY)

test:
	@go test ./...

clean:
	@rm -rf $(BIN_DIR)

migrate:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

rollback:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down

status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" status

create:
	@if [ -z "$(name)" ]; then \
		echo "Uso: make create name=nome_da_migration"; \
	else \
		goose -dir $(MIGRATIONS_DIR) create $(name) sql; \
	fi
