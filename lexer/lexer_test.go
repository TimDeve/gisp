package main

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/token"
)

func TestLexWithSingleNumber(t *testing.T) {
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

func TestLexWithListOfSingleNumbers(t *testing.T) {
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
