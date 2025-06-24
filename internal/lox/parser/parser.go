package parser

import (
	ast "github.com/IanChenTTT/loxGO/internal/lox/ast"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
)

type Parser struct {
	tokens  []t.Token
	current int
}

/*
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
*/

//
//MAIN
//

// newParser reutrn new Parser struct
// new copy method no interface
func NewParser(tokens []t.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}
func (p *Parser) Run() {
	p.expression()
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
}
func (p *Parser) equality() ast.Expr {
	//exprComp = comparison()
	return p.equality() //TODO
}
func (p *Parser) comparison() ast.Expr {
	var expr ast.Expr
	return expr //TODO
}
func (p *Parser) term() ast.Expr {
	var expr ast.Expr
	return expr //TODO
}
func (p *Parser) factor() ast.Expr {
	var expr ast.Expr
	return expr //TODO
}
func (p *Parser) unary() ast.Expr {
	var expr ast.Expr
	return expr //TODO
}
func (p *Parser) primary() ast.Expr {
	var expr ast.Expr
	return expr //TODO
}

//
// UTIL
//

// advance check current token is EOF
// if not return current token and
// advance current + 1
// else return token before EOF
func (p *Parser) advance() t.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

// check return tok type match current tok type
func (p *Parser) check(tok t.Token) bool {
	if p.isAtEnd() {
		return false
	}
	return p.tokens[p.current].Types == tok.Types
}

// peek return Parser.current
func (p *Parser) peek() t.Token {
	return p.tokens[p.current]
}

// previous return Parser.current-1
func (p *Parser) previous() t.Token {
	return p.tokens[p.current-1]
}

// isAtEnd return Parser.current is EOF
func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].Types == t.EOF
}
