.PHONY: all build run clean

BINARY := main
SRC := main.go
BIN_DIR := bin

all: run

build: $(BIN_DIR)/$(BINARY)

$(BIN_DIR)/$(BINARY): $(SRC)
	@mkdir -p $(BIN_DIR)
	@go build -o $@ $<

run: build
	@./$(BIN_DIR)/$(BINARY)

clean:
	@rm -rf $(BIN_DIR)