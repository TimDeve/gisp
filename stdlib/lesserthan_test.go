package stdlib

import (
	"reflect"
	"testing"

	"gisp/value"
)

func TestLesserThanShouldReturnFalseWithNoArguments(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(false)

	result, err := equal([]value.Value{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserThanShouldReturnTrueWithOneArgument(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewBoolean(false),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserThanAnArgGreaterThanAnOther(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewNumber(2),
		value.NewNumber(4),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserThanAnArgGreaterThanMultipleOthers(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewNumber(1),
		value.NewNumber(2),
		value.NewNumber(4),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserThanAnArgNotGreaterThanMultipleOthers(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(false)

	result, err := equal([]value.Value{
		value.NewNumber(1),
		value.NewNumber(10),
		value.NewNumber(2),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserThanAnArgEqualToAnOther(t *testing.T) {
	equal, ok := GetFunc("<")
	if !ok {
		t.Errorf("< function not found")
		return
	}

	expected := value.NewBoolean(false)

	result, err := equal([]value.Value{
		value.NewNumber(10),
		value.NewNumber(10),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}
