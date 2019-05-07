package eval

import (
	"github.com/TimDeve/gisp/lexer"
	"github.com/TimDeve/gisp/parser"
	"github.com/TimDeve/gisp/value"
)

func Eval(input string) (value.Value, error) {
	tokens := lexer.Lex(input)

	values, err := parser.Parse(tokens)
	if err != nil {
		return value.Number{0.0}, err
	}

	return values[0], err
}
