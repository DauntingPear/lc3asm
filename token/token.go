package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	PERIOD    = "."
	HASH      = "#"

	IDENT    = "IDENT"    // Labels
	REGISTER = "REGISTER" // Registers R0-8

	ADD = "ADD"
	NOT = "NOT"
	AND = "AND"

	END = "END"

	INT = "INT"
	HEX = "HEX"
)

var keywords = map[string]TokenType{
	"ADD": ADD,
	"NOT": NOT,
	"AND": AND,
	"END": END,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
