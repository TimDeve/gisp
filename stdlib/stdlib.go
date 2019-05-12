package stdlib

import (
	"errors"

	"github.com/TimDeve/gisp/value"
)

func add(values []value.Value) (value.Value, error) {
	result := 0.0
	for _, val := range values {
		if value.IsNumber(val) {
			v := val.(value.Number)
			result = result + v.GetValue()
		} else {
			return value.NewNothing(), errors.New("Add error: value passed not a number")
		}
	}
	return value.NewNumber(result), nil
}

func substract(values []value.Value) (value.Value, error) {
	result := 0.0
	for i, val := range values {
		if value.IsNumber(val) {
			v := val.(value.Number)
			if i == 0 && len(values) > 1 {
				result = result + v.GetValue()
			} else {
				result = result - v.GetValue()
			}
		} else {
			return value.NewNothing(), errors.New("Add error: value passed not a number")
		}
	}
	return value.NewNumber(result), nil
}

func equal(values []value.Value) (value.Value, error) {
	result := value.NewBoolean(true)

	if len(values) == 0 {
		return value.NewNothing(), errors.New("Wrong number of argugments: 0")
	} else if len(values) == 1 {
		return result, nil
	}

	for i := 1; i < len(values); i++ {
		if !values[i].Equals(values[i-1]) {
			return value.NewBoolean(false), nil
		}
	}

	return result, nil
}

func greaterThan(values []value.Value) (value.Value, error) {
	if len(values) == 0 {
		return value.NewBoolean(false), nil
	} else if len(values) == 1 {
		return value.NewBoolean(true), nil
	}

	for i := 0; i < len(values); i++ {
		if !value.IsNumber(values[i]) {
			return value.NewNothing(), errors.New("All arguments must be numbers")
		}
		if i != 0 {
			prevVal := values[i-1].(value.Number)
			currentVal := values[i].(value.Number)

			if prevVal.GetValue() <= currentVal.GetValue() {
				return value.NewBoolean(false), nil
			}
		}
	}

	return value.NewBoolean(true), nil
}

func greaterOrEqual(values []value.Value) (value.Value, error) {
	if len(values) == 0 || len(values) == 1 {
		return value.NewBoolean(true), nil
	}

	for i := 0; i < len(values); i++ {
		if !value.IsNumber(values[i]) {
			return value.NewNothing(), errors.New("All arguments must be numbers")
		}
		if i != 0 {
			prevVal := values[i-1].(value.Number)
			currentVal := values[i].(value.Number)

			if prevVal.GetValue() < currentVal.GetValue() {
				return value.NewBoolean(false), nil
			}
		}
	}

	return value.NewBoolean(true), nil
}

var stdlibMap = map[string]func([]value.Value) (value.Value, error){
	"+":  add,
	"-":  substract,
	"=":  equal,
	">":  greaterThan,
	">=": greaterOrEqual,
	"<":  lesserThan,
	"<=": lesserOrEqual,
}

func GetFunc(name string) (lib func([]value.Value) (value.Value, error), ok bool) {
	lib, ok = stdlibMap[name]
	return
}
