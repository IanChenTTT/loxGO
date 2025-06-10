package main

import (
	"strconv"
)

type Literal interface {
	string | float32 | int
}
type Token struct {
	types   TokenType
	lexemes string
	literal any
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
	switch t.literal.(type) {
	case string:
		str, _ := t.literal.(string)
		return t.types.String() + " " + t.lexemes + str
	case rune: // single char
		in, _ := t.literal.(rune)
		return t.types.String() + " " + t.lexemes + string(in)
	case int: //int 32 default golang int type
		in, _ := t.literal.(int)
		s := strconv.Itoa(in)
		return t.types.String() + " " + t.lexemes + s
	case float64:
		in, _ := t.literal.(float64)
		s := strconv.FormatFloat(in, 'E', -1, 32)
		return t.types.String() + " " + t.lexemes + s
	case nil:
		return t.types.String() + " " + t.lexemes + NIL.String()
	default:
		return "token not found"
	}
}
