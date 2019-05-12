package parser

import (
	"errors"
	"strconv"

	"github.com/TimDeve/gisp/token"
	"github.com/TimDeve/gisp/value"
)

func Parse(tokens []token.Token) (values []value.Value, err error) {
	for _, tok := range tokens {
		val, err := parseToken(tok)
		if err != nil {
			return values, err
		}
		values = append(values, val)
	}

	return
}

func parseToken(tok token.Token) (value.Value, error) {
	switch tok.Type {
	case token.UNKNOWN:
		return value.NewNothing(), errors.New("Parse error: invalid token")
	case token.NUMBER:
		num, err := parseNumber(tok)
		if err != nil {
			return value.NewNothing(), err
		}
		return num, nil
	case token.SYMBOL:
		b, ok := parseBoolean(tok)
		if ok {
			return b, nil
		}
		sym := parseSymbol(tok)
		return sym, nil
	case token.SEXP:
		return parseSexp(tok)
	default:
		return value.NewNothing(), errors.New("Parse error: unknown token")
	}
}

func parseNumber(tok token.Token) (num value.Number, err error) {
	f, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return num, err
	}
	num = value.NewNumber(f)
	return
}

func parseBoolean(tok token.Token) (val value.Boolean, ok bool) {
	ok = false

	if tok.Literal == "true" {
		return value.NewBoolean(true), true
	} else if tok.Literal == "false" {
		return value.NewBoolean(false), true
	}

	return
}

func parseSymbol(tok token.Token) (sym value.Symbol) {
	return value.NewSymbol(tok.Literal)
}

func parseSexp(tok token.Token) (value.Sexp, error) {
	vals, err := Parse(tok.Children)
	return value.NewSexp(vals), err
}
