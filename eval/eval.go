package eval

import (
	"errors"

	"gisp/lexer"
	"gisp/parser"
	"gisp/stdlib"
	"gisp/value"
)

func Eval(input string) (value.Value, error) {
	tokens := lexer.Lex(input)

	values, err := parser.Parse(tokens)
	if err != nil {
		return value.NewNumber(0.0), err
	}

	var newValues []value.Value

	for _, val := range values {
		if value.IsSexp(val) {
			newVal, err := evalExpression(val.(value.Sexp))
			if err != nil {
				return value.NewNothing(), err
			}
			newValues = append(newValues, newVal)
		} else {
			newValues = append(newValues, val)
		}
	}

	return newValues[len(values)-1], err
}

func evalValue(val value.Value) (value.Value, error) {
	if value.IsSexp(val) {
		v := val.(value.Sexp)
		return evalExpression(v)
	}

	return val, nil
}

func evalExpression(expr value.Sexp) (value.Value, error) {
	values := expr.GetValue()
	if value.IsSymbol(values[0]) {
		firstSym := values[0].(value.Symbol)
		b, ok := getBuiltin(firstSym.GetValue())
		if ok {
			return b(values[1:])
		}
	}

	var newValues []value.Value
	for _, val := range values {
		if value.IsSexp(val) {
			newVal, err := evalExpression(val.(value.Sexp))
			if err != nil {
				return value.NewNothing(), err
			}
			newValues = append(newValues, newVal)
		} else {
			newValues = append(newValues, val)
		}
	}

	first := newValues[0]

	if value.IsSymbol(first) {
		rest := newValues[1:]

		firstSym := first.(value.Symbol)

		f, ok := stdlib.GetFunc(firstSym.GetValue())
		if !ok {
			return value.NewNothing(), errors.New("symbol not found")
		}
		return f(rest)
	}

	return value.NewSexp(newValues), nil
}
