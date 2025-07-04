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
primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" | "(" expression "," expression ")";
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

// Run start recursive parsing from list of token
// TODO error handling
func (p *Parser) Run() (ast.Expr, g.ErrState) {
	var eState g.ErrState
	expr, errs := p.expression() // TODO AST tree error handling
	if errs != nil {
		eState.HadError = true
	}
	return expr, eState
}

func (p *Parser) expression() (ast.Expr, []error) {
	return p.equality()
}

// expression Wrapper return new ast.Expr, and error
// TODO when append error something should happend like New in the error section
func (p *Parser) expressionWrapper(errs *[]error) ast.Expr {
	expr1, err1 := p.expression() //this line sucks
	// case 1 (expression)
	if err1 != nil { // left expression
		*errs = append(*errs, err1...)
		return expr1
	}
	return expr1
}

//BINARY

// common operation stand for common binary operation
func (p *Parser) common(fn func() (ast.Expr, []error), typs ...t.TokenType) (ast.Expr, []error) {
	var expr ast.Expr
	expr, err := fn() // start from left expression
	for p.match(typs...) {
		tok := p.previous()
		var right ast.Expr
		right, err1 := fn()
		err = append(err, err1...)
		expr = &ast.Binary{
			Left:     expr, // this append previous exprion aka left expression
			Operator: tok,
			Right:    right,
		}
	}
	return expr, err
}

// comparison(...)* this would keep matching (==|!=)
// ((==|!=)comparison)* => ==a==b like this
func (p *Parser) equality() (ast.Expr, []error) {
	return p.common(p.comparison, t.BANG_EQUAL, t.EQUAL_EQUAL)
}

// term(...)* this would keep matching (>|>=|<|<=)
// ((>|>=|<|<=)term)* => > a > b like this
func (p *Parser) comparison() (ast.Expr, []error) {
	return p.common(p.term, t.GREATER, t.GREATER_EQUAL, t.LESS, t.LESS_EQUAL)
}

// factor(...)* this would keep matching (-|+)
// ((-|+)factor)* => - b + a like this
func (p *Parser) term() (ast.Expr, []error) {
	return p.common(p.factor, t.MINUS, t.PLUS)
}

// unary(...)* this would keep matching (*|/)
// ((*|/)unary)* => * b / a like this
func (p *Parser) factor() (ast.Expr, []error) {
	return p.common(p.unary, t.SLASH, t.STAR)
}

// unary will recursive find - or !
// or not found it just return terminal literal
func (p *Parser) unary() (ast.Expr, []error) {
	if p.match(t.BANG, t.MINUS) {
		tok := p.previous()
		var right ast.Expr
		right, err := p.unary()
		return &ast.Unary{
			Operator: tok,
			Right:    right,
		}, err
	}
	return p.primary()
}

// primary is last rule that match the terminal
func (p *Parser) primary() (ast.Expr, []error) {
	var expr ast.Expr
	var errs []error
	errs = nil
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
		var expr1 ast.Expr
		var expr2 ast.Expr
		expr1 = p.expressionWrapper(&errs) // TODO better error handling
		// case 1 (expression)
		tok, eState := p.consume(t.RIGHT_PAREN, "Expect ')' after expression")
		switch tok.Types {
		case t.COMMA: // in case of 234, 123 it will just execute 234, not handling right expression
			p.advance() // skip current to next expression
			expr2 := p.expressionWrapper(&errs)
			p.advance()
			return expr2, errs
		default:
			if eState.HadError {
				errs = append(errs, New(tok, eState.S))
				return expr2, errs //empty ast.Expr, errs
			}
			//fmt.Println(ast.NewASTPrinter().Print(expr1))
			return expr1, errs // (expression), errs
		}
	} else {
		errs = append(errs, New(p.peek(), "expect expression"))
	}
	return expr, errs
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
	return p.peek(), eState
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().Types == t.SEMICOLON {
			return
		}
		switch p.peek().Types {
		case t.CLASS:
		case t.FUN:
		case t.VAR:
		case t.FOR:
		case t.IF:
		case t.WHILE:
		case t.PRINT:
		case t.RETURN:
			return
		}
	}
}

// Error is struct method that generate new parserError
func (e *parserError) Error() string {
	return e.s
}

// New set local eState
// print the error token
// return parserError msg
func New(tok t.Token, msg string) error {
	var eState g.ErrState //TODO

	eState.HadError = true
	// just report error back to user
	eState.ErnoToken(tok, msg)

	// then keep tracking what wrong
	return &parserError{msg} //TODO
}
