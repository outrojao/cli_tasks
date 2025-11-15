.PHONY: all build run test clean

BINARY := main
SRC := ./cmd/main.go
BIN_DIR := bin

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