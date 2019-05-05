package token

type TokenType string

const (
	NUMBER  TokenType = "NUMBER"
	SYMBOL  TokenType = "SYMBOL"
	UNKNOWN TokenType = "UNKNOWN"
	SEXP    TokenType = "SEXP"
)

type Token struct {
	Type     TokenType
	Literal  string
	Children []Token
}
