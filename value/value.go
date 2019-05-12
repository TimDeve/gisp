package value

import (
	"fmt"
	"strings"
)

type valueType string

const (
	nothing valueType = "nothing"
	boolean valueType = "boolean"
	number  valueType = "number"
	symbol  valueType = "symbol"
	sexp    valueType = "sexp"
)

type Value interface {
	getType() valueType
	String() string
}

type Nothing struct{}

func (n Nothing) getType() valueType {
	return nothing
}

func (n Nothing) String() string {
	return ""
}

func NewNothing() Nothing {
	return Nothing{}
}

func IsNothing(val Value) bool {
	return val.getType() == nothing
}

type Boolean struct {
	value bool
}

func (n Boolean) getType() valueType {
	return boolean
}

func (n *Boolean) GetValue() bool {
	return n.value
}

func (n Boolean) String() string {
	return fmt.Sprint(n.GetValue())
}

func NewBoolean(val bool) Boolean {
	return Boolean{value: val}
}

func IsBoolean(val Value) bool {
	return val.getType() == boolean
}

type Number struct {
	value float64
}

func (n Number) getType() valueType {
	return number
}

func (n *Number) GetValue() float64 {
	return n.value
}

func (n Number) String() string {
	return fmt.Sprint(n.GetValue())
}

func NewNumber(val float64) Number {
	return Number{value: val}
}

func IsNumber(val Value) bool {
	return val.getType() == number
}

type Symbol struct {
	value string
}

func (s Symbol) getType() valueType {
	return symbol
}

func (s *Symbol) GetValue() string {
	return s.value
}

func (s Symbol) String() string {
	return s.GetValue()
}

func NewSymbol(val string) Symbol {
	return Symbol{value: val}
}

func IsSymbol(val Value) bool {
	return val.getType() == symbol
}

type Sexp struct {
	value []Value
}

func (s Sexp) getType() valueType {
	return sexp
}

func (s *Sexp) GetValue() []Value {
	return s.value
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

func NewSexp(val []Value) Sexp {
	return Sexp{value: val}
}

func IsSexp(val Value) bool {
	return val.getType() == sexp
}
