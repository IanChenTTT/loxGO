package main

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
func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
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
	s.tokens = append(s.tokens, Token{
		EOF,
		"",
		IDENTIFIER.String(),
		s.line,
	})
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
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
		break
	case '"':
		if err := s.literalString(); err != nil {
			eState.erno(s.line, err.Error())
		}
		break
	default:
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

// literalString is function read multiline string
func (s *Scanner) literalString() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.peek() != '"' || s.isAtEnd() { //not found pair /* */
		return New("string was proper close")
	}
	s.advance() // current is "
	val := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, val)
	return nil
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
func (s *Scanner) addToken(typ TokenType, literals ...string) {
	if len(literals) >= 1 {
		s.addTokenS(typ, literals[0])
		return
	}
	s.addTokenS(typ, IDENTIFIER.String())
}
func (s *Scanner) addTokenS(typ TokenType, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, Token{
		typ,
		text,
		literal,
		s.line,
	})
}
