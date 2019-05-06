package value

type ValueType string

const (
	NUMBER ValueType = "NUMBER"
	SYMBOL ValueType = "SYMBOL"
	SEXP   ValueType = "SEXP"
)

type Value interface {
	GetType() ValueType
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

type Symbol struct {
	Value string
}

func (s Symbol) GetType() ValueType {
	return SYMBOL
}

func (s *Symbol) GetValue() string {
	return s.Value
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
