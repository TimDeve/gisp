package parser

import (
	"errors"
	"reflect"
	"testing"

	"gisp/token"
	"gisp/value"
)

func TestShouldConvertTokenToValues(t *testing.T) {
	result, err := Parse([]token.Token{
		{Type: token.NUMBER, Literal: "1.0", Children: nil},
		{Type: token.SYMBOL, Literal: "+", Children: nil},
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

func TestShouldConvertBooleansToValues(t *testing.T) {
	result, err := Parse([]token.Token{
		{Type: token.SYMBOL, Literal: "true", Children: nil},
		{Type: token.SYMBOL, Literal: "false", Children: nil},
	})

	expected := []value.Value{
		value.NewBoolean(true),
		value.NewBoolean(false),
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
		{Type: token.SEXP, Literal: "", Children: []token.Token{
			{Type: token.NUMBER, Literal: "1.0", Children: nil},
			{Type: token.SYMBOL, Literal: "+", Children: nil},
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
		{Type: token.UNKNOWN, Literal: "", Children: nil},
		{Type: token.NUMBER, Literal: "2.0", Children: nil},
	})

	if err == nil {
		t.Errorf("Expected error for UNKNOWN value but received nil error")
	}

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("Wrong error.\nExpected:\n%s\nReceived:\n%s", expectedError.Error(), err.Error())
	}
}
