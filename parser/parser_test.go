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
