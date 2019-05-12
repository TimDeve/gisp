package value

import (
	"reflect"
	"testing"
)

func TestShouldStringifySexp(t *testing.T) {
	sexp := Sexp{[]Value{
		Symbol{"add"},
		Number{1.0},
		Number{1.0},
	}}

	result := sexp.String()
	expected := "(add 1 1)"

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldStringifyNestedSexp(t *testing.T) {
	sexp := Sexp{[]Value{
		Symbol{"add"},
		Number{1.0},
		Number{1.0},
		Sexp{[]Value{
			Symbol{"add"},
			Number{1.0},
			Number{1.0},
		}},
	}}

	result := sexp.String()
	expected := "(add 1 1 (add 1 1))"

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldCompareSexp(t *testing.T) {
	sexp1 := Sexp{[]Value{
		Symbol{"add"},
		Number{1.0},
		Number{1.0},
		Sexp{[]Value{
			Symbol{"add"},
			Number{1.0},
			Number{1.0},
			Boolean{true},
		}},
	}}

	sexp2 := Sexp{[]Value{
		Symbol{"add"},
		Number{1.0},
		Number{1.0},
		Sexp{[]Value{
			Symbol{"add"},
			Number{1.0},
			Number{1.0},
			Boolean{true},
		}},
	}}

	if !sexp1.Equals(sexp2) {
		t.Errorf("Not equal")
	}
}
