package value

// Types that formula can return and cell can hold
type ValueType int

const (
	NilType     ValueType = iota // Empty cell
	BooleanType                  // `TRUE` or `FALSE` literals or boolean functions
	IntType
	FloatType
	StringType // Text cell
	AreaType   // Range of cells (e.g A1:A10). Only 1D allowed
	ErrorType
)

type Value interface {
	Type() ValueType
	String() string
	ToType(string, ValueType) Value
}
