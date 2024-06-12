package ast

import (
	"bytes"
	"fmt"
	"lc3asm-parser/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ADD, AND
type ThreeRegisterStatement struct {
	Token           token.Token // The opcode
	DataRegister    *Register
	SourceRegisters [2]*Register
}

func (trs *ThreeRegisterStatement) statementNode()       {}
func (trs *ThreeRegisterStatement) TokenLiteral() string { return trs.Token.Literal }
func (trs *ThreeRegisterStatement) String() string {
	var out bytes.Buffer

	out.WriteString(trs.TokenLiteral() + " ")
	out.WriteString(trs.DataRegister.String() + ",")
	out.WriteString(trs.SourceRegisters[0].String() + ",")
	out.WriteString(trs.SourceRegisters[1].String())

	return out.String()
}

// ADD_i, AND_i
type TwoRegisterImmediate struct {
	Token          token.Token // The opcode
	DataRegister   *Register
	SourceRegister *Register
	Immediate      int
}

func (tri *TwoRegisterImmediate) statementNode()       {}
func (tri *TwoRegisterImmediate) TokenLiteral() string { return tri.Token.Literal }
func (tri *TwoRegisterImmediate) String() string {
	var out bytes.Buffer

	out.WriteString(tri.TokenLiteral() + " ")
	out.WriteString(tri.DataRegister.String() + ",")
	out.WriteString(tri.SourceRegister.String() + ",")
	out.WriteString("#" + fmt.Sprintf("%d", tri.Immediate))

	return out.String()
}

// LD, LDI, LEA, ST, STI
type RegisterLabelStatement struct {
	Token    token.Token // The opcode
	Register *Register
	Label    *Label
}

func (rls *RegisterLabelStatement) statementNode()       {}
func (rls *RegisterLabelStatement) TokenLiteral() string { return rls.Token.Literal }
func (rls *RegisterLabelStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rls.TokenLiteral() + " ")
	out.WriteString(rls.Register.String() + " ")
	out.WriteString(rls.Label.String())

	return out.String()
}

// LDR, STR
type TwoRegisterOffsetStatement struct {
	Token         token.Token // The opcode
	LeftRegister  *Register
	RightRegister *Register
	Offset        int
}

func (tro *TwoRegisterOffsetStatement) statementNode()       {}
func (tro *TwoRegisterOffsetStatement) TokenLiteral() string { return tro.Token.Literal }
func (tro *TwoRegisterOffsetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(tro.TokenLiteral() + " ")
	out.WriteString(tro.LeftRegister.String() + ",")
	out.WriteString(tro.RightRegister.String() + ",")
	out.WriteString("#" + fmt.Sprintf("%d", tro.Offset))

	return out.String()
}

// NOT
type TwoRegisterStatement struct {
	Token          token.Token // The opcode
	DataRegister   *Register
	SourceRegister *Register
}

func (tr *TwoRegisterStatement) statementNode()       {}
func (tr *TwoRegisterStatement) TokenLiteral() string { return tr.Token.Literal }
func (tr *TwoRegisterStatement) String() string {
	var out bytes.Buffer

	out.WriteString(tr.TokenLiteral() + " ")
	out.WriteString(tr.DataRegister.String() + ",")
	out.WriteString(tr.SourceRegister.String())

	return out.String()
}

type Opcode struct {
	Token   token.Token
	Literal string
}

func (o *Opcode) statementNode()       {}
func (o *Opcode) TokenLiteral() string { return o.Token.Literal }
func (o *Opcode) String() string {
	var out bytes.Buffer

	out.WriteString(o.TokenLiteral())

	return out.String()
}

type Register struct {
	Token token.Token
	Value int
	ID    int
}

func (r *Register) statementNode()       {}
func (r *Register) TokenLiteral() string { return r.Token.Literal }
func (r *Register) String() string {
	var out bytes.Buffer

	out.WriteString(r.TokenLiteral())

	return out.String()
}

type Label struct {
	Token token.Token
	Value string
}

func (l *Label) statementNode()       {}
func (l *Label) TokenLiteral() string { return l.Token.Literal }
func (l *Label) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral())

	return out.String()
}

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
func (b *BranchStatement) String() string {
	var out bytes.Buffer

	out.WriteString(b.TokenLiteral())

	return out.String()
}

// JSR
type SingleLabelStatement struct {
	Token token.Token
	Label *Label
}

func (sls *SingleLabelStatement) statementNode()       {}
func (sls *SingleLabelStatement) TokenLiteral() string { return sls.Token.Literal }
func (sls *SingleLabelStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sls.TokenLiteral() + " ")
	out.WriteString(sls.Label.String())

	return out.String()
}

type SingleRegisterStatement struct {
	Token    token.Token
	Register *Register
}

func (srs *SingleRegisterStatement) statementNode()       {}
func (srs *SingleRegisterStatement) TokenLiteral() string { return srs.Token.Literal }
func (srs *SingleRegisterStatement) String() string {
	var out bytes.Buffer

	out.WriteString(srs.TokenLiteral() + " ")
	out.WriteString(srs.Register.String())

	return out.String()
}

type NoArgStatement struct {
	Token token.Token
}

func (nas *NoArgStatement) statementNode()       {}
func (nas *NoArgStatement) TokenLiteral() string { return nas.Token.Literal }
func (nas *NoArgStatement) String() string {
	var out bytes.Buffer

	out.WriteString(nas.TokenLiteral())

	return out.String()
}

type IntegerDirectiveStatement struct {
	Token token.Token
	Value int
}

func (ids *IntegerDirectiveStatement) statementNode()       {}
func (ids *IntegerDirectiveStatement) TokenLiteral() string { return ids.Token.Literal }
func (ids *IntegerDirectiveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ids.TokenLiteral() + " ")
	out.WriteString("#" + fmt.Sprintf("%d", ids.Value))

	return out.String()
}

type HexDirectiveStatement struct {
	Token token.Token
	Value int
}

func (hds *HexDirectiveStatement) statementNode()       {}
func (hds *HexDirectiveStatement) TokenLiteral() string { return hds.Token.Literal }
func (hds *HexDirectiveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(hds.TokenLiteral() + " ")
	out.WriteString(hds.Token.Literal)

	return out.String()
}

type NoArgDirective struct {
	Token token.Token
}

func (nad *NoArgDirective) statementNode()       {}
func (nad *NoArgDirective) TokenLiteral() string { return nad.Token.Literal }
func (nad *NoArgDirective) String() string {
	var out bytes.Buffer

	out.WriteString(nad.TokenLiteral())

	return out.String()
}

type StringDirectiveStatement struct {
	Token token.Token
	Value string
}

func (sds *StringDirectiveStatement) statementNode()       {}
func (sds *StringDirectiveStatement) TokenLiteral() string { return sds.Token.Literal }
func (sds *StringDirectiveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(sds.TokenLiteral() + " ")
	out.WriteString(`"` + sds.Value + `"`)

	return out.String()
}

type HexTrapStatement struct {
	Token token.Token
	Value int
}

func (ht *HexTrapStatement) statementNode()       {}
func (ht *HexTrapStatement) TokenLiteral() string { return ht.Token.Literal }
func (ht *HexTrapStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ht.TokenLiteral() + " ")
	out.WriteString(ht.Token.Literal)

	return out.String()
}

type NoArgTrapStatement struct {
	Token token.Token
}

func (nat *NoArgTrapStatement) statementNode()       {}
func (nat *NoArgTrapStatement) TokenLiteral() string { return nat.Token.Literal }
func (nat *NoArgTrapStatement) String() string {
	var out bytes.Buffer

	out.WriteString(nat.TokenLiteral())

	return out.String()
}
