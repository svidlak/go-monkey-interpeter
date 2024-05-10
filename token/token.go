package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identtifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	MULTIPLY  = "*"
	DIVISINON = "/"

	//delitmiters
	COMMA     = ","
	SEMICOLON = ";"
	DOT       = "."

	//parenthesis
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
