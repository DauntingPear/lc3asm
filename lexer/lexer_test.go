package lexer

import (
	"fmt"
	"testing"

	"lc3asm-parser/token"
)

func TestNextToken(t *testing.T) {
	input := `ADD R5,R5,R5;
.END
#44
x44
ADD R3,R4,R5
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

		{token.OPCODE, "ADD"},
		{token.REGISTER, "R3"},
		{token.COMMA, ","},
		{token.REGISTER, "R4"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
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

func TestOperationOpcodes(t *testing.T) {
	input := `ADD R5,R5,R5
ADD R1,R2,#1
NOT R1,R2
AND R1,R0,R2
AND R2,R2,#4`

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

		{token.OPCODE, "ADD"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "1"},

		{token.OPCODE, "NOT"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},

		{token.OPCODE, "AND"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.REGISTER, "R0"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},

		{token.OPCODE, "AND"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "4"},
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

func TestLoadOpcodes(t *testing.T) {
	input := `LD R1,LABEL
LDI R2,LABEL;
LDR R4,R1,#4
LEA R6,LABEL`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPCODE, "LD"},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		{token.OPCODE, "LDI"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},
		{token.SEMICOLON, ";"},

		{token.OPCODE, "LDR"},
		{token.REGISTER, "R4"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
		{token.HASH, "#"},
		{token.INT, "4"},

		{token.OPCODE, "LEA"},
		{token.REGISTER, "R6"},
		{token.COMMA, ","},
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

func TestStoreOpcodes(t *testing.T) {
	input := `ST R3,LABEL
STI R3,LABEL
STR R2,R1,LABEL`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPCODE, "ST"},
		{token.REGISTER, "R3"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		{token.OPCODE, "STI"},
		{token.REGISTER, "R3"},
		{token.COMMA, ","},
		{token.IDENT, "LABEL"},

		{token.OPCODE, "STR"},
		{token.REGISTER, "R2"},
		{token.COMMA, ","},
		{token.REGISTER, "R1"},
		{token.COMMA, ","},
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

func TestMovementOpcodes(t *testing.T) {
	input := `JMP R1
JSR LABEL
JSRR R4
RET
RTI`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPCODE, "JMP"},
		{token.REGISTER, "R1"},

		{token.OPCODE, "JSR"},
		{token.IDENT, "LABEL"},

		{token.OPCODE, "JSRR"},
		{token.REGISTER, "R4"},

		{token.OPCODE, "RET"},

		{token.OPCODE, "RTI"},
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

func TestTrapCodes(t *testing.T) {
	input := `TRAP x22
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
		{token.TRAP, "TRAP"},
		{token.HEX, "x22"},

		{token.TRAP, "GETC"},
		{token.TRAP, "OUT"},
		{token.TRAP, "PUTS"},
		{token.TRAP, "IN"},
		{token.TRAP, "PUTSP"},
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

func TestDirectives(t *testing.T) {
	input := `.END
.BEGIN
.STRINGZ
.FILL
.BLKW`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PERIOD, "."},
		{token.DIRECTIVE, "END"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "BEGIN"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "STRINGZ"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "FILL"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "BLKW"},
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

func TestInts(t *testing.T) {
	input := `#-1 #-3 #2`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.HASH, "#"},
		{token.INT, "-1"},

		{token.HASH, "#"},
		{token.INT, "-3"},

		{token.HASH, "#"},
		{token.INT, "2"},
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

func TestHex(t *testing.T) {
	input := `.END
.BEGIN
.STRINGZ
.FILL
.BLKW`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PERIOD, "."},
		{token.DIRECTIVE, "END"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "BEGIN"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "STRINGZ"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "FILL"},

		{token.PERIOD, "."},
		{token.DIRECTIVE, "BLKW"},
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

func TestComments(t *testing.T) {
	input := `NOT R5,R5; This is a comment
; This is a comment
	;comment indented
ADD`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.OPCODE, "NOT"},
		{token.REGISTER, "R5"},
		{token.COMMA, ","},
		{token.REGISTER, "R5"},
		{token.COMMENT, "; This is a comment"},

		{token.COMMENT, "; This is a comment"},

		{token.INDENT, "INDENT"},
		{token.COMMENT, ";comment indented"},
		{token.DEDENT, "DEDENT"},

		{token.OPCODE, "ADD"},
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

func TestString(t *testing.T) {
	input := `.STRINGZ "this is a string"
"this is another string"`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PERIOD, "."},
		{token.DIRECTIVE, "STRINGZ"},
		{token.STRING, "this is a string"},

		{token.STRING, "this is another string"},
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

func Test(t *testing.T) {
	input := `
;;+------------------------------------------------------------------------------+
;;|                                                                              |
;;|  Author: Adrian Brady                                                        |
;;|  Date: 04/01/2024                                                            |
;;|  Purpose: Breakout Game Lab for Computer Engineering class spring 2024 MATC. |
;;|                                                                              |
;;+------------------------------------------------------------------------------+

;;+------------------------------+
;;|       Initialization         |
;;+------------------------------+
.orig x3000
START:
  ; Reset Game values
  JSR ResetGameSR

  ; Initial Game Setup and Frame Buffer Initialization
  ; Preconditions: None
  ; Postconditions: Video buffer filled with black pixels, R0/R1/R4/R6 preserved
  JSR InitFrameSR

  ; Initial boundary and Gameplay Objects Drawing
  ; Preconditions: None
  ; Postconditions: Draws game boundaries, initializes game environment,
  ; draws ball. Register returns not considered. R7 is return address.
  JSR InitializeGameSR

  ; Main Game loop
  ; Preconditions: None
  ; Postconditions: End of game, program has finished
  JSR GameLoopSR

HALT


;;+------------------------------+
;;|       Game Init Section      |
;;+------------------------------+

;----------------------------
;; Resets game constants
;----------------------------
RESET_RET .FILL 0
ResetGameSR:
  ST R7,RESET_RET
  AND R0,R0,#0
  ST R0,BRICK_COLOR
  ST R0,WALL_COLOR
  ADD R0,R0,#2
  ST R0,BRICK_COL
  LD R0,ZERO
  ADD R0,R0,#5
  ST R0,BALL_X
  ST R0,BALL_Y
  LD R0,ZERO
  ADD R0,R0,#1
  ST R0,BALL_Y_DIR
  ST R0,BALL_X_DIR
  ADD R0,R0,#2
  ST R0,Bricks_Remaining
  LD R0,ZERO
  ADD R0,R0,#8
  ST R0,PADDLE_POS
  LD R0,ZERO
  ST R0,LEFT_COLOR
  ST R0,RIGHT_COLOR
  LD R7,RESET_RET
  RET
`
l := New(input)

	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}

		fmt.Println(tok)

	}
	t.Error("Errored")
}
