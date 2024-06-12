package ast

import (
	"lc3asm-parser/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ThreeRegisterStatement{
				Token: token.Token{Type: token.OPCODE, Literal: "ADD"},
				DataRegister: &Register{
					Token: token.Token{Type: token.REGISTER, Literal: "R3"},
					Value: 3,
				},
				SourceRegisters: [2]*Register{
					&Register{
						Token: token.Token{Type: token.REGISTER, Literal: "R2"},
						Value: 2,
					},
					&Register{
						Token: token.Token{Type: token.REGISTER, Literal: "R6"},
						Value: 6,
					},
				},
			},
		},
	}

	if program.String() != "ADD R3,R2,R6" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
