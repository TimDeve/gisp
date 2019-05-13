package token

type Type string

const (
	NUMBER  Type = "NUMBER"
	SYMBOL  Type = "SYMBOL"
	UNKNOWN Type = "UNKNOWN"
	SEXP    Type = "SEXP"
)

type Token struct {
	Type     Type
	Literal  string
	Children []Token
}
