package lox

import (
	"strconv"
)

type Scanner struct {
	source  string
	srcRune []rune
	tokens  []Token
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
func (s *Scanner) scanTokens() errState {
	var eState errState
	eState.hadError = false
	for !s.isAtEnd() {
		s.start = s.current
		eState = s.scanToken()
	}
	s.addToken(EOF)
	return eState
}

func (s *Scanner) scanToken() errState {
	var eState errState
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
		break
	case ')':
		s.addToken(RIGHT_PAREN)
		break
	case '{':
		s.addToken(LEFT_BRACE)
		break
	case '}':
		s.addToken(RIGHT_PAREN)
		break
	case ',':
		s.addToken(COMMA)
		break
	case '.':
		s.addToken(DOT)
		break
	case '-':
		s.addToken(MINUS)
		break
	case '+':
		s.addToken(PLUS)
		break
	case ';':
		s.addToken(SEMICOLON)
		break
	case '*':
		s.addToken(STAR)
		break
	case '!':
		s.addToken(Iff(s.match('!'), BANG_EQUAL, BANG))
		break
	case '=':
		s.addToken(Iff(s.match('='), EQUAL_EQUAL, EQUAL))
		break
	case '<':
		s.addToken(Iff(s.match('='), LESS_EQUAL, LESS))
		break
	case '>':
		s.addToken(Iff(s.match('='), GREATER_EQUAL, GREATER))
		break
	case '/':
		if err := s.comment(); err != nil {
			eState.erno(s.line, err.Error())
		}
		break
	case ' ': // here just discard the white space
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
		break
	case '"': // string
		if err := s.literalString(); err != nil {
			eState.erno(s.line, err.Error())
		}
		break
	case '\'': //char
		if err := s.literalChar(); err != nil {
			eState.erno(s.line, err.Error())
		}
		break

	// C for IDENTIFIER(variable and the like)
	// regex form as follow
	// [a-zA-Z_][a-zA-Z_0-9]* a
	// if first c is numeric than it's literal
	// if first c is alpha than check regex,keyword , or IDENTIFIER/variable
	default:
		if s.isDigi(c) {
			if err := s.number(); err != nil {
				eState.erno(s.line, err.Error())
			}
			break
		} else if s.isAlpha(c) {
			if err := s.identifier(); err != nil {
				eState.erno(s.line, err.Error())
			}
			break
		}
		// unidentify
		eState.erno(s.line, "Unexpected character: ")
		break
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
			return New("comment token was not close")
		}
		s.advance()
		if !s.match('/') {
			return New("comment token was not close found *")
		}
	} else {
		s.addToken(SLASH)
	}
	return nil
}

func (s *Scanner) literalChar() error {
	c := s.peek()
	if c >= 0 && c <= 127 {
		if s.peekNext() != '\'' {
			return New("char was not properly close found: " + string(s.peekNext()))
		}
		s.advance()
		s.advance()
		s.addToken(CHAR, c)
		return nil
	}
	return New("not valid ascii char")
}

// literalString is function read multiline string
func (s *Scanner) literalString() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.peek() != '"' || s.isAtEnd() { //not found pair /* */
		return New("string was was not close")
	}
	s.advance() // current is "
	val := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, val)
	return nil
}

// number scan threw enter number
// need detect false number TODO
// need detect method inject TODO
func (s *Scanner) number() error {
	var isFloat bool
	for {
		if s.peek() >= '0' && s.peek() <= '9' {
			s.advance()
			continue
		}
		if s.peek() == '.' {
			if !s.isDigi(s.peekNext()) {
				// TODO probally 123.method() is allowed need to fix this line
				s.advance() // advance current . rune
				return New("number was not properly form last digit is .")
			}
			isFloat = true
			s.advance() // advance current . rune
			continue
		}
		break
	}
	if isFloat {
		val, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
		if err != nil {
			return New("INTERNAL float convert fail")
		}
		s.addToken(FLOAT, val)
		return nil
	}
	i64, err := strconv.ParseInt(s.source[s.start:s.current], 10, 32)
	if err != nil {
		return New("INTERNAL int convert fail")
	}
	s.addToken(INT, int32(i64))
	return nil
}

// identifier return it's keyword or user identifier
// TODO check maybe ? potential error
func (s *Scanner) identifier() error {
	for s.isAlpha(s.peek()) || s.isDigi(s.peek()) {
		s.advance()
	}
	val, prs := keywords[s.source[s.start:s.current]]
	if !prs {
		s.addToken(IDENTIFIER)
		return nil
	}
	s.addToken(val)
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
func (s *Scanner) addToken(typ TokenType, literals ...any) {
	if len(literals) >= 1 {
		// char literal, string literal , int literal, float literal
		s.addTokenS(typ, literals[0])
		return
	}
	// I just put nil better idea?
	s.addTokenS(typ, nil)
}
func (s *Scanner) addTokenS(typ TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		typ,
		text,
		literal,
		s.line,
	})
}
