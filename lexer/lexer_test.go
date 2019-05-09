package lexer

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/token"
)

func TestShouldLexWithSingleNumber(t *testing.T) {
	result := Lex("1")
	expected := []token.Token{
		token.Token{
			Type:     token.NUMBER,
			Literal:  "1",
			Children: nil,
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithListOfSingleNumbers(t *testing.T) {
	result := Lex("(1 3 4 5)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "1", nil},
			{token.NUMBER, "3", nil},
			{token.NUMBER, "4", nil},
			{token.NUMBER, "5", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithListOfMultiNumber(t *testing.T) {
	result := Lex("(100 32)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "100", nil},
			{token.NUMBER, "32", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithFloat(t *testing.T) {
	result := Lex("(100.45 32.22)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "100.45", nil},
			{token.NUMBER, "32.22", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldReturnUnknownForMalformedFloat(t *testing.T) {
	result := Lex("(10..0.45 2 32.22.4)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.UNKNOWN, "10..0.45", nil},
			{token.NUMBER, "2", nil},
			{token.UNKNOWN, "32.22.4", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleSymbols(t *testing.T) {
	result := Lex("(+ 2 3.50)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.SYMBOL, "+", nil},
			{token.NUMBER, "2", nil},
			{token.NUMBER, "3.50", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleMultiCharacterSymbols(t *testing.T) {
	result := Lex("(add 2 3.50)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.SYMBOL, "add", nil},
			{token.NUMBER, "2", nil},
			{token.NUMBER, "3.50", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleNestedSexp(t *testing.T) {
	result := Lex("(add 2 (add 3 4.50))")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.SYMBOL, "add", nil},
			{token.NUMBER, "2", nil},
			{token.SEXP, "", []token.Token{
				{token.SYMBOL, "add", nil},
				{token.NUMBER, "3", nil},
				{token.NUMBER, "4.50", nil},
			}},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleEmptySexp(t *testing.T) {
	result := Lex("(1 2 ())")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "1", nil},
			{token.NUMBER, "2", nil},
			{token.SEXP, "", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldReturnUnknownForUnclosedParen(t *testing.T) {
	result := Lex("(1 2 (1 2)")
	expected := []token.Token{
		{token.UNKNOWN, "(1 2 (1 2)", nil},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleAllWhiteSpace(t *testing.T) {
	result := Lex("(1\r2\v\n \f \t 3)")
	expected := []token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "1", nil},
			{token.NUMBER, "2", nil},
			{token.NUMBER, "3", nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}
