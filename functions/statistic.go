package functions

import V "github.com/stepann0/excel/value"

func Sum(area []V.Value) V.Value {
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

func Avg(area []V.Value) V.Value {
	if len(area) == 0 {
		return V.FromFloat(0)
	}
	S := Sum(area).(V.Number[float64])
	return V.Number[float64]{Val: S.Val / float64(len(area))}
}
