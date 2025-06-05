package main

type Token struct {
	types   TokenType
	lexemes string
	literal string //TODO fixed type
	line    int
}
type Tokenize interface {
	token(t Token) error
	toString() string
}

func (t *Token) token(parToken Token) error {
	*t = parToken
	return nil
}
func (t *Token) toString() string {
	return t.types.String() + " " + t.lexemes + t.literal
}
