package stdlib

import (
	"reflect"
	"testing"

	"gisp/value"
)

func TestLesserOrEqualThanShouldReturnTrueWithNoArguments(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserOrEqualThanShouldReturnTrueWithOneArgument(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
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

func TestLesserOrEqualThanAnArgLesserThanAnOther(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
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

func TestLesserOrEqualThanAnArgLesserThanMultipleOthers(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
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

func TestLesserOrEqualThanAnArgNotLesserThanMultipleOthers(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
		return
	}

	expected := value.NewBoolean(false)

	result, err := equal([]value.Value{
		value.NewNumber(1),
		value.NewNumber(10),
		value.NewNumber(3),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestLesserOrEqualThanAnArgEqualToAnOther(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
		return
	}

	expected := value.NewBoolean(true)

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

func TestLesserOrEqualThanShouldReturnErrorIfNotAllValuesNumbers(t *testing.T) {
	equal, ok := GetFunc("<=")
	if !ok {
		t.Errorf("<= function not found")
		return
	}

	expectedErrorMessage := "all arguments must be numbers"

	_, err := equal([]value.Value{
		value.NewBoolean(true),
		value.NewNumber(1.0),
	})

	if err == nil {
		t.Errorf("Should return an error")
		return
	}

	if err.Error() != expectedErrorMessage {
		t.Errorf("\nError message should be:\n%s\nWas:\n%s", expectedErrorMessage, err.Error())
	}
}
