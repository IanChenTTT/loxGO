package ast

import (
	"fmt"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
	"strconv"
	"strings"
)

type ASTPrinter struct{}

func NewASTPrinter() *ASTPrinter {
	return &ASTPrinter{}
}

func (astp *ASTPrinter) Print(expr Expr) any {
	return expr.accept(astp)
}
func (ast *ASTPrinter) visitBinary(expr Binary) any {
	return ast.parenthesize(expr.Operator.Lexemes, expr.Left, expr.Right)
}
func (ast *ASTPrinter) visitGrouping(expr Grouping) any {
	return ast.parenthesize("group", expr.Expr)
}
func (ast *ASTPrinter) visitLiteral(expr Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	fmt.Printf("%T type and value %v", expr.Value, expr.Value)
	switch expr.Value.(type) {
	case string:
		str, _ := expr.Value.(string)
		return str
	case float64:
		f, _ := expr.Value.(float64)
		s := strconv.FormatFloat(f, 'E', -1, 32)
		return s
	case int:
		fmt.Println("it should enter here1")
		s := strconv.Itoa(expr.Value.(int))
		return s
	case int32:
		fmt.Println("it should enter here2")
		in, _ := expr.Value.(rune)
		return string(in)
	}
	return ""
}
func (ast *ASTPrinter) visitUnary(expr Unary) any {
	return ast.parenthesize(expr.Operator.Lexemes, expr.Right)
}
func (ast *ASTPrinter) parenthesize(name string, exprs ...Expr) string {
	var build strings.Builder
	fmt.Fprintf(&build, "(%s", name)
	for _, expr := range exprs {
		fmt.Fprintf(&build, " ")
		s, ok := expr.accept(ast).(string)
		if !ok {
			fmt.Printf("error parenthesize %v %T\n", expr, expr)
			return ""
		}
		fmt.Fprintf(&build, s)
	}
	fmt.Fprintf(&build, ")")
	return build.String()
}

// test function
func Run() {
	var ex Expr
	ex = &Binary{
		Left: &Unary{
			Operator: t.Token{
				Types:   t.MINUS,
				Lexemes: "-",
				Literal: nil,
				Line:    0,
			},
			Right: &Literal{
				Value: 123,
			},
		},
		Operator: t.Token{
			Types:   t.STAR,
			Lexemes: "*",
			Literal: nil,
			Line:    0,
		},
		Right: &Grouping{
			Expr: &Literal{Value: 45.67},
		},
	}
	ast := NewASTPrinter()
	fmt.Println(ast.Print(ex))
}
