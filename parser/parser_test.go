package parser

import (
	"errors"
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/token"
	"github.com/TimDeve/gisp/value"
)

func TestShouldConvertTokenToValues(t *testing.T) {
	result, err := Parse([]token.Token{
		{token.NUMBER, "1.0", nil},
		{token.SYMBOL, "+", nil},
	})

	expected := []value.Value{
		value.NewNumber(1.0),
		value.NewSymbol("+"),
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestShouldConvertTokenWithSexpToValues(t *testing.T) {
	result, err := Parse([]token.Token{
		{token.SEXP, "", []token.Token{
			{token.NUMBER, "1.0", nil},
			{token.SYMBOL, "+", nil},
		}},
	})

	expected := []value.Value{
		value.NewSexp(
			[]value.Value{
				value.NewNumber(1.0),
				value.NewSymbol("+"),
			},
		),
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestShouldExitWithErrorIfUnknownTokenPresent(t *testing.T) {
	expectedError := errors.New("Parse error: invalid token")

	_, err := Parse([]token.Token{
		{token.UNKNOWN, "", nil},
		{token.NUMBER, "2.0", nil},
	})

	if err == nil {
		t.Errorf("Expected error for UNKNOWN value but received nil error")
	}

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("Wrong error.\nExpected:\n%s\nReceived:\n%s", expectedError.Error(), err.Error())
	}
}
