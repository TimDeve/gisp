package eval

import (
	"testing"

	"github.com/TimDeve/gisp/value"
)

func TestParseInt(t *testing.T) {
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

func TestParseFloat(t *testing.T) {
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
