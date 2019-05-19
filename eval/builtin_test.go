package eval

import (
	"errors"
	"gisp/value"
	"reflect"
	"testing"
)

func TestIfBuiltInShouldReturnAnErrorIfGivenNoArgs(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expectedError := errors.New("wrong number of args: 0")

	_, err := ifBuiltIn([]value.Value{})

	if err == nil {
		t.Errorf("Expected error for no args but got nil")
		return
	}

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("Wrong error.\nExpected:\n%s\nReceived:\n%s", expectedError.Error(), err.Error())
	}
}

func TestIfBuiltInShouldReturnErrorIfOnlyHasAConditionButNoBranches(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expectedError := errors.New("wrong number of args: 1")

	_, err := ifBuiltIn([]value.Value{value.NewBoolean(true)})

	if err == nil {
		t.Errorf("Expected error for one args but got nil")
		return
	}

	if err != nil && err.Error() != expectedError.Error() {
		t.Errorf("Wrong error.\nExpected:\n%s\nReceived:\n%s", expectedError.Error(), err.Error())
	}
}

func TestIfBuiltInShouldReturnFirstBranchIfConditionalIsTrue(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(1)

	result, err := ifBuiltIn([]value.Value{value.NewBoolean(true), value.NewNumber(1), value.NewNumber(2)})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldReturnSecondBranchIfConditionalIsFalse(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(2)

	result, err := ifBuiltIn([]value.Value{value.NewBoolean(false), value.NewNumber(1), value.NewNumber(2)})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldReturnNothingIfConditionalIsFalseAndThereIsNoSecondBranch(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNothing()

	result, err := ifBuiltIn([]value.Value{value.NewBoolean(false), value.NewNumber(1)})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldReturnFirstBranchIfConditionalNotABoolean(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(1)

	result, err := ifBuiltIn([]value.Value{value.NewNumber(666), value.NewNumber(1), value.NewNumber(2)})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldEvaluateTheConditional(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(2)

	result, err := ifBuiltIn([]value.Value{
		value.NewSexp([]value.Value{
			value.NewSymbol("if"),
			value.NewBoolean(true),
			value.NewBoolean(false),
		}),
		value.NewNumber(1),
		value.NewNumber(2),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldEvaluateFirstBranchWhenConditionalIsTrue(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(6)

	result, err := ifBuiltIn([]value.Value{
		value.NewBoolean(true),
		value.NewSexp([]value.Value{
			value.NewSymbol("+"),
			value.NewNumber(3),
			value.NewNumber(3),
		}),
		value.NewNumber(2),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestIfBuiltInShouldEvaluateSecondBranchWhenConditionalIsFalse(t *testing.T) {
	ifBuiltIn, ok := getBuiltin("if")
	if !ok {
		t.Errorf("if builtin not found")
		return
	}

	expected := value.NewNumber(7)

	result, err := ifBuiltIn([]value.Value{
		value.NewBoolean(false),
		value.NewSexp([]value.Value{
			value.NewSymbol("+"),
			value.NewNumber(3),
			value.NewNumber(3),
		}),
		value.NewSexp([]value.Value{
			value.NewSymbol("+"),
			value.NewNumber(3),
			value.NewNumber(4),
		}),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}
