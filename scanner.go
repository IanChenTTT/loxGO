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
func (s *Scanner) scanTokens() ([]Token, errState) {
	var t []Token
	var eState errState
	for s.isAtEnd() {
		s.start = s.current
		eState = s.scanToken()
	}
	t = append(t, Token{
		EOF,
		"",
		"", //TODO fixed in token.go
		s.line,
	})
	return t, eState
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
	default:
		eState.erno(s.line, "Unexpected character")
		break
	}
	return eState
}

// match rune to current word
func (s *Scanner) match(r rune) bool {
	if s.isAtEnd() {
		return false
	}
	//because advance func execute once
	//before match func current already+1
	if s.srcRune[s.current] == r {
		return true
	}
	s.current++ // already look ahead one
	return false
}

// move current rune to next
// return old rune
func (s *Scanner) advance() rune {
	r := s.srcRune[s.current]
	s.current++
	return r
}
func (s *Scanner) addToken(typ TokenType) {
	s.addTokenS(typ, "")
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
