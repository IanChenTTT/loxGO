package token

import (
	"strconv"
)

type Literal interface {
	string | float32 | int
}
type Token struct {
	Types   TokenType
	Lexemes string
	Literal any
	Line    int
}
type Tokenize interface {
	token(t Token) error
	toString() string
}

func (t *Token) token(parToken Token) error {
	*t = parToken
	return nil
}
func (t *Token) ToString() string {
	switch t.Literal.(type) {
	case string:
		str, _ := t.Literal.(string)
		return "type: " + t.Types.String() + " lexemes: " + t.Lexemes + " literal: " + str
	case int32: // single char rune is allias to int 32 wtf
		switch t.Types {
		case CHAR:
			in, _ := t.Literal.(rune)
			return "type: " + t.Types.String() + " lexemes: " + t.Lexemes + " literal: " + string(in)
		case INT:
			in, _ := t.Literal.(int32)
			s := strconv.FormatInt(int64(in), 10)
			return "type: " + t.Types.String() + " lexemes: " + t.Lexemes + " literal: " + s
		default:
			return "int token type doesn't match literal"
		}
	case float64:
		in, _ := t.Literal.(float64)
		s := strconv.FormatFloat(in, 'E', -1, 32)
		return "type: " + t.Types.String() + " lexemes: " + t.Lexemes + " literal: " + s
	case nil:
		return "type: " + t.Types.String() + " lexemes: " + t.Lexemes + " literal: " + NIL.String()
	default:
		return "token not found"
	}
}
