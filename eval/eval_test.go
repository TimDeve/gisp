package eval

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/value"
)

func TestEvalInt(t *testing.T) {
	result, err := Eval("1")

	if err != nil {
		t.Errorf(`Eval("1") returned an error: %s`, err.Error())
		return
	}

	if result.GetType() != value.NUMBER {
		t.Errorf(`Did not return a number`)
		return
	}

	number := result.(value.Number)
	if number.GetValue() != 1.0 {
		t.Errorf(`Eval("1") = %f; want 1.0`, result)
	}
}

func TestEvalFloat(t *testing.T) {
	result, err := Eval("1.35")

	if err != nil {
		t.Errorf(`Eval("1.35") returned an error: %s`, err.Error())
	}

	if result.GetType() != value.NUMBER {
		t.Errorf(`Did not return a number`)
		return
	}

	number := result.(value.Number)
	if number.GetValue() != 1.35 {
		t.Errorf(`Eval("1.35") = %f; want 1.35`, result)
	}
}

func TestEvalList(t *testing.T) {
	result, err := Eval("(1 1.0 2.4)")

	expected := value.Sexp{
		[]value.Value{
			value.Number{1.0},
			value.Number{1.0},
			value.Number{2.4},
		},
	}

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestEvalAdd(t *testing.T) {
	result, err := Eval("(+ 1 1)")

	if err != nil {
		t.Errorf(`Eval("(+ 1 1)") returned an error: %s`, err.Error())
	}

	if result.GetType() != value.NUMBER {
		t.Errorf(`Did not return a number`)
		return
	}

	number := result.(value.Number)
	if number.GetValue() != 2.0 {
		t.Errorf(`Eval("(+ 1 1)") = %f; want 2.0`, result)
	}
}

func TestEvalReturnsLastValue(t *testing.T) {
	result, err := Eval("(+ 1 1) (+ 1 2)")

	if err != nil {
		t.Errorf(`Eval("(+ 1 1) (+ 1 2)") returned an error: %s`, err.Error())
	}

	if result.GetType() != value.NUMBER {
		t.Errorf(`Did not return a number`)
		return
	}

	number := result.(value.Number)
	if number.GetValue() != 3.0 {
		t.Errorf(`Eval("(+ 1 1) (+ 1 2)") = %f; want 2.0`, result)
	}
}
