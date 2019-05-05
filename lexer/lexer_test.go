package main

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/token"
)

func TestShouldLexWithSingleNumber(t *testing.T) {
	result := Lex("1")
	expected := []token.Token{
		token.Token{
			Type:    token.NUMBER,
			Literal: "1",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal. Expected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithListOfSingleNumbers(t *testing.T) {
	result := Lex("(1 3 4 5)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.NUMBER, "1"},
		{token.NUMBER, "3"},
		{token.NUMBER, "4"},
		{token.NUMBER, "5"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal. Expected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithListOfMultiNumber(t *testing.T) {
	result := Lex("(100 32)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.NUMBER, "100"},
		{token.NUMBER, "32"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal. Expected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldLexWithFloat(t *testing.T) {
	result := Lex("(100.45 32.22)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.NUMBER, "100.45"},
		{token.NUMBER, "32.22"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal. Expected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldReturnUnknownForMalformedFloat(t *testing.T) {
	result := Lex("(10..0.45 2 32.22.4)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.UNKNOWN, "10..0.45"},
		{token.NUMBER, "2"},
		{token.UNKNOWN, "32.22.4"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal. Expected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleSymbols(t *testing.T) {
	result := Lex("(+ 2 3.50)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.SYMBOL, "+"},
		{token.NUMBER, "2"},
		{token.NUMBER, "3.50"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldHandleMultiCharacterSymbols(t *testing.T) {
	result := Lex("(add 2 3.50)")
	expected := []token.Token{
		{token.LEFT_PAREN, "("},
		{token.SYMBOL, "add"},
		{token.NUMBER, "2"},
		{token.NUMBER, "3.50"},
		{token.RIGHT_PAREN, ")"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}
