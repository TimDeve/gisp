package token

type TokenType string

const (
	LEFT_PAREN  TokenType = "LEFT_PAREN"
	RIGHT_PAREN TokenType = "RIGHT_PAREN"
	NUMBER      TokenType = "NUMBER"
	UNKNOWN     TokenType = "UNKNOWN"
)

type Token struct {
	Type    TokenType
	Literal string
}
