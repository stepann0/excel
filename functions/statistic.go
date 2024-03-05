package functions

import V "github.com/stepann0/excel/value"

func Sum(area []V.Value) V.Value {
	S := 0.0
	for _, v := range area {
		switch it := v.(type) {
		case V.Int:
			S += float64(it)
		case V.Float:
			S += float64(it)
		case V.Boolean:
			S += float64(it.ToType("sum", V.FloatType, true).(V.Float))
		}
	}
	return V.Float(S)
}

func Avg(area []V.Value) V.Value {
	if len(area) == 0 {
		return V.Float(0)
	}
	S := Sum(area).(V.Float)
	return V.Float(float64(S) / float64(len(area)))
}
