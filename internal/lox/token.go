package lox

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
		return "type: " + t.types.String() + " lexemes: " + t.lexemes + " literal: " + str
	case int32: // single char rune is allias to int 32 wtf
		switch t.types {
		case CHAR:
			in, _ := t.literal.(rune)
			return "type: " + t.types.String() + " lexemes: " + t.lexemes + " literal: " + string(in)
		case INT:
			in, _ := t.literal.(int32)
			s := strconv.FormatInt(int64(in), 10)
			return "type: " + t.types.String() + " lexemes: " + t.lexemes + " literal: " + s
		default:
			return "int token type doesn't match literal"
		}
	case float64:
		in, _ := t.literal.(float64)
		s := strconv.FormatFloat(in, 'E', -1, 32)
		return "type: " + t.types.String() + " lexemes: " + t.lexemes + " literal: " + s
	case nil:
		return "type: " + t.types.String() + " lexemes: " + t.lexemes + " literal: " + NIL.String()
	default:
		return "token not found"
	}
}
