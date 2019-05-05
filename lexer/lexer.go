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

	tok, newCharIndex := tokenize(input, charIndex)

	accumulator = append(accumulator, tok)

	return lex(input, newCharIndex, accumulator)
}

func tokenize(input string, charIndex int) (tok token.Token, newCharIndex int) {
	ch := input[charIndex]

	switch ch {
	case '(':
		return newToken(token.LEFT_PAREN, ch), charIndex + 1
	case ')':
		return newToken(token.RIGHT_PAREN, ch), charIndex + 1
	default:
		if isDigit(ch) {
			return readNumber(input, charIndex)
		}
		return readSymbol(input, charIndex)
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

func readNumber(input string, charIndex int) (tok token.Token, newCharIndex int) {
	literalSlice := []byte{input[charIndex]}
	newCharIndex = charIndex + 1
	numbersOfPeriod := 0

	for newCharIndex < len(input) &&
		(isDigit(input[newCharIndex]) || input[newCharIndex] == '.') {

		literalSlice = append(literalSlice, input[newCharIndex])

		if input[newCharIndex] == '.' {
			numbersOfPeriod = numbersOfPeriod + 1
		}

		newCharIndex = newCharIndex + 1
	}

	if numbersOfPeriod > 1 {
		return token.Token{Type: token.UNKNOWN, Literal: string(literalSlice)}, newCharIndex
	}

	return token.Token{Type: token.NUMBER, Literal: string(literalSlice)}, newCharIndex
}

func readSymbol(input string, charIndex int) (tok token.Token, newCharIndex int) {
	literalSlice := []byte{input[charIndex]}
	newCharIndex = charIndex + 1

	for newCharIndex < len(input) &&
		(input[newCharIndex] != '(' && input[newCharIndex] != ')' && input[newCharIndex] != ' ') {

		literalSlice = append(literalSlice, input[newCharIndex])

		newCharIndex = newCharIndex + 1
	}

	return token.Token{Type: token.SYMBOL, Literal: string(literalSlice)}, newCharIndex
}
