package eval

import "testing"

func TestParseInt(t *testing.T) {
	result, err := Eval("1")

	if err != nil {
		t.Errorf(`Eval("1") returned an error: %s`, err.Error())
	}

	if result != 1.0 {
		t.Errorf(`Eval("1") = %f; want 1.0`, result)
	}
}

func TestParseFloat(t *testing.T) {
	result, err := Eval("1.35")

	if err != nil {
		t.Errorf(`Eval("1.35") returned an error: %s`, err.Error())
	}

	if result != 1.35 {
		t.Errorf(`Eval("1.35") = %f; want 1.35`, result)
	}
}
