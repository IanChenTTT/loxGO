#! /bin/sh

astPrinter="./internal/lox/ast/astPrinter.go"
ast="./internal/lox/ast/ast.go"

renameFile() {
    echo "File exists: $1"
    funcName=$2

    # Get directory, filename, and extension
    dir=$(dirname "$1")
    base=$(basename "$1")
    name="${base%.*}"      # Remove extension
    ext="${base##*.}"      # Extract extension

    temp="$dir/$name"      # Path without extension

    # Rename file to remove extension
    mv "$1" "$temp"
    echo "Renamed to: $temp"

    # Do something here (e.g., sleep, compile, edit, etc.)
    $funcName 

    # Rename back to original
    mv "$temp" "$1"
    echo "Renamed back to: $1"
}
astBuild(){
  # if astPrinter no need
  # rebuild prog if necessary
  make ast
  # run prog with some arguments
  ./builds/genAST $1 
}
# if you create new Expr type , kill old ast.go 
if [ -f "$astPrinter" ] && [ ! -f "$ast" ]; then 
  renameFile $astPrinter astBuild
  exit 1
fi

astBuild "$@"


