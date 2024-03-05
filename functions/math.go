package functions

import (
	"math"
	"math/rand"

	V "github.com/stepann0/excel/value"
)

func OneArgReturnFloat(fn func(float64) float64, a V.Value) V.Float {
	number := a.ToType("", V.FloatType, true).(V.Float)
	return V.Float(fn(float64(number)))
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
	return V.Float(rand.Float64())
}
