package stdlib

import (
	"reflect"
	"testing"

	"github.com/TimDeve/gisp/value"
)

func TestShouldReturnFalseOkWhenFunctionDoesNotExist(t *testing.T) {
	_, ok := GetFunc("thisshouldnotexist")

	if ok {
		t.Errorf("GetFunc(\"thisshouldnotexist\") returned true. Expected false")
	}
}

func TestAddShouldReturnZeroWhenNoArgument(t *testing.T) {
	add, _ := GetFunc("+")

	expected := value.NewNumber(0)

	result, err := add([]value.Value{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldAddTwoNumbers(t *testing.T) {
	add, _ := GetFunc("+")

	expected := value.NewNumber(2.0)

	result, err := add([]value.Value{
		value.NewNumber(1.0),
		value.NewNumber(1.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestShouldAddNNumbers(t *testing.T) {
	add, _ := GetFunc("+")

	expected := value.NewNumber(15.0)

	result, err := add([]value.Value{
		value.NewNumber(1.0),
		value.NewNumber(2.0),
		value.NewNumber(3.0),
		value.NewNumber(4.0),
		value.NewNumber(5.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestSubstractShouldReturnZeroWhenNoArgument(t *testing.T) {
	add, _ := GetFunc("-")

	expected := value.NewNumber(0)

	result, err := add([]value.Value{})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestSubstractShouldMakeSingleNumberNegative(t *testing.T) {
	add, _ := GetFunc("-")

	expected := value.NewNumber(-42)

	result, err := add([]value.Value{
		value.NewNumber(42),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestSubstractShouldSubstractTwoNumbers(t *testing.T) {
	add, _ := GetFunc("-")

	expected := value.NewNumber(2.0)

	result, err := add([]value.Value{
		value.NewNumber(3.0),
		value.NewNumber(1.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestSubstractShouldSubtractNNumbers(t *testing.T) {
	add, _ := GetFunc("-")

	expected := value.NewNumber(1)

	result, err := add([]value.Value{
		value.NewNumber(15.0),
		value.NewNumber(2.0),
		value.NewNumber(3.0),
		value.NewNumber(4.0),
		value.NewNumber(5.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%+v\nReceived:\n%+v", expected, result)
	}
}

func TestEqualShouldBeReturnErrorWithNoValue(t *testing.T) {
	equal, ok := GetFunc("=")
	if !ok {
		t.Errorf("= function not found")
		return
	}

	_, err := equal([]value.Value{})

	if err == nil {
		t.Errorf("Should return an error")
		return
	}

	if err.Error() != "Wrong number of argugments: 0" {
		t.Errorf("\nError message should be:\nWrong number of argugments: 0\nWas:\n%s", err.Error())
	}
}

func TestEqualShouldBeTrueWithOneValue(t *testing.T) {
	equal, ok := GetFunc("=")
	if !ok {
		t.Errorf("= function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewNumber(15.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestEqualShouldBeTrueWithEqualValues(t *testing.T) {
	equal, ok := GetFunc("=")
	if !ok {
		t.Errorf("= function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewNumber(15.0),
		value.NewNumber(15.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestEqualShouldBeTrueWithMoreThanTwoEqualValues(t *testing.T) {
	equal, ok := GetFunc("=")
	if !ok {
		t.Errorf("= function not found")
		return
	}

	expected := value.NewBoolean(true)

	result, err := equal([]value.Value{
		value.NewNumber(15.0),
		value.NewNumber(15.0),
		value.NewNumber(15.0),
		value.NewNumber(15.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}

func TestEqualShouldBeFalseWithDifferentValues(t *testing.T) {
	equal, ok := GetFunc("=")
	if !ok {
		t.Errorf("= function not found")
		return
	}

	expected := value.NewBoolean(false)

	result, err := equal([]value.Value{
		value.NewNumber(15.0),
		value.NewNumber(1.0),
	})

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Not equal.\nExpected:\n%#v\nReceived:\n%#v", expected, result)
	}
}
