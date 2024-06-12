package ast

import "lc3asm-parser/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ADD, AND
type ThreeRegisterStatement struct {
	Token           token.Token // The opcode
	DataRegister    *Register
	SourceRegisters [2]*Register
}

func (trs *ThreeRegisterStatement) statementNode()       {}
func (trs *ThreeRegisterStatement) TokenLiteral() string { return trs.Token.Literal }

// ADD_i, AND_i
type TwoRegisterImmediate struct {
	Token          token.Token // The opcode
	DataRegister   *Register
	SourceRegister *Register
	Immediate      int
}

func (tri *TwoRegisterImmediate) statementNode()       {}
func (tri *TwoRegisterImmediate) TokenLiteral() string { return tri.Token.Literal }

// LD, LDI, LEA, ST, STI
type RegisterLabelStatement struct {
	Token    token.Token // The opcode
	Register *Register
	Label    *Label
}

func (rls *RegisterLabelStatement) statementNode()       {}
func (rls *RegisterLabelStatement) TokenLiteral() string { return rls.Token.Literal }

// LDR, STR
type TwoRegisterOffsetStatement struct {
	Token         token.Token // The opcode
	LeftRegister  *Register
	RightRegister *Register
	Offset        int
}

func (tro *TwoRegisterOffsetStatement) statementNode()       {}
func (tro *TwoRegisterOffsetStatement) TokenLiteral() string { return tro.Token.Literal }

// NOT
type TwoRegisterStatement struct {
	Token          token.Token // The opcode
	DataRegister   *Register
	SourceRegister *Register
}

func (tr *TwoRegisterStatement) statementNode()       {}
func (tr *TwoRegisterStatement) TokenLiteral() string { return tr.Token.Literal }

type Opcode struct {
	Token   token.Token
	Literal string
}

func (o *Opcode) statementNode()       {}
func (o *Opcode) TokenLiteral() string { return o.Token.Literal }

type Register struct {
	Token token.Token
	Value int
	ID    int
}

func (r *Register) statementNode()       {}
func (r *Register) TokenLiteral() string { return r.Token.Literal }

type Label struct {
	Token token.Token
	Value string
}

func (l *Label) statementNode()       {}
func (l *Label) TokenLiteral() string { return l.Token.Literal }

// BR
type BranchStatement struct {
	Token token.Token
	N     bool
	Z     bool
	P     bool
	Label *Label
}

func (b *BranchStatement) statementNode()       {}
func (b *BranchStatement) TokenLiteral() string { return b.Token.Literal }

// JSR
type SingleLabelStatement struct {
	Token token.Token
	Label *Label
}

func (sls *SingleLabelStatement) statementNode()       {}
func (sls *SingleLabelStatement) TokenLiteral() string { return sls.Token.Literal }

type SingleRegisterStatement struct {
	Token    token.Token
	Register *Register
}

func (srs *SingleRegisterStatement) statementNode()       {}
func (srs *SingleRegisterStatement) TokenLiteral() string { return srs.Token.Literal }

type NoArgStatement struct {
	Token token.Token
}

func (nas *NoArgStatement) statementNode()       {}
func (nas *NoArgStatement) TokenLiteral() string { return nas.Token.Literal }

type IntegerDirectiveStatement struct {
	Token token.Token
	Value int
}

func (ids *IntegerDirectiveStatement) statementNode()       {}
func (ids *IntegerDirectiveStatement) TokenLiteral() string { return ids.Token.Literal }

type HexDirectiveStatement struct {
	Token token.Token
	Value int
}

func (hds *HexDirectiveStatement) statementNode()       {}
func (hds *HexDirectiveStatement) TokenLiteral() string { return hds.Token.Literal }

type NoArgDirective struct {
	Token token.Token
}

func (nad *NoArgDirective) statementNode()       {}
func (nad *NoArgDirective) TokenLiteral() string { return nad.Token.Literal }

type StringDirectiveStatement struct {
	Token token.Token
	Value string
}

func (sds *StringDirectiveStatement) statementNode()       {}
func (sds *StringDirectiveStatement) TokenLiteral() string { return sds.Token.Literal }

type HexTrapStatement struct {
	Token token.Token
	Value int
}

func (ht *HexTrapStatement) statementNode()       {}
func (ht *HexTrapStatement) TokenLiteral() string { return ht.Token.Literal }

type NoArgTrapStatement struct {
	Token token.Token
}

func (nat *NoArgTrapStatement) statementNode()       {}
func (nat *NoArgTrapStatement) TokenLiteral() string { return nat.Token.Literal }
