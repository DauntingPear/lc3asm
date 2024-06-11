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

func TestTwoRegisterOffsetStatements(t *testing.T) {
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

func TestBranchStatements(t *testing.T) {
	input := `
BR LABEL
BRn MYLabel
BRnzp THELabel
BRp ALabel
BRpzn SomeLabel`

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
		expectedOpcodeLiteral string
		expectedNBitValue     bool
		expectedZBitValue     bool
		expectedPBitValue     bool
		expectedLabelLiteral  string
	}{
		{"BR", true, true, true, "LABEL"},
		{"BRn", true, false, false, "MYLabel"},
		{"BRnzp", true, true, true, "THELabel"},
		{"BRp", false, false, true, "ALabel"},
		{"BRpzn", true, true, true, "SomeLabel"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testBranchStatement(
			t,
			stmt,
			tt.expectedOpcodeLiteral,
			tt.expectedNBitValue,
			tt.expectedZBitValue,
			tt.expectedPBitValue,
			tt.expectedLabelLiteral,
		) {
			return
		}
	}
}

func testBranchStatement(
	t *testing.T,
	s ast.Statement,
	ol string,
	n bool,
	z bool,
	p bool,
	literal string,
) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	b, ok := s.(*ast.BranchStatement)
	if !ok {
		t.Errorf("s not *ast.BranchStatement. got=%T", s)
		return false
	}

	if b.N != n {
		t.Errorf("b.N not '%t'. got=%t", n, b.N)
		return false
	}

	if b.Z != z {
		t.Errorf("b.Z not '%t'. got=%t", z, b.Z)
		return false
	}

	if b.P != p {
		t.Errorf("b.P not '%t'. got=%t", p, b.P)
		return false
	}

	if b.Label == nil {
		t.Fatalf("b.Label is nil")
	}

	if b.Label.Value != literal {
		t.Errorf("b.Label.Value not '%s'. got=%s", b.Label.Value, literal)
		return false
	}

	return true
}

func TestSingleLabelStatements(t *testing.T) {
	input := `
JSR MyLabel
JSR MySubroutine
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
		expectedOpcodeLiteral string
		expectedLiteral       string
	}{
		{"JSR", "MyLabel"},
		{"JSR", "MySubroutine"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSingleLabelStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedLiteral) {
			return
		}
	}
}

func testSingleLabelStatement(t *testing.T, s ast.Statement, ol string, literal string) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	sls, ok := s.(*ast.SingleLabelStatement)
	if !ok {
		t.Errorf("s not *ast.BranchStatement. got=%T", s)
		return false
	}

	if sls.Label.Value != literal {
		t.Errorf("sls.Label.Value not '%s'. got=%s", literal, sls.Label.Value)
		return false
	}

	return true
}

func TestSingleRegisterStatements(t *testing.T) {
	input := `
JMP R4
JSRR R0
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
		expectedOpcodeLiteral   string
		expectedRegisterLiteral string
		expectedRegisterID      int
	}{
		{"JMP", "R4", 4},
		{"JSRR", "R0", 0},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSingleRegisterStatement(t, stmt, tt.expectedOpcodeLiteral, tt.expectedRegisterLiteral, tt.expectedRegisterID) {
			return
		}
	}
}

func testSingleRegisterStatement(t *testing.T, s ast.Statement, ol string, rl string, rid int) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	srs, ok := s.(*ast.SingleRegisterStatement)
	if !ok {
		t.Errorf("s not *ast.SingleRegisterStatement. got=%T", s)
		return false
	}

	if srs.Register.TokenLiteral() != rl {
		t.Errorf("srs.Register.TokenLiteral() not '%s'. got=%s", rl, srs.Register.TokenLiteral())
		return false
	}

	if srs.Register.ID != rid {
		t.Errorf("srs.Register.ID not '%d'. got=%d", rid, srs.Register.ID)
		return false
	}

	return true
}

func TestNoArgStatements(t *testing.T) {
	input := `
RET
RTI
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
		expectedOpcodeLiteral string
	}{
		{"RET"},
		{"RTI"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testNoArgStatement(t, stmt, tt.expectedOpcodeLiteral) {
			return
		}
	}
}

func testNoArgStatement(t *testing.T, s ast.Statement, ol string) bool {
	if s.TokenLiteral() != ol {
		t.Errorf("s.TokenLiteral not '%s'. got=%q", ol, s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.NoArgStatement)
	if !ok {
		t.Errorf("s not *ast.ThreeRegisterStatement. got=%T", s)
		return false
	}

	return true
}
