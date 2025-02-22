package ast

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
	Token           token.Token
	Opcode          *Opcode
	DataRegister    *Register
	SourceRegisters [2]*Register
}

func (trs *ThreeRegisterStatement) statementNode()       {}
func (trs *ThreeRegisterStatement) TokenLiteral() string { return trs.Token.Literal }

// ADD_i, AND_i
type TwoRegisterImmediate struct {
	Token          token.Token
	Opcode         *Opcode
	DataRegister   *Register
	SourceRegister *Register
	Immediate      int
}

func (tri *TwoRegisterImmediate) statementNode()       {}
func (tri *TwoRegisterImmediate) TokenLiteral() string { return tri.Token.Literal }

// LD, LDI, LEA, ST, STI
type RegisterLabelStatement struct {
	Token    token.Token
	Opcode   *Opcode
	Register *Register
	Label    *Label
}

// LDR, STR
func (rls *RegisterLabelStatement) statementNode()       {}
func (rls *RegisterLabelStatement) TokenLiteral() string { return rls.Token.Literal }

type TwoRegisterOffset struct {
	Token         token.Token
	Opcode        *Opcode
	LeftRegister  *Register
	RightRegister *Register
	Offset        int
}

func (tro *TwoRegisterOffset) statementNode()       {}
func (tro *TwoRegisterOffset) TokenLiteral() string { return tro.Token.Literal }

// NOT
type TwoRegister struct {
	Token          token.Token
	Opcode         *Opcode
	DataRegister   *Register
	SourceRegister *Register
}

func (tr *TwoRegister) statementNode()       {}
func (tr *TwoRegister) TokenLiteral() string { return tr.Token.Literal }

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
