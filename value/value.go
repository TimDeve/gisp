package value

import (
	"fmt"
	"strings"
)

type ValueType string

const (
	NOTHING ValueType = "NOTHING"
	NUMBER  ValueType = "NUMBER"
	SYMBOL  ValueType = "SYMBOL"
	SEXP    ValueType = "SEXP"
)

type Value interface {
	GetType() ValueType
	String() string
}

type Nothing struct{}

func (n Nothing) GetType() ValueType {
	return NOTHING
}

func (n Nothing) String() string {
	return ""
}

type Number struct {
	Value float64
}

func (n Number) GetType() ValueType {
	return NUMBER
}

func (n *Number) GetValue() float64 {
	return n.Value
}

func (n Number) String() string {
	return fmt.Sprint(n.GetValue())
}

type Symbol struct {
	Value string
}

func (s Symbol) GetType() ValueType {
	return SYMBOL
}

func (s *Symbol) GetValue() string {
	return s.Value
}

func (s Symbol) String() string {
	return s.GetValue()
}

type Sexp struct {
	Value []Value
}

func (s Sexp) GetType() ValueType {
	return SEXP
}

func (s *Sexp) GetValue() []Value {
	return s.Value
}

func (s Sexp) String() string {
	var str strings.Builder
	str.WriteString("(")

	for i, val := range s.GetValue() {
		str.WriteString(val.String())

		if i != len(s.GetValue())-1 {
			str.WriteString(" ")
		}
	}

	str.WriteString(")")
	return str.String()
}
