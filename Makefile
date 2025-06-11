BUILD_DIR=./builds
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)
APP=lox
AST=genAST
BIN_APP= $(BUILD_DIR)/$(APP)
BIN_AST= $(BUILD_DIR)/$(AST)

# if func  init have execute order
# rearrange the build target order
# EXAMPLE
MAIN_TARGET= ./cmd/lox/lox.go
AST_TARGET= ./cmd/tool/generateAST.go
.PHONY: all build run clean test fmt lint ast
	
# MAIN APP
all: run
build: $(BUILD_DIR)
	go build -o $(BIN_APP) $(MAIN_TARGET)
run: build
	./$(BIN_APP)
# internal took for generate AST
ast: $(BUILD_DIR)
	go build -o $(BIN_AST) $(AST_TARGET)
clean:
	go clean
	rm -rf $(BUILD_DIR)
