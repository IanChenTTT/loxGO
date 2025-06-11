package lox
import (
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
)

type Expr interface {
	ExprNode()
}
type Binary struct {
  left Expr 
  operator t.Token 
  right Expr  
  }
func ( b *Binary ) Binary(in Binary) {
  b.left=in.left
  b.operator=in.operator
  b.right=in.right
  }
func (b *Binary) ExprNode() {
}
type Grouping struct {
  expression Expr 
  }
func ( g *Grouping ) Grouping(in Grouping) {
  g.expression=in.expression
  }
func (g *Grouping) ExprNode() {
}
type Literal struct {
  value any 
  }
func ( l *Literal ) Literal(in Literal) {
  l.value=in.value
  }
func (l *Literal) ExprNode() {
}
type Unary struct {
  operator t.Token 
  right Expr 
  }
func ( g *Unary ) Unary(in Unary) {
  g.operator=in.operator
  g.right=in.right
  }
func (g *Unary) ExprNode() {
}
