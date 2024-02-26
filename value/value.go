package value

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
	String() string
}
