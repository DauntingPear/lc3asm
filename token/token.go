package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Symbols
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	PERIOD    = "."
	HASH      = "#"

	// Labels and registers
	IDENT    = "IDENT"    // Labels
	REGISTER = "REGISTER" // Registers R0-8

	// Opcodes
	ADD  = "ADD"
	NOT  = "NOT"
	AND  = "AND"
	LD   = "LD"
	LDI  = "LDI"
	LDR  = "LDR"
	LEA  = "LEA"
	ST   = "ST"
	STI  = "STI"
	STR  = "STR"
	BR   = "BR"
	JMP  = "JMP"
	JSR  = "JSR"
	JSRR = "JSRR"
	RET  = "RET"
	RTI  = "RTI"

	// Directives
	END     = "END"
	FILL    = "FILL"
	BLKW    = "BLKW"
	STRINGZ = "STRINGZ"
	ORIG    = "ORIG"

	// Traps
	GETC  = "GETC"
	OUT   = "OUT"
	PUTS  = "PUTS"
	IN    = "IN"
	PUTSP = "PUTSP"
	HALT  = "HALT"
	TRAP  = "TRAP"

	// Number Types
	INT = "INT"
	HEX = "HEX"

	// Indentation
	INDENT = "INDENT"
	DEDENT = "DEDENT"
)

var keywords = map[string]TokenType{
	// Opcodes
	"ADD":  ADD,
	"NOT":  NOT,
	"AND":  AND,
	"LD":   LD,
	"LDI":  LDI,
	"LDR":  LDR,
	"LEA":  LEA,
	"ST":   ST,
	"STI":  STI,
	"STR":  STR,
	"BR":   BR,
	"JMP":  JMP,
	"JSR":  JSR,
	"JSRR": JSRR,
	"RET":  RET,
	"RTI":  RTI,

	// Directives
	"END":     END,
	"ORIG":    ORIG,
	"FILL":    FILL,
	"BLKW":    BLKW,
	"STRINGZ": STRINGZ,

	// TRAPS
	"TRAP":  TRAP,
	"GETC":  GETC,
	"OUT":   OUT,
	"PUTS":  PUTS,
	"IN":    IN,
	"PUTSP": PUTSP,
	"HALT":  HALT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
