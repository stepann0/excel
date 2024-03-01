package functions

import (
	"math"
	"math/rand"

	V "github.com/stepann0/excel/value"
)

func OneArgReturnFloat(fn func(float64) float64, a V.Value) V.Number[float64] {
	var number float64
	switch a := a.(type) {
	case V.Number[float64]:
		number = a.Val
	case V.Number[int]:
		number = float64(a.Val)
	default:
		println("math functions")
		V.TypeError()
	}
	return V.Number[float64]{Val: fn(number)}
}

func Sin(a []V.Value) V.Value {
	return OneArgReturnFloat(math.Sin, a[0])
}

func Cos(a []V.Value) V.Value {
	return OneArgReturnFloat(math.Cos, a[0])
}

func Abs(a []V.Value) V.Value {
	return OneArgReturnFloat(math.Abs, a[0])
}

func Exp(a []V.Value) V.Value {
	return OneArgReturnFloat(math.Exp, a[0])
}

func Rand(_ []V.Value) V.Value {
	return V.FromFloat(rand.Float64())
}
