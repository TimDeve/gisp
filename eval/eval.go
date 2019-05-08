package eval

import (
	"errors"

	"github.com/TimDeve/gisp/lexer"
	"github.com/TimDeve/gisp/parser"
	"github.com/TimDeve/gisp/stdlib"
	"github.com/TimDeve/gisp/value"
)

func Eval(input string) (value.Value, error) {
	tokens := lexer.Lex(input)

	values, err := parser.Parse(tokens)
	if err != nil {
		return value.Number{0.0}, err
	}

	var newValues []value.Value

	for _, val := range values {
		if val.GetType() == value.SEXP {
			newVal, err := EvalExpression(val.(value.Sexp))
			if err != nil {
				return value.Number{0.0}, err
			}
			newValues = append(newValues, newVal)
		} else {
			newValues = append(newValues, val)
		}
	}

	return newValues[len(values)-1], err
}

func EvalExpression(expr value.Sexp) (value.Value, error) {
	values := expr.GetValue()
	first := values[0]
	rest := values[1:len(values)]

	if first.GetType() == value.SYMBOL {
		v := first.(value.Symbol)
		f, ok := stdlib.GetLib(v.GetValue())
		if !ok {
			return value.Number{}, errors.New("Symbol not found")
		}
		res, err := f(rest)
		return res, err
	} else {
		return expr, nil
	}
}
