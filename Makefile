BUILD_DIR=./builds
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)
APP=lox
BIN= $(BUILD_DIR)/$(APP)

# if func  init have execute order
# rearrange the build target order
# EXAMPLE
TARGET= ./global.go ./error.go ./tokenType.go ./token.go ./scanner.go ./main.go
.PHONY: all build run clean test fmt lint

all: run
build: $(BUILD_DIR)
	go build -o $(BIN) $(TARGET)
run: build
	./$(BIN)
clean:
	go clean
	rm -rf $(BUILD_DIR)
