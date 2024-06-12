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

	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be '%s', got '%s' instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
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
	case token.DIRECTIVE:
		stmt := p.parseDirectiveStatement()
		return stmt
	case token.TRAP:
		stmt := p.parseTrapStatement()
		return stmt
	default:
		return nil
	}
}

// TODO: Write this function
func (p *Parser) parseDirectiveStatement() ast.Statement {
	switch p.curToken.Literal {
	case "BLKW": // .BLKW ####
		stmt := p.parseIntDirective()
		return stmt
	case "ORIG", "FILL": // .ORIG x####, .FILL x#####
		stmt := p.parseHexDirective()
		return stmt
	case "END": // .END
		stmt := p.parseNoArgDirective()
		return stmt
	case "STRINGZ": // .STRINGZ "String here"
		stmt := p.parseStringDirective()
		return stmt
	default:
		return nil
	}

}

func (p *Parser) parseIntDirective() ast.Statement {
	directive := p.curToken

	if !p.expectPeek(token.HASH) {
		return nil
	}

	if !p.expectPeek(token.INT) {
		return nil
	}

	num, err := strconv.Atoi(p.curToken.Literal)
	if err != nil {
		return nil
	}

	stmt := &ast.IntegerDirectiveStatement{
		Token: directive,
		Value: num,
	}

	return stmt
}

func (p *Parser) parseHexDirective() ast.Statement {
	directive := p.curToken

	if !p.expectPeek(token.HEX) {
		return nil
	}

	num, err := strconv.ParseInt(p.curToken.Literal[1:], 16, 64)
	if err != nil {
		return nil
	}

	stmt := &ast.HexDirectiveStatement{
		Token: directive,
		Value: int(num),
	}

	return stmt
}

func (p *Parser) parseNoArgDirective() ast.Statement {
	directive := p.curToken

	stmt := &ast.NoArgDirective{
		Token: directive,
	}

	return stmt
}

func (p *Parser) parseStringDirective() ast.Statement {
	directive := p.curToken

	if !p.expectPeek(token.STRING) {
		return nil
	}

	string := p.curToken.Literal

	stmt := &ast.StringDirectiveStatement{
		Token: directive,
		Value: string,
	}

	return stmt
}

// TODO: Write this function
func (p *Parser) parseTrapStatement() ast.Statement {
	return nil
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
	case "JMP", "JSRR":
		stmt := p.parseSingleRegisterStatement()
		return stmt
	case "JSR", "BR",
		"BRn", "BRnz", "BRnzp", "BRnp", "BRnpz",
		"BRp", "BRpn", "BRpnz", "BRpz", "BRpzn",
		"BRz", "BRzp", "BRzpn", "BRzn", "BRznp":
		stmt := p.parseSingleLabelStatement()
		return stmt
	case "RET", "RTI":
		stmt := p.parseNoArgStatement()
		return stmt
	default:
		return nil
	}
}

// TODO: Write this function
func (p *Parser) parseNoArgStatement() ast.Statement {
	opcode := p.curToken

	stmt := &ast.NoArgStatement{Token: opcode}

	return stmt
}

// TODO: Write this function
func (p *Parser) parseSingleRegisterStatement() ast.Statement {
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

	stmt := &ast.SingleRegisterStatement{
		Token:    opcode,
		Register: register,
	}

	return stmt
}

// TODO: Write this function
func (p *Parser) parseSingleLabelStatement() ast.Statement {
	var stmt ast.Statement

	if isBR(p.curToken.Literal) {
		stmt = p.parseBRStatement()
	} else {
		stmt = p.parseJSRStatement()
	}

	return stmt
}

func (p *Parser) parseBRStatement() ast.Statement {
	opcode := p.curToken

	stmt := &ast.BranchStatement{Token: opcode}

	if len(p.curToken.Literal) == 2 {
		stmt.N = true
		stmt.Z = true
		stmt.P = true
	} else {
		branches := p.curToken.Literal[1:]
		for _, ch := range branches {
			switch ch {
			case 'n':
				stmt.N = true
			case 'z':
				stmt.Z = true
			case 'p':
				stmt.P = true
			}
		}
	}

	if !p.expectPeek(token.IDENT) {
		return stmt
	}

	label := &ast.Label{Token: p.curToken, Value: p.curToken.Literal}

	stmt.Label = label

	return stmt
}

func (p *Parser) parseJSRStatement() ast.Statement {
	opcode := p.curToken

	stmt := &ast.SingleLabelStatement{Token: opcode}

	if !p.expectPeek(token.IDENT) {
		return stmt
	}

	label := &ast.Label{Token: p.curToken, Value: p.curToken.Literal}

	stmt.Label = label

	return stmt
}

func isBR(literal string) bool {
	if len(literal) >= 2 {
		if literal[0] == 'B' && literal[1] == 'R' {
			return true
		}
	}

	return false
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
		p.peekError(t)
		return false
	}
}

func (p *Parser) parseTwoRegisterOffsetStatement() ast.Statement {
	opcode := p.curToken

	if !p.expectPeek(token.REGISTER) {
		return nil
	}

	registerID := string(p.curToken.Literal[1])
	num, err := strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	leftRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		return nil
	}

	if !p.expectPeek(token.REGISTER) {
		return nil
	}

	registerID = string(p.curToken.Literal[1])
	num, err = strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	rightRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		return nil
	}

	if !p.expectPeek(token.HASH) {
		return nil
	}

	if !p.expectPeek(token.INT) {
		return nil
	}

	offsetValue := string(p.curToken.Literal)
	offset, err := strconv.Atoi(offsetValue)
	if err != nil {
		return nil
	}

	stmt := &ast.TwoRegisterOffsetStatement{Token: opcode,
		LeftRegister:  leftRegister,
		RightRegister: rightRegister,
		Offset:        offset,
	}

	return stmt
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
		return nil
	}

	if !p.expectPeek(token.IDENT) {
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
		return nil
	}

	if !p.expectPeek(token.REGISTER) {
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
		return nil
	}

	registerID := string(p.curToken.Literal[1])
	num, err := strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	dataRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		return nil
	}

	if !p.expectPeek(token.REGISTER) {
		return nil
	}

	registerID = string(p.curToken.Literal[1])
	num, err = strconv.Atoi(registerID)
	if err != nil {
		return nil
	}

	sourceRegister := &ast.Register{Token: p.curToken, ID: num}

	if !p.expectPeek(token.COMMA) {
		return nil
	}

	if p.peekTokenIs(token.HASH) {
		p.nextToken() // consume above checked token
		if !p.expectPeek(token.INT) {
			return nil
		}
		registerID = string(p.curToken.Literal[1])
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
		p.peekError(p.curToken.Type)
		return nil
	}
}
