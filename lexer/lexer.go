package main

import (
	"github.com/TimDeve/gisp/token"
)

func Lex(input string) []token.Token {
	return lex(input, 0, []token.Token{})
}

func lex(input string, charIndex int, accumulator []token.Token) []token.Token {
	if charIndex >= len(input) {
		return accumulator
	}

	if input[charIndex] == ' ' {
		return lex(input, charIndex+1, accumulator)
	}

	accumulator = append(accumulator, tokenize(input[charIndex]))

	return lex(input, charIndex+1, accumulator)
}

func tokenize(ch byte) token.Token {
	switch ch {
	case '(':
		return newToken(token.LEFT_PAREN, ch)
	case ')':
		return newToken(token.RIGHT_PAREN, ch)
	default:
		if isDigit(ch) {
			return newToken(token.NUMBER, ch)
		}
		return newToken(token.UNKNOWN, 0)
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string([]byte{ch}),
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
