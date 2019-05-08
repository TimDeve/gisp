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
	var newValues []value.Value
	for _, val := range expr.GetValue() {
		if val.GetType() == value.SEXP {
			newVal, err := EvalExpression(val.(value.Sexp))
			if err != nil {
				return value.Number{}, err
			}
			newValues = append(newValues, newVal)
		} else {
			newValues = append(newValues, val)
		}
	}

	first := newValues[0]

	if first.GetType() == value.SYMBOL {
		rest := newValues[1:len(newValues)]

		firstSym := first.(value.Symbol)
		f, ok := stdlib.GetFunc(firstSym.GetValue())
		if !ok {
			return value.Number{}, errors.New("Symbol not found")
		}
		return f(rest)
	} else {
		return value.Sexp{newValues}, nil
	}
}
