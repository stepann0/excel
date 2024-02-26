package data

// Types that formula can return and cell can hold
type ValueType int

const (
	NilType ValueType = iota // Empty cell
	NumberType
	StringType  // Text cell
	BooleanType // `TRUE` or `FALSE` literals or boolean functions
	AreaType    // Range of cells (e.g A1:A10). Only 1D allowed
	ErrorType
)

type Value interface {
	Type() ValueType
}

// Numbers
type Number[T float64 | int] struct {
	Val T
}

func (n Number[T]) Type() ValueType { return NumberType }

// String
type String struct {
	Val string
}

func (s String) Type() ValueType { return StringType }

// Bool
type Boolean struct {
	Val bool
}

func (b Boolean) Type() ValueType { return BooleanType }

// Area
type Area struct {
	Val []Value
}

func (a Area) Type() ValueType { return AreaType }

// Error
type Error struct {
	Msg error
}

func (e Error) Type() ValueType { return ErrorType }
