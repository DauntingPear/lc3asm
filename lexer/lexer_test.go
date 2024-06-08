package lexer

import (
	"testing"

	"lc3asm-parser/token"
)

func TestNextToken(t *testing.T) {
	input := `ADD R5,R5,R5;
.END
#44
x44
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

		{token.HEX, "x44"},
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

func TestTabIndentation(t *testing.T) {
	input := `LABEL:
	R5,R5,R5;
	ADD R5,R5,#1
NOT R1,R1
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "LABEL"},
		{token.COLON, ":"},

		{token.INDENT, "INDENT"},

		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.SEMICOLON, ";"},

		{token.ADD, "ADD"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		{token.DEDENT, "DEDENT"},

		{token.NOT, "NOT"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
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
