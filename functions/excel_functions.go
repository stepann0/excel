package functions

import (
	"math"

	V "github.com/stepann0/excel/value"
)

func Sin(a V.Value) V.Number[float64] {
	res := math.Sin(a.(V.Number[float64]).Val)
	return V.Number[float64]{Val: res}
}

func Sum(area []V.Value) V.Number[float64] {
	S := 0.0
	for _, v := range area {
		switch it := v.(type) {
		case V.Number[int]:
			S += float64(it.Val)
		case V.Number[float64]:
			S += it.Val
		case V.Boolean:
			S += it.ToFloat().Val
		}
	}
	return V.Number[float64]{Val: S}
}
