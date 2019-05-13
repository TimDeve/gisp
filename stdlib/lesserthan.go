package stdlib

import (
	"errors"

	"gisp/value"
)

func lesserThan(values []value.Value) (value.Value, error) {
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

			if prevVal.GetValue() >= currentVal.GetValue() {
				return value.NewBoolean(false), nil
			}
		}
	}

	return value.NewBoolean(true), nil
}
