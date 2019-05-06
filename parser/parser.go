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

func parseToken(tok token.Token) (value value.Value, err error) {
	switch tok.Type {
	case token.UNKNOWN:
		return value, errors.New("Parse error: invalid token")
	case token.NUMBER:
		num, err := parseNumber(tok)
		if err != nil {
			return value, err
		}
		return num, nil
	case token.SYMBOL:
		sym := parseSymbol(tok)
		return sym, nil
	case token.SEXP:
		return parseSexp(tok)
	default:
		return value, errors.New("Parse error: unknown token")
	}
}

func parseNumber(tok token.Token) (num value.Number, err error) {
	f, err := strconv.ParseFloat(tok.Literal, 64)
	if err != nil {
		return num, err
	}
	num = value.Number{f}
	return
}

func parseSymbol(tok token.Token) (sym value.Symbol) {
	return value.Symbol{tok.Literal}
}

func parseSexp(tok token.Token) (sexp value.Sexp, err error) {
	vals, err := Parse(tok.Children)
	sexp.Value = vals
	return
}
