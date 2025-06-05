BUILD_DIR=./builds
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)
APP=lox
BIN= $(BUILD_DIR)/$(APP)

.PHONY: all build run clean test fmt lint

all: run
build: $(BUILD_DIR)
	go build -o $(BIN) .
run: build
	./$(BIN)
clean:
	go clean
	rm -rf $(BUILD_DIR)
