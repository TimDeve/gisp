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

var stdlibMap = map[string]func([]value.Value) (value.Value, error){
	"+": add,
	"-": substract,
}

func GetFunc(name string) (lib func([]value.Value) (value.Value, error), ok bool) {
	lib, ok = stdlibMap[name]
	return
}
