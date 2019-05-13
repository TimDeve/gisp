package lexer

import (
	"reflect"
	"testing"

	"gisp/token"
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

func TestShouldLexWithSingleNegativeNumber(t *testing.T) {
	result := Lex("-1")
	expected := []token.Token{
		token.Token{
			Type:     token.NUMBER,
			Literal:  "-1",
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
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "1", Children: nil},
			{Type: token.NUMBER, Literal: "3", Children: nil},
			{Type: token.NUMBER, Literal: "4", Children: nil},
			{Type: token.NUMBER, Literal: "5", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithListOfMultiNumber(t *testing.T) {
	result := Lex("(100 32)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "100", Children: nil},
			{Type: token.NUMBER, Literal: "32", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithFloat(t *testing.T) {
	result := Lex("(100.45 32.22)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "100.45", Children: nil},
			{Type: token.NUMBER, Literal: "32.22", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldReturnUnknownForMalformedFloat(t *testing.T) {
	result := Lex("(10..0.45 2 32.22.4)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.UNKNOWN, Literal: "10..0.45", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.UNKNOWN, Literal: "32.22.4", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleSymbols(t *testing.T) {
	result := Lex("(+ 2 3.50)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.SYMBOL, Literal: "+", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.NUMBER, Literal: "3.50", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleMultiCharacterSymbols(t *testing.T) {
	result := Lex("(add 2 3.50)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.SYMBOL, Literal: "add", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.NUMBER, Literal: "3.50", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleNestedSexp(t *testing.T) {
	result := Lex("(add 2 (add 3 4.50))")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.SYMBOL, Literal: "add", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.SEXP, Literal: "", Children: []token.Token{
				{Type: token.SYMBOL, Literal: "add", Children: nil},
				{Type: token.NUMBER, Literal: "3", Children: nil},
				{Type: token.NUMBER, Literal: "4.50", Children: nil},
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
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "1", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.SEXP, Literal: "", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldReturnUnknownForUnclosedParen(t *testing.T) {
	result := Lex("(1 2 (1 2)")
	expected := []token.Token{
		{Type: token.UNKNOWN, Literal: "(1 2 (1 2)", Children: nil},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleAllWhiteSpace(t *testing.T) {
	result := Lex("(1\r2\v\n \f \t 3)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "1", Children: nil},
			{Type: token.NUMBER, Literal: "2", Children: nil},
			{Type: token.NUMBER, Literal: "3", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestShouldHandleWhiteSpaceAfterASymbol(t *testing.T) {
	result := Lex("(add\n3)")
	expected := []token.Token{
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.SYMBOL, Literal: "add", Children: nil},
			{Type: token.NUMBER, Literal: "3", Children: nil},
		}},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}
