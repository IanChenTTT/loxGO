#! /bin/sh
# rebuild prog if necessary
make ast
# run prog with some arguments
./builds/genAST "$@"
