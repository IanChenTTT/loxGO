package lox

import (
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
)

type Expr interface {
	ExprNode()
	accept(visitor Visitor[any]) any
}
type Visitor[T any] interface {
	visitBinary(expr Binary) T
	visitGrouping(expr Grouping) T
	visitLiteral(expr Literal) T
	visitUnary(expr Unary) T
}
type Binary struct {
	left     Expr
	operator t.Token
	right    Expr
}

func (b *Binary) Binary(in Binary) {
	b.left = in.left
	b.operator = in.operator
	b.right = in.right
}

func (b *Binary) ExprNode() {
}
func (b *Binary) accept(visitor Visitor[any]) any {
	visitor.visitBinary(*b)
	return *b
}

type Grouping struct {
	expression Expr
}

func (g *Grouping) Grouping(in Grouping) {
	g.expression = in.expression
}

func (g *Grouping) ExprNode() {
}
func (g *Grouping) accept(visitor Visitor[any]) any {
	visitor.visitGrouping(*g)
	return *g
}

type Literal struct {
	value any
}

func (l *Literal) Literal(in Literal) {
	l.value = in.value
}

func (l *Literal) ExprNode() {
}
func (l *Literal) accept(visitor Visitor[any]) any {
	visitor.visitLiteral(*l)
	return *l
}

type Unary struct {
	operator t.Token
	right    Expr
}

func (g *Unary) Unary(in Unary) {
	g.operator = in.operator
	g.right = in.right
}

func (g *Unary) ExprNode() {
}
func (g *Unary) accept(visitor Visitor[any]) any {
	visitor.visitUnary(*g)
	return *g
}
