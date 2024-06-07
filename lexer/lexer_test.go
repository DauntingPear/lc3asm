package lexer

import (
	"testing"

	"lc3asm-parser/token"
)

func TestNextToken(t *testing.T) {
	input := `ADD R5,R5,R5;
.END
#44
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ADD, "ADD"},

		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.SEMICOLON, ";"},

		{token.PERIOD, "."},
		{token.END, "END"},

		{token.HASH, "#"},
		{token.INT, "44"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
