package main

import (
	"fmt"
	"github.com/IanChenTTT/loxGO/internal/tool"
	"os"
)

func main() {
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) != 2 {
		panic("usage ./genAST relativePathToFolder")
	}
	tool.GenAST(os.Args[1])
}
