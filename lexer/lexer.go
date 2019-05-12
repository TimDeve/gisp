package lexer

import (
	"github.com/TimDeve/gisp/token"
)

func Lex(input string) []token.Token {
	return lex(input, 0, nil)
}

func lex(input string, charIndex int, accumulator []token.Token) []token.Token {
	if charIndex >= len(input) {
		return accumulator
	}

	if isWhiteSpace(input[charIndex]) {
		return lex(input, charIndex+1, accumulator)
	}

	tok, newCharIndex := tokenize(input, charIndex)

	accumulator = append(accumulator, tok)

	return lex(input, newCharIndex, accumulator)
}

func tokenize(input string, charIndex int) (tok token.Token, newCharIndex int) {
	if input[charIndex] == '(' {
		return readSexp(input, charIndex)
	} else if couldBeNumber(input, charIndex) {
		return readNumber(input, charIndex)
	} else {
		return readSymbol(input, charIndex)
	}
}

func couldBeNumber(input string, charIndex int) bool {
	ch := input[charIndex]
	return isDigit(ch) ||
		(ch == '-' || ch == '+') && (charIndex+1 <= len(input) && isDigit(input[charIndex+1]))
}

func isWhiteSpace(b byte) bool {
	return b == ' ' ||
		b == '\n' ||
		b == '\t' ||
		b == '\r' ||
		b == '\v' ||
		b == '\f'
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
		(input[newCharIndex] != '(' && input[newCharIndex] != ')' && !isWhiteSpace(input[newCharIndex])) {

		literalSlice = append(literalSlice, input[newCharIndex])

		newCharIndex = newCharIndex + 1
	}

	return token.Token{Type: token.SYMBOL, Literal: string(literalSlice)}, newCharIndex
}

func readSexp(input string, charIndex int) (tok token.Token, newCharIndex int) {
	literalSlice := []byte{input[charIndex]}
	newCharIndex = charIndex + 1
	numberOfLeftParen := 1
	numberOfRightParen := 0

	for {
		if newCharIndex >= len(input) {
			return token.Token{Type: token.UNKNOWN, Literal: string(literalSlice)}, newCharIndex
		}

		literalSlice = append(literalSlice, input[newCharIndex])

		if input[newCharIndex] == '(' {
			numberOfLeftParen = numberOfLeftParen + 1
		}

		if input[newCharIndex] == ')' {
			numberOfRightParen = numberOfRightParen + 1

			if numberOfLeftParen == numberOfRightParen {
				literalWithoutSurroundingParen := literalSlice[1 : len(literalSlice)-1]
				return token.Token{
					Type:     token.SEXP,
					Literal:  "",
					Children: Lex(string(literalWithoutSurroundingParen)),
				}, newCharIndex + 1
			}
		}

		newCharIndex = newCharIndex + 1
	}
}
