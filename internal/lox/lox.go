package lox

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/IanChenTTT/loxGO/internal/lox/ast"
	g "github.com/IanChenTTT/loxGO/internal/lox/global"
	"github.com/IanChenTTT/loxGO/internal/lox/parser"
	s "github.com/IanChenTTT/loxGO/internal/lox/scanner"
)

// check is a private error checker
// which implied  system error, not user error
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func RunFile(path string) {
	var estate g.ErrState
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
	estate = run(&buf) //TODO
	if estate.HadError {
		panic("run return error")
	}
}
func RunPrompt() {
	var state g.ErrState
	red := bufio.NewReader(os.Stdin)
	fmt.Println("Enter expression, ctrl-d twice to end of file:")
	var buf []byte
	for {
		fmt.Print(">")
		text, err := red.ReadBytes('\n')
		buf = append(buf, text...)
		if err != nil {
			if err == io.EOF {
				fmt.Println() // flush buffer
				state = run(&buf)
				state.HadError = false //TODO
				break
			}
			check(err)
		}
	}
}
func run(src *[]byte) g.ErrState {
	var scanner s.Scanner
	scanner.Scanner(string(*src))
	eState := scanner.ScanTokens()
	for i, tok := range scanner.Tokens {
		fmt.Printf("scan: %d  %s\n", i, tok.ToString())
	}
	if len(scanner.Tokens) == 1 {
		return eState //  EOF handle TODO better EOR handle
	}
	parsed := parser.NewParser(scanner.Tokens)
	astPrint := ast.NewASTPrinter()
	expr, eState := parsed.Run()
	if !eState.HadError {
		fmt.Println(astPrint.Print(expr))
	}
	return eState
}
