package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var eState errState

func init() {

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func runFile(path string) {
	var estate errState
	fil, err := os.Open(path)
	check(err)
	stat, err := fil.Stat()
	check(err)
	buf := make([]byte, stat.Size())

	n, err := io.ReadFull(bufio.NewReader(fil), buf)
	check(err)
	fmt.Printf("%d bytes %s", n, string(buf[:n]))

	err = fil.Close()
	check(err)
	estate = run(buf[:n]) //TODO
	if !estate.hadError {
		panic("run return error")
	}
}
func runPrompt() {
	var state errState
	red := bufio.NewReader(os.Stdin)
	fmt.Println("Enter expression, ctrl-d to end of file:")
	var buf []byte
	for {
		fmt.Print(">")
		text, err := red.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF")
				state = run(buf)
				state.hadError = false //TODO
				break
			}
			check(err)
		}
		buf = append(buf, text...)
	}
}
func run(src []byte) errState {
	var scanner Scanner
	scanner.Scanner(string(src))
	tokens, eState := scanner.scanTokens()
	for i, tok := range tokens {
		fmt.Printf("scan: %d tok: %v", i, tok)
	}
	return eState
}
func main() {
	fmt.Println(len(os.Args), os.Args)
	if len(os.Args) > 2 {
		panic("Usage: jlox [script]")
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
