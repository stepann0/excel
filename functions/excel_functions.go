package functions

import (
	V "github.com/stepann0/excel/value"
)

var FuncList map[string]FunctionProto

type ExcelFunction func([]V.Value) V.Value
type ArgCFunc func([]V.Value) bool // Example: atLeast(1), lessThan(3)

type FunctionProto struct {
	fn       ExcelFunction
	argCheck ArgCFunc
}

func (f *FunctionProto) Call(args []V.Value) V.Value {
	if !f.argCheck(args) {
		V.ArgCountError()
	}
	return f.fn(args)
}

func init() {
	if FuncList != nil {
		return
	}

	FuncList = map[string]FunctionProto{
		// Math
		"sin": {
			fn:       Sin,
			argCheck: exactly(1),
		},
		"cos": {
			fn:       Cos,
			argCheck: exactly(1),
		},
		"abs": {
			fn:       Abs,
			argCheck: exactly(1),
		},
		"exp": {
			fn:       Exp,
			argCheck: exactly(1),
		},
		"rand": {
			fn:       Rand,
			argCheck: lessThan(2),
		},
		// Statistics
		"sum": {
			fn:       Sum,
			argCheck: atLeast(0),
		},
		"avg": {
			fn:       Avg,
			argCheck: atLeast(0),
		},
	}
}

func atLeast(n int) ArgCFunc {
	return func(args []V.Value) bool {
		return len(args) >= n
	}
}

func exactly(n int) ArgCFunc {
	return func(v []V.Value) bool {
		return len(v) == n
	}
}

func lessThan(n int) ArgCFunc {
	return func(args []V.Value) bool {
		return len(args) < n
	}
}
