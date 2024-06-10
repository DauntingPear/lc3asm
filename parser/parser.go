package parser

import (
	"fmt"
	"lc3asm-parser/ast"
	"lc3asm-parser/lexer"
	"lc3asm-parser/token"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.OPCODE:
		stmt := p.parseOpcodeSatement()
		return stmt
	default:
		return nil
	}
}

func (p *Parser) parseOpcodeSatement() ast.Statement {
	switch p.curToken.Literal {
	case "ADD", "AND":
		stmt := p.parseOperationOpcodeStatement()
		return stmt
	case "NOT":
		stmt := p.parseTwoRegisterStatement()
		return stmt
	case "ST", "STI", "LD", "LDI", "LEA":
		stmt := p.parseRegisterLabelStatement()
		return stmt
	case "STR", "LDR":
		stmt := p.parseTwoRegisterOffsetStatement()
		return stmt
	default:
		return nil
	}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

		return nil
	}
}

func (p *Parser) parseRegisterLabelStatement() ast.Statement {
	opcode := p.curToken

	if !p.expectPeek(token.REGISTER) {
		return nil
	}

	registerID := string(p.curToken.Literal[1])
	num, err := strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	register := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		fmt.Println("ERR: expected comma after register")
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		fmt.Println("ERR: expected label at label")
		return nil
	}

	label := &ast.Label{Token: p.curToken, Value: p.curToken.Literal}

	stmt := &ast.RegisterLabelStatement{
		Token:    opcode,
		Register: register,
		Label:    label,
	}

	return stmt

}

func (p *Parser) parseTwoRegisterStatement() ast.Statement {
	opcode := p.curToken

	if !p.expectPeek(token.REGISTER) {
		return nil
	}

	registerID := string(p.curToken.Literal[1])
	num, err := strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	dataRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		fmt.Println("ERR: expected comma after data register")
		return nil
	}

	if !p.expectPeek(token.REGISTER) {
		fmt.Println("ERR: expected register at source register 1")
		return nil
	}

	registerID = string(p.curToken.Literal[1])
	num, err = strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	sourceRegister := &ast.Register{Token: p.curToken, ID: num}

	stmt := &ast.TwoRegisterStatement{Token: opcode,
		DataRegister:   dataRegister,
		SourceRegister: sourceRegister,
	}

	return stmt
}

func (p *Parser) parseOperationOpcodeStatement() ast.Statement {
	opcodeToken := p.curToken

	if !p.expectPeek(token.REGISTER) {
		fmt.Println("ERR: expected register at data register")
		return nil
	}

	registerID := string(p.curToken.Literal[1])
	num, err := strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	dataRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		fmt.Println("ERR: expected comma after data register")
		return nil
	}

	if !p.expectPeek(token.REGISTER) {
		fmt.Println("ERR: expected register at source register 1")
		return nil
	}

	registerID = string(p.curToken.Literal[1])
	num, err = strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	sourceRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		fmt.Println("ERR: expected comma after source register 1")
		return nil
	}

	if p.peekTokenIs(token.HASH) {
		p.nextToken() // consume above checked token
		if !p.expectPeek(token.INT) {
			fmt.Println("ERR: expected int after #")
			return nil
		}
		registerID = string(p.curToken.Literal[1])
		fmt.Println(p.curToken.Literal)
		num, err = strconv.Atoi(registerID)
		if err != nil {
			return nil
		}
		// int logic
		stmt := &ast.TwoRegisterImmediate{
			Token:          opcodeToken,
			DataRegister:   dataRegister,
			SourceRegister: sourceRegister,
			Immediate:      num,
		}

		return stmt
	} else if p.peekTokenIs(token.REGISTER) {
		p.nextToken() // consume above checked token
		// 3 register logic

		registerID = string(p.curToken.Literal[1])
		fmt.Println(p.curToken.Literal)
		num, err = strconv.Atoi(registerID)
		if err != nil {
			return nil
		}
		sourceRegister2 := &ast.Register{Token: p.curToken, ID: num}

		stmt := &ast.ThreeRegisterStatement{
			Token:           opcodeToken,
			DataRegister:    dataRegister,
			SourceRegisters: [2]*ast.Register{sourceRegister, sourceRegister2},
		}

		return stmt
	} else {
		fmt.Printf("ERR: expected expected register or hash after source register 1, got=%s\n", p.curToken.Literal)
		return nil
	}
}
