# Building interpreter

## Build

- run lox executable 
    1. run cmd mode
    ```bash
        make run 
    ```
    2. build
    ```bash
        make build 
    ```
    3. run file
    ```bash
        make build && ./builds/lox file.lox
    ```
-  generate AST node file
    1. run
        - recommend path ./internal/lox/ast/ast.go 
    ```bash
    ./ast.sh ./internal/lox/YourPath.go
    ```
    2. build
    ```bash
    make ast 
    ```
- clean builds folder // TODO fix build folder 
    1.
    ```bash
    make clean 
    ```

## Channel Log

1. CH4 rough scanner finish, add additinal float, char , multi comment support
2. CH5 working ...

## TODO

1. add test suite for internal, eg: internal/token/_test.go 

# # Note

1. [ch3](./note/ch3.md)
2. [ch4](./note/ch4.md)

