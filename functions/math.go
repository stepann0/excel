package functions

import (
	"math"
	"math/rand"

	V "github.com/stepann0/excel/value"
)

func OneArgReturnFloat(fn func(float64) float64, a V.Value) V.Float {
	var number float64
	switch a := a.(type) {
	case V.Float:
		number = a.Val
	case V.Int:
		number = float64(a.Val)
	default:
		println("math functions")
		V.TypeError()
	}
	return V.Float{Val: fn(number)}
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
	return V.Float{rand.Float64()}
}
