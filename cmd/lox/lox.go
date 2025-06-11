package main

import (
	"fmt"
	"github.com/IanChenTTT/loxGO/internal/lox"
	"os"
)

func main() {
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) > 2 {
		panic("Usage: ./lox [script].lox")
	} else if len(os.Args) == 2 {
		lox.RunFile(os.Args[1])
	} else {
		lox.RunPrompt()
	}
}
