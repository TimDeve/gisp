package stdlib

import (
	"errors"

	"github.com/TimDeve/gisp/value"
)

func add(values []value.Value) (value.Value, error) {
	result := 0.0
	for _, val := range values {
		if val.GetType() == value.NUMBER {
			v := val.(value.Number)
			result = result + v.GetValue()
		} else {
			return value.Number{}, errors.New("Add error: value passed not a number")
		}
	}
	return value.Number{result}, nil
}

var stdlibMap = map[string]func([]value.Value) (value.Value, error){
	"+":   add,
	"add": add,
}

func GetLib(name string) (lib func([]value.Value) (value.Value, error), ok bool) {
	lib, ok = stdlibMap[name]
	return
}
