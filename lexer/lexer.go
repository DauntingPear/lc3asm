package lexer

import (
	"fmt"
	"lc3asm-parser/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	indentation  int
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '.':
		tok = newToken(token.PERIOD, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '#':
		tok = newToken(token.HASH, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok = l.parseLetter()
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Read indentation characters
func (l *Lexer) readIndentation() token.Token {
	l.readChar()
	indents := 0

	// Consumes tab or spaces for counting indentation levels
	for l.ch == '\t' || l.ch == ' ' {
		indents++
		l.readChar()
	}

	var tok token.Token

	// Checks if indentation has decreased
	if indents < l.indentation {
		fmt.Println("lt")
		l.indentation = l.indentation - indents
		tok.Type = token.DEDENT
		tok.Literal = "DEDENT"
		return tok
	}

	// Checks if indentation has incresed
	if indents > l.indentation {
		fmt.Println("gt")
		l.indentation = l.indentation + indents
		tok.Type = token.INDENT
		tok.Literal = "INDENT"
		return tok
	}

	// If no change in indentation, return empty token
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) peekChars(num int) byte {
	if l.readPosition+num >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition+num]
	}
}

func isRegister(r byte, num byte) bool {
	if r == 'R' && isDigit(num) {
		return true
	} else {
		return false
	}
}

func (l *Lexer) parseLetter() token.Token {
	var tok token.Token
	if isRegister(l.ch, l.peekChar()) {
		tok.Literal = l.readRegister()
		tok.Type = token.REGISTER
		return tok
	} else if isHex(l.ch, l.peekChar()) {
		tok.Literal = l.readHex()
		tok.Type = token.HEX
		return tok
	}

	tok.Literal = l.readIdentifier()
	tok.Type = token.LookupIdent(tok.Literal)

	return tok
}

func (l *Lexer) readRegister() string {
	position := l.position
	l.readChar()
	l.readChar()
	return l.input[position:l.position]
}

func (l *Lexer) readHex() string {
	position := l.position
	l.readChar()
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isHex(ch byte, num byte) bool {
	if ch == 'x' && isDigit(num) {
		return true
	}
	return false
}
