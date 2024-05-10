package lexer

import (
	"testing"

	"github.com/svidlak/go-monkey-interpeter/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},,;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVISINON, "/"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.DOT, "."},
	}
}
