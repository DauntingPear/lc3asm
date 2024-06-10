package parser

import (
	"lc3asm-parser/ast"
	"lc3asm-parser/lexer"
	"testing"
)

func TestThreeRegisterStatements(t *testing.T) {
	input := `
ADD R5,R5,R3
AND R3,R2,R4
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedOpcodeLiteral    string
		expectedRegisterLiterals [3]string
	}{
		{"ADD", [3]string{"R5", "R5", "R3"}},
		{"AND", [3]string{"R3", "R2", "R4"}},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testThreeRegisterStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedRegisterLiterals) {
			return
		}
	}
}

func testThreeRegisterStatement(t *testing.T, s ast.Statement, ol string, rls [3]string) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	trs, ok := s.(*ast.ThreeRegisterStatement)
	if !ok {
		t.Errorf("s not *ast.ThreeRegisterStatement. got=%T", s)
		return false
	}

	if trs.DataRegister.TokenLiteral() != rls[0] {
		t.Errorf("trs.DataRegister.TokenLiteral not '%s'. got=%s", rls[0], trs.DataRegister.TokenLiteral())
		return false
	}

	return true
}

func TestTwoRegisterStatements(t *testing.T) {
	input := `
NOT R2,R1
NOT R2,R6
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedOpcodeLiteral    string
		expectedRegisterLiterals [2]string
	}{
		{"NOT", [2]string{"R2", "R1"}},
		{"NOT", [2]string{"R2", "R6"}},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testTwoRegisterStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedRegisterLiterals) {
			return
		}
	}
}

func testTwoRegisterStatement(t *testing.T, s ast.Statement, ol string, rls [2]string) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	trs, ok := s.(*ast.TwoRegisterStatement)
	if !ok {
		t.Errorf("s not *ast.TwoRegisterStatement. got=%T", s)
		return false
	}

	if trs.DataRegister.TokenLiteral() != rls[0] {
		t.Errorf("trs.DataRegister.TokenLiteral not '%s'. got=%s", rls[0], trs.DataRegister.TokenLiteral())
		return false
	}

	return true
}

func TestRegisterLabelStatements(t *testing.T) {
	input := `
LD R5, LABEL
LDI R2, MYLABEL
LEA R7, SomeLabel
ST R0, aLabel
STI R2, LabelHere
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 5 {
		t.Fatalf("program.Statements does not contain 5 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedOpcodeLiteral   string
		expectedRegisterLiteral string
		expectedLabelLiteral    string
	}{
		{"LD", "R5", "LABEL"},
		{"LDI", "R2", "MYLABEL"},
		{"LEA", "R7", "SomeLabel"},
		{"ST", "R0", "aLabel"},
		{"STI", "R2", "LabelHere"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testRegisterLabelStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedRegisterLiteral, tt.expectedLabelLiteral) {
			return
		}
	}
}

func testRegisterLabelStatement(t *testing.T, s ast.Statement, ol string, rl string, ll string) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	rls, ok := s.(*ast.RegisterLabelStatement)
	if !ok {
		t.Errorf("s not *ast.ThreeRegisterStatement. got=%T", s)
		return false
	}

	if rls.Register.TokenLiteral() != rl {
		t.Errorf("rls.Register.TokenLiteral not '%s'. got=%s", rl, rls.Register.TokenLiteral())
		return false
	}

	if rls.Label.TokenLiteral() != ll {
		t.Errorf("rls.Label.TokenLiteral not '%s'. got=%s", ll, rls.Label.TokenLiteral())
	}

	return true
}

func TestTwoRegisterOffsetStatement(t *testing.T) {
	input := `
STR R1, R3, #21
LDR R2, R4, #10
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedOpcodeLiteral        string
		expectedLeftRegisterLiteral  string
		expectedRightRegisterLiteral string
		expectedIntegerLiteralValue  int
	}{
		{"STR", "R1", "R3", 21},
		{"LDR", "R2", "R4", 10},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testTwoRegisterOffsetStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedLeftRegisterLiteral, tt.expectedRightRegisterLiteral, tt.expectedIntegerLiteralValue) {
			return
		}
	}
}

func testTwoRegisterOffsetStatement(t *testing.T, s ast.Statement, ol string, ll string, rl string, il int) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	tros, ok := s.(*ast.TwoRegisterOffsetStatement)
	if !ok {
		t.Errorf("s not *ast.ThreeRegisterStatement. got=%T", s)
		return false
	}

	if tros.LeftRegister.TokenLiteral() != ll {
		t.Errorf("rls.LeftRegister.TokenLiteral not '%s'. got=%s", ll, tros.LeftRegister.TokenLiteral())
		return false
	}

	if tros.RightRegister.TokenLiteral() != rl {
		t.Errorf("rls.RightRegister.TokenLiteral not '%s'. got=%s", rl, tros.RightRegister.TokenLiteral())
		return false
	}

	if tros.Offset != il {
		t.Errorf("rls.Label.TokenLiteral not '%s'. got=%d", ll, tros.Offset)
	}

	return true
}
