package ast

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
	Left     Expr
	Operator t.Token
	Right    Expr
}

func (b *Binary) Binary(in Binary) {
	b.Left = in.Left
	b.Operator = in.Operator
	b.Right = in.Right
}

func (b *Binary) ExprNode() {
}
func (b *Binary) accept(visitor Visitor[any]) any {
	return visitor.visitBinary(*b)
}

type Grouping struct {
	Expr Expr
}

func (g *Grouping) Grouping(in Grouping) {
	g.Expr = in.Expr
}

func (g *Grouping) ExprNode() {
}
func (g *Grouping) accept(visitor Visitor[any]) any {
	return visitor.visitGrouping(*g)
}

type Literal struct {
	Value any
}

func (l *Literal) Literal(in Literal) {
	l.Value = in.Value
}

func (l *Literal) ExprNode() {
}
func (l *Literal) accept(visitor Visitor[any]) any {
	return visitor.visitLiteral(*l)
}

type Unary struct {
	Operator t.Token
	Right    Expr
}

func (g *Unary) Unary(in Unary) {
	g.Operator = in.Operator
	g.Right = in.Right
}

func (g *Unary) ExprNode() {
}
func (g *Unary) accept(visitor Visitor[any]) any {
	return visitor.visitUnary(*g)
}
