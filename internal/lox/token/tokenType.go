package token

type TokenType int

const (
	// Single-character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	// One or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	AND
	OR
	// Literals
	IDENTIFIER
	STRING
	CHAR
	INT
	FLOAT
	//TODO
	begin_const_identifier
	NIL
	TRUE
	FALSE
	end_const_identifier
	//Keywords src/go/token/token.go beg if for indexing
	keyword_beg
	CLASS
	ELSE
	FUN
	FOR
	IF
	PRINT
	RETURN
	SUPER
	THIS
	VAR
	WHILE
	keyword_end

	EOF
)

var TokenName = map[TokenType]string{
	LEFT_PAREN:  "(",
	RIGHT_PAREN: ")",
	LEFT_BRACE:  "{",
	RIGHT_BRACE: "}",
	COMMA:       ",",
	DOT:         ".",
	MINUS:       "-",
	PLUS:        "+",
	SEMICOLON:   ";",
	SLASH:       "/",
	STAR:        "*",

	// One or two character tokens
	BANG:          "!",
	BANG_EQUAL:    "!=",
	EQUAL:         "=",
	EQUAL_EQUAL:   "==",
	GREATER:       ">",
	GREATER_EQUAL: ">=",
	LESS:          "<",
	LESS_EQUAL:    "<=",

	// Literals
	IDENTIFIER: "identifier",
	STRING:     "string",
	CHAR:       "char",
	INT:        "int",
	FLOAT:      "float",

	// Keywords
	AND:    "and",
	CLASS:  "class",
	ELSE:   "else",
	FALSE:  "false",
	FUN:    "fun",
	FOR:    "for",
	IF:     "if",
	NIL:    "nil",
	OR:     "or",
	PRINT:  "print",
	RETURN: "return",
	SUPER:  "super",
	THIS:   "this",
	TRUE:   "true",
	VAR:    "var",
	WHILE:  "while",

	EOF: "eof",
}

// make it private global variable
// use LookUp to find match keyword
var keywords map[string]TokenType          //key:string val: int
var const_identifiers map[string]TokenType //key:string val: int
func init() {
	keywords = make(map[string]TokenType, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[TokenName[i]] = i
	}
	const_identifiers = make(map[string]TokenType, end_const_identifier-(begin_const_identifier+1))
	for i := begin_const_identifier + 1; i < end_const_identifier; i++ {
		const_identifiers[TokenName[i]] = i
	}
}

// Lookup functon search keyword
// return identifier if not found , else keyword
func Lookup(idet string) TokenType {
	if tok, ok := keywords[idet]; ok {
		return tok
	}
	if tok, ok := const_identifiers[idet]; ok {
		return tok
	}
	return IDENTIFIER
}

// String return token string
func (t TokenType) String() string {
	return TokenName[t]
}
