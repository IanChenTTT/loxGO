package scanner

import (
	g "github.com/IanChenTTT/loxGO/internal/lox/global"
	t "github.com/IanChenTTT/loxGO/internal/lox/token"

	"strconv"
)

type Scanner struct {
	source  string
	srcRune []rune
	Tokens  []t.Token
	start   int
	current int
	line    int
}

func (s *Scanner) Scanner(source string) {
	s.source = source
	s.srcRune = []rune(s.source)
}

// scanTokens determine leximine
// convert leximine to token
func (s *Scanner) ScanTokens() g.ErrState {
	var eState g.ErrState
	eState.HadError = false
	for !s.isAtEnd() {
		s.start = s.current
		eState = s.scanToken()
	}
	s.start = s.current // Make sure last Lexemes is at end
	s.addToken(t.EOF)
	return eState
}

func (s *Scanner) scanToken() g.ErrState {
	var eState g.ErrState
	c := s.advance()
	switch c {
	case '(':
		s.addToken(t.LEFT_PAREN)
	case ')':
		s.addToken(t.RIGHT_PAREN)
	case '{':
		s.addToken(t.LEFT_BRACE)
	case '}':
		s.addToken(t.RIGHT_PAREN)
	case ',':
		s.addToken(t.COMMA)
	case '.':
		s.addToken(t.DOT)
	case '-':
		s.addToken(t.MINUS)
	case '+':
		s.addToken(t.PLUS)
	case ';':
		s.addToken(t.SEMICOLON)
	case '*':
		s.addToken(t.STAR)
	case '!':
		s.addToken(g.Iff(s.match('!'), t.BANG_EQUAL, t.BANG))
	case '=':
		s.addToken(g.Iff(s.match('='), t.EQUAL_EQUAL, t.EQUAL))
	case '<':
		s.addToken(g.Iff(s.match('='), t.LESS_EQUAL, t.LESS))
	case '>':
		s.addToken(g.Iff(s.match('='), t.GREATER_EQUAL, t.GREATER))
	case '?':
		s.addToken(t.CONDITION)
	case ':':
		s.addToken(t.COLON)
	case '/':
		if err := s.comment(); err != nil {
			eState.Erno(s.line, err.Error())
		}
	case ' ': // here just discard the white space
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
	case '"': // string
		if err := s.literalString(); err != nil {
			eState.Erno(s.line, err.Error())
		}
	case '\'': //char
		if err := s.literalChar(); err != nil {
			eState.Erno(s.line, err.Error())
		}

	// C for IDENTIFIER(variable and the like)
	// regex form as follow
	// [a-zA-Z_][a-zA-Z_0-9]* a
	// if first c is numeric than it's literal
	// if first c is alpha than check regex,keyword , or IDENTIFIER/variable
	default:
		if s.isDigi(c) {
			if err := s.number(); err != nil {
				eState.Erno(s.line, err.Error())
			}
			break
		} else if s.isAlpha(c) {
			if err := s.identifier(); err != nil {
				eState.Erno(s.line, err.Error())
			}
			break
		}
		// unidentify
		eState.Erno(s.line, "Unexpected character: ")
	}
	return eState
}

//
// MAIN
//

// comment make single line comment
// and multiline comment
func (s *Scanner) comment() error {
	// single line comment goes until end of line -> //
	// multy line comment goest until end of match -> /**/
	if s.match('/') {
		for s.peek() != '\n' && !s.isAtEnd() {
			s.advance()
		}
	} else if s.match('*') {
		for s.peek() != '*' && !s.isAtEnd() {
			if s.peek() == '\n' {
				s.line++
			}
			s.advance()
		}
		if s.peek() != '*' || s.isAtEnd() { //not found pair /* */
			return g.New("comment token was not close")
		}
		s.advance()
		if !s.match('/') {
			return g.New("comment token was not close found *")
		}
	} else {
		s.addToken(t.SLASH)
	}
	return nil
}

func (s *Scanner) literalChar() error {
	c := s.peek()
	if c >= 0 && c <= 127 {
		if s.peekNext() != '\'' {
			// if any error occur skip current line like c
			// single line operation !!!!
			for s.peek() != '\n' {
				s.advance()
			}
			return g.New("char was not properly close found: " + s.source[s.start:s.current])
		}
		s.advance()
		s.advance()
		s.addToken(t.CHAR, c)
		return nil
	}
	return g.New("not valid ascii char")
}

// literalString is function read multiline string
func (s *Scanner) literalString() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	// not found pair ""
	// so it read till the end
	// multi line operation !!!!
	if s.peek() != '"' || s.isAtEnd() { //not found pair /* */
		return g.New("string was was not close: " + s.source[s.start:s.current])
	}
	s.advance() // current is "
	val := s.source[s.start+1 : s.current-1]
	s.addToken(t.STRING, val)
	return nil
}

// number scan threw enter number
// need detect false number TODO
// need detect method inject TODO
func (s *Scanner) number() error {
	var isFloat bool
	for {
		if s.peek() >= '0' && s.peek() <= '9' { // keep scanning if it's number
			s.advance()
			continue
		}
		if s.peek() == '.' { // it's a float 123.
			if !s.isDigi(s.peekNext()) {
				s.advance() // advance current . rune
				return g.New("number was not properly form last digit is .")
			}
			isFloat = true
			s.advance() // advance current . rune
			continue
		}
		break
	}
	// detect end of literal
	// TODO fix this shit
	// ex:
	// 123[a-zA-Z]+ => no
	// 123 +-/*)' ''EOF''\n' => yes
	if s.peek() != ' ' &&
		s.peek() != '+' &&
		s.peek() != '-' &&
		s.peek() != '*' &&
		s.peek() != '/' &&
		s.peek() != ')' &&
		s.peek() != ',' &&
		s.peek() != '=' &&
		s.peek() != '?' &&
		s.peek() != ':' &&
		s.peek() != '<' &&
		s.peek() != '>' &&
		s.peek() != '\n' &&
		s.peek() != 0 {
		e := g.New("not a number missing a seperator ? " + s.source[s.start:s.current] + string(s.peek()))
		s.advance() // advance error rune //TODO
		return e
	}
	if isFloat {
		val, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
		if err != nil {
			return g.New("INTERNAL float convert fail")
		}
		s.addToken(t.FLOAT, val)
		return nil
	}
	i64, err := strconv.ParseInt(s.source[s.start:s.current], 10, 32)
	if err != nil {
		return g.New("INTERNAL int convert fail")
	}
	s.addToken(t.INT, int(i64))
	return nil
}

// identifier return it's keyword or user identifier
// TODO check maybe ? potential error
func (s *Scanner) identifier() error {
	for s.isAlpha(s.peek()) || s.isDigi(s.peek()) {
		s.advance()
	}
	s.addToken(t.Lookup(s.source[s.start:s.current]))
	return nil
}
func (s *Scanner) isAlpha(r rune) bool {
	if (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r == '_') {
		return true
	}
	return false
}
func (s *Scanner) isDigi(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

//
// UTIL
//

// peek return current word
// if current = end return 0
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0 // 0 null character \0
	}
	return s.srcRune[s.current]
}

// peekNext return current +1 word
// if current+1 = end return 0
func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.srcRune[s.current+1]
}

// move current rune to next
// return old rune
func (s *Scanner) advance() rune {
	r := s.srcRune[s.current]
	s.current++
	return r
}

// match rune to current word
func (s *Scanner) match(r rune) bool {
	if s.isAtEnd() {
		return false
	}
	//because advance func execute once
	//before match func current already+1
	if s.srcRune[s.current] != r {
		return false
	}
	s.current++ // already look ahead one
	return true
}

// isAtEnd function return current greater then source
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// addToken is wrapper to add token ,
// default token literal type is IDENTIFIER
func (s *Scanner) addToken(typ t.TokenType, literals ...any) {
	if len(literals) >= 1 {
		// char literal, string literal , int literal, float literal
		s.addTokenS(typ, literals[0])
		return
	}
	// I just put nil better idea?
	s.addTokenS(typ, nil)
}
func (s *Scanner) addTokenS(typ t.TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.Tokens = append(s.Tokens, t.Token{
		Types:   typ,
		Lexemes: text,
		Literal: literal,
		Line:    s.line,
	})
}
