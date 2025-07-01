package token

import (
	"fmt"
	"strconv"
	"strings"
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
	var build strings.Builder
	fmt.Fprintf(&build, "type: %-10s  lexemes: %-10s", t.Types.String(), t.Lexemes)
	switch t.Literal.(type) {
	case string:
		str, _ := t.Literal.(string)
		fmt.Fprintf(&build, "literal: %-10s", str)
		return build.String()
	case int:
		in, _ := t.Literal.(int)
		str := strconv.FormatInt(int64(in), 10)
		fmt.Fprintf(&build, "literal: %-10s", str)
		return build.String()
	case int32: // single char rune is allias to int 32 wtf
		in, _ := t.Literal.(rune)
		fmt.Fprintf(&build, "literal: %-10s", string(in))
		return build.String()
	case float64:
		in, _ := t.Literal.(float64)
		str := strconv.FormatFloat(in, 'E', -1, 32)
		fmt.Fprintf(&build, "literal: %-10s", str)
		return build.String()
	case nil:
		fmt.Fprintf(&build, "literal: %-10s", NIL.String())
		return build.String()
	default:
		return "token not found"
	}
}
