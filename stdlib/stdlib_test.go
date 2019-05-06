package stdlib

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/value"
)

func TestShouldAddTwoNumbers(t *testing.T) {
	add, _ := GetLib("add")

	expected := value.Number{2.0}

	result, err := add([]value.Value{
		value.Number{1.0},
		value.Number{1.0},
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}
