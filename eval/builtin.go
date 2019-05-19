package eval

import (
	"fmt"
	"gisp/value"
)

var builtinMap map[string]builtIn

type builtIn = func(values []value.Value) (value.Value, error)

func init() {
	builtinMap = map[string]builtIn{
		"if": ifBuiltIn,
	}
}

func ifBuiltIn(values []value.Value) (value.Value, error) {
	vLen := len(values)
	if vLen == 2 || vLen == 3 {
		conditional, err := evalValue(values[0])
		if err != nil {
			return value.NewNothing(), err
		}
		if value.IsBoolean(conditional) {
			c := conditional.(value.Boolean)

			if !c.GetValue() {
				if vLen < 3 {
					return value.NewNothing(), nil
				}
				return evalValue(values[2])
			}
		}

		return evalValue(values[1])
	}

	return value.NewNothing(), fmt.Errorf("wrong number of args: %d", vLen)
}

func getBuiltin(name string) (lib builtIn, ok bool) {
	lib, ok = builtinMap[name]
	return
}
