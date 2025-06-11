BUILD_DIR=./builds
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)
APP=lox
BIN= $(BUILD_DIR)/$(APP)

# if func  init have execute order
# rearrange the build target order
# EXAMPLE
MAIN_TARGET= ./cmd/lox/lox.go
.PHONY: all build run clean test fmt lint

all: run
build: $(BUILD_DIR)
	go build -o $(BIN) $(MAIN_TARGET)
run: build
	./$(BIN)
clean:
	go clean
	rm -rf $(BUILD_DIR)
