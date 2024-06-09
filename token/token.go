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
	MINUS     = "-"
	QUOTE     = `"`

	// Keywords
	IDENT     = "IDENT"    // Labels
	REGISTER  = "REGISTER" // Registers R0-8
	OPCODE    = "OPCODE"
	DIRECTIVE = "DIRECTIVE"
	TRAP      = "TRAP"
	COMMENT   = "COMMENT"

	// Number Types
	INT = "INT"
	HEX = "HEX"

	// Indentation
	INDENT = "INDENT"
	DEDENT = "DEDENT"
)

var keywords = map[string]TokenType{
	// Opcodes
	"ADD":  OPCODE,
	"NOT":  OPCODE,
	"AND":  OPCODE,
	"LD":   OPCODE,
	"LDI":  OPCODE,
	"LDR":  OPCODE,
	"LEA":  OPCODE,
	"ST":   OPCODE,
	"STI":  OPCODE,
	"STR":  OPCODE,
	"BR":   OPCODE,
	"JMP":  OPCODE,
	"JSR":  OPCODE,
	"JSRR": OPCODE,
	"RET":  OPCODE,
	"RTI":  OPCODE,

	// Directives
	"END":     DIRECTIVE,
	"ORIG":    DIRECTIVE,
	"FILL":    DIRECTIVE,
	"BLKW":    DIRECTIVE,
	"STRINGZ": DIRECTIVE,
	"BEGIN":   DIRECTIVE,

	// TRAPS
	"TRAP":  TRAP,
	"GETC":  TRAP,
	"OUT":   TRAP,
	"PUTS":  TRAP,
	"IN":    TRAP,
	"PUTSP": TRAP,
	"HALT":  TRAP,

	// Registers
	"R0": REGISTER,
	"R1": REGISTER,
	"R2": REGISTER,
	"R3": REGISTER,
	"R4": REGISTER,
	"R5": REGISTER,
	"R6": REGISTER,
	"R7": REGISTER,

	"BRn":   OPCODE,
	"BRnz":  OPCODE,
	"BRnzp": OPCODE,
	"BRnp":  OPCODE,
	"BRnpz": OPCODE,

	"BRz":   OPCODE,
	"BRzn":  OPCODE,
	"BRznp": OPCODE,
	"BRzp":  OPCODE,
	"BRzpn": OPCODE,

	"BRp":   OPCODE,
	"BRpz":  OPCODE,
	"BRpzn": OPCODE,
	"BRpn":  OPCODE,
	"BRpnz": OPCODE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
