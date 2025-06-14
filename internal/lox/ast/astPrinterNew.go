package ast

import (
	"fmt"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
	"strconv"
	"strings"
)

type ASTPrinter struct{}

func (astp *ASTPrinter) print(expr Expr) any {
	return expr.accept(astp)
}
func (ast *ASTPrinter) visitBinary(expr Binary) any {
	return ast.parenthesize(expr.operator.Lexemes, expr.left, expr.right)
}
func (ast *ASTPrinter) visitGrouping(expr Grouping) any {
	return ast.parenthesize("group", expr.expression)
}
func (ast *ASTPrinter) visitLiteral(expr Literal) any {
	if expr.value == nil {
		return "nil"
	}
	switch expr.value.(type) {
	case string:
		str, _ := expr.value.(string)
		return str
	case float64:
		f, _ := expr.value.(float64)
		s := strconv.FormatFloat(f, 'E', -1, 32)
		return s
	case int:
		s := strconv.Itoa(expr.value.(int))
		return s
	case int32:
		in, _ := expr.value.(rune)
		return string(in)
	}
	return ""
}
func (ast *ASTPrinter) visitUnary(expr Unary) any {
	return ast.parenthesize(expr.operator.Lexemes, expr.right)
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
		left: &Unary{
			operator: t.Token{
				Types:   t.MINUS,
				Lexemes: "-",
				Literal: nil,
				Line:    0,
			},
			right: &Literal{
				value: 123,
			},
		},
		operator: t.Token{
			Types:   t.STAR,
			Lexemes: "*",
			Literal: nil,
			Line:    0,
		},
		right: &Grouping{
			expression: &Literal{value: 45.67},
		},
	}
	ast := ASTPrinter{}
	fmt.Println(ast.print(ex))
}
