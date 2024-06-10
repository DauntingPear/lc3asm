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
		{token.OPCODE, "ADD"},

		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.SEMICOLON, ";"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "END"},

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

		{token.OPCODE, "ADD"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		{token.DEDENT, "DEDENT"},

		{token.OPCODE, "NOT"},
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

func TestSpaceIndentation(t *testing.T) {
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

		{token.OPCODE, "ADD"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		{token.DEDENT, "DEDENT"},

		{token.OPCODE, "NOT"},
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

func TestBranchLexeme(t *testing.T) {
	input := `BRnzp
BRn
BRp
BRz
BRnp
BR
BRfsafda
BRnzp LABEL`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPCODE, "BRnzp"},
		{token.OPCODE, "BRn"},
		{token.OPCODE, "BRp"},
		{token.OPCODE, "BRz"},
		{token.OPCODE, "BRnp"},
		{token.OPCODE, "BR"},
		{token.IDENT, "BRfsafda"},
		{token.OPCODE, "BRnzp"},
		{token.IDENT, "LABEL"},
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

func TestLexemes(t *testing.T) {
	input := `ADD R5,R5,R5
ADD R1,R2,#1
NOT R1,R2
AND R1,R0,R2
AND R2,R2,#4
LD R1,LABEL
LDI R2,LABEL;
LDR R4,R1,#4
LEA R6,LABEL

ST R3,LABEL
STI R3,LABEL
STR R2,R1,LABEL

BRnzp LABEL

JMP R1
JSR LABEL
JSRR R4
RET
RTI

LABEL:
	ADD R5,R5,#1
ADD R1,R1,R1

.END
.BEGIN
.STRINGZ
.FILL
.BLKW

TRAP x22
GETC
OUT
PUTS
IN
PUTSP
HALT`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// ADD R5,R5,R5
		// 0-5
		{token.OPCODE, "ADD"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},

		// ADD R1,R2,#1
		// 6-12
		{token.OPCODE, "ADD"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		// NOT R1,R2
		// 13-16
		{token.OPCODE, "NOT"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},

		// AND R1,R0,R2
		// 17-22
		{token.OPCODE, "AND"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R0"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},

		//AND R2,R2,#4
		// 23-29
		{token.OPCODE, "AND"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "4"},

		// LD R1,LABEL
		// 30-33
		{token.OPCODE, "LD"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		// LDI R2,LABEL;
		// 34-38
		{token.OPCODE, "LDI"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},
		{token.SEMICOLON, ";"},

		// LDR R4,R1,#4
		// 39-45
		{token.OPCODE, "LDR"},
		{token.REGISTER, "R4"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "4"},

		// LEA R6,LABEL
		// 46-49
		{token.OPCODE, "LEA"},
		{token.REGISTER, "R6"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		// ST R3,LABEL
		// 50-53
		{token.OPCODE, "ST"},
		{token.REGISTER, "R3"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		// STI R3,LABEL
		// 54-57
		{token.OPCODE, "STI"},
		{token.REGISTER, "R3"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		// STR R2,R1,LABEL
		// 58-63
		{token.OPCODE, "STR"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		// BRnzp LABEL
		// 64-65
		{token.OPCODE, "BRnzp"},
		{token.IDENT, "LABEL"},

		// JMP R1
		// 66-67
		{token.OPCODE, "JMP"},
		{token.REGISTER, "R1"},

		// JSR LABEL
		// 68-69
		{token.OPCODE, "JSR"},
		{token.IDENT, "LABEL"},

		// JSRR R4
		// 70-71
		{token.OPCODE, "JSRR"},
		{token.REGISTER, "R4"},

		// RET
		// 72
		{token.OPCODE, "RET"},

		// RTI
		// 73
		{token.OPCODE, "RTI"},

		// LABEL:
		// 74-75
		{token.IDENT, "LABEL"},
		{token.COLON, ":"},

		// INDENT
		// 76
		{token.INDENT, "INDENT"},

		// ADD R5,R5,#1
		// 77-83
		{token.OPCODE, "ADD"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		// DEDENT
		// 84
		{token.DEDENT, "DEDENT"},

		// ADD R1,R1,R1
		// 85-90
		{token.OPCODE, "ADD"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},

		// .END
		// 91-92
		{token.PERIOD, "."},
		{token.DIRECTIVE, "END"},

		// .BEGIN
		// 93-94
		{token.PERIOD, "."},
		{token.DIRECTIVE, "BEGIN"},

		// .STRINGZ
		// 95-96
		{token.PERIOD, "."},
		{token.DIRECTIVE, "STRINGZ"},

		// .FILL
		// 97-98
		{token.PERIOD, "."},
		{token.DIRECTIVE, "FILL"},

		// .BLKW
		// 99-100
		{token.PERIOD, "."},
		{token.DIRECTIVE, "BLKW"},

		// TRAP x22
		// 101-102
		{token.TRAP, "TRAP"},
		{token.HEX, "x22"},

		// GETC
		// 103
		{token.TRAP, "GETC"},

		// OUT
		// 104
		{token.TRAP, "OUT"},

		// PUTS
		// 105
		{token.TRAP, "PUTS"},

		// IN
		// 106
		{token.TRAP, "IN"},

		// PUTSP
		// 107
		{token.TRAP, "PUTSP"},

		// HALT
		// 108
		{token.TRAP, "HALT"},
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
