package token

import "fmt"

type TokenType int

const (
	// Sigle character tokens.
	LPAREN TokenType = iota
	RPAREN
	LBRACE
	RBRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	NOT
	NEQ
	ASSIGN
	EQUAL
	GT
	GEQ
	LT
	LEQ

	// Literals
	IDENT
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

// token description
var tokenNames = []string{
	// Sigle character tokens.
	"LPAREN",
	"RPAREN",
	"LBRACE",
	"RBRACE",
	"COMMA",
	"DOT",
	"MINUS",
	"PLUS",
	"SEMICOLON",
	"SLASH",
	"STAR",
	// One or two character tokens.
	"NOT",
	"NEQ",
	"ASSIGN",
	"EQUAL",
	"GT",
	"GEQ",
	"LT",
	"LEQ",

	// Literals
	"IDENT",
	"STRING",
	"NUMBER",

	// keywords
	"AND",
	"CLASS",
	"ELSE",
	"FALSE",
	"FUN",
	"FOR",
	"IF",
	"NIL",
	"OR",
	"PRINT",
	"RETURN",
	"SUPER",
	"THIS",
	"TRUE",
	"VAR",
	"WHILE",

	"EOF",
}

type Object interface{}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal Object
	Line    int
}

func New(typ TokenType, lex string, lit Object, lin int) *Token {
	return &Token{
		Type:    typ,
		Lexeme:  lex,
		Literal: lit,
		Line:    lin,
	}
}

func (t *Token) ToString() string {
	if t.Literal != nil {
		return fmt.Sprintf("<%s, '%s', '%v'>", tokenNames[t.Type], t.Lexeme, t.Literal)
	}
	return fmt.Sprintf("<%s, '%s'>", tokenNames[t.Type], t.Lexeme)
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
