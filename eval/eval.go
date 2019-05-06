package eval

import "strconv"

func Eval(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}
