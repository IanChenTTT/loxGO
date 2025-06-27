package parser

import (
	ast "github.com/IanChenTTT/loxGO/internal/lox/ast"
	g "github.com/IanChenTTT/loxGO/internal/lox/global"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"
)

type Parser struct {
	tokens  []t.Token
	current int
}
type parserError struct {
	s string
}

/*
lowest priority to highest priority
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
func (p *Parser) Run() ast.Expr {
	return p.expression()
}

func (p *Parser) expression() ast.Expr {
	return p.equality()
}

//BINARY

// common operation stand for common binary operation
func (p *Parser) common(fn func() ast.Expr, typs ...t.TokenType) ast.Expr {
	var expr ast.Expr
	expr = fn() // start from left expression
	for p.match(typs...) {
		tok := p.previous()
		var right ast.Expr
		right = fn()
		expr = &ast.Binary{
			Left:     expr, // this append previous exprion aka left expression
			Operator: tok,
			Right:    right,
		}
	}
	return expr
}

// comparison(...)* this would keep matching (==|!=)
// ((==|!=)comparison)* => ==a==b like this
func (p *Parser) equality() ast.Expr {
	return p.common(p.comparison, t.BANG_EQUAL, t.EQUAL_EQUAL)
}

// term(...)* this would keep matching (>|>=|<|<=)
// ((>|>=|<|<=)term)* => > a > b like this
func (p *Parser) comparison() ast.Expr {
	return p.common(p.term, t.GREATER, t.GREATER_EQUAL, t.LESS, t.LESS_EQUAL)
}

// factor(...)* this would keep matching (-|+)
// ((-|+)factor)* => - b + a like this
func (p *Parser) term() ast.Expr {
	return p.common(p.factor, t.MINUS, t.PLUS)
}

// unary(...)* this would keep matching (*|/)
// ((*|/)unary)* => * b / a like this
func (p *Parser) factor() ast.Expr {
	return p.common(p.unary, t.SLASH, t.STAR)
}

// unary will recursive find - or !
// or not found it just return terminal literal
func (p *Parser) unary() ast.Expr {
	if p.match(t.BANG, t.MINUS) {
		tok := p.previous()
		var right ast.Expr
		right = p.unary()
		return &ast.Unary{
			Operator: tok,
			Right:    right,
		}
	}
	return p.primary()
}

// primary is last rule that match the terminal
func (p *Parser) primary() ast.Expr {
	var expr ast.Expr
	if p.match(t.FALSE) {
		expr = &ast.Literal{Value: false}
	} else if p.match(t.TRUE) {
		expr = &ast.Literal{Value: true}
	} else if p.match(t.NIL) {
		expr = &ast.Literal{Value: nil}
	} else if p.match(t.INT, t.STRING, t.CHAR, t.FLOAT) {
		expr = &ast.Literal{
			Value: p.previous().Literal,
		}
	} else if p.match(t.LEFT_PAREN) {
		expr = p.expression()
		tok, eState := p.consume(t.RIGHT_PAREN, "Expect ')' after expression")
		if eState.HadError {
			New(tok, eState.S)
		}
	}
	return expr
}

//
// ERROR
//

// consume is a parser checker make sure expression enclose
func (p *Parser) consume(typ t.TokenType, msg string) (t.Token, g.ErrState) {
	var eState g.ErrState
	eState.HadError = false
	if p.check(typ) {
		return p.advance(), eState
	}
	// error occur
	eState.HadError = true
	eState.S = msg
	return t.Token{}, eState
}

//
// UTIL
//

// match series of tokenType to current token type
func (p *Parser) match(types ...t.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	for _, typ := range types {
		if p.check(typ) {
			p.advance()
			return true
		}
	}
	return false
}

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
func (p *Parser) check(tok t.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.tokens[p.current].Types == tok
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

//
// ERROR
//

func (e *parserError) Error() string {
	return e.s
}
func New(tok t.Token, msg string) error {
	var eState g.ErrState

	// just report error back to user
	eState.ErnoToken(tok, msg)

	// then keep tracking what wrong
	return &parserError{msg} //TODO
}
