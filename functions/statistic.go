package functions

import V "github.com/stepann0/excel/value"

func Sum(area []V.Value) V.Value {
	S := 0.0
	for _, v := range area {
		switch it := v.(type) {
		case V.Int:
			S += float64(it.Val)
		case V.Float:
			S += it.Val
		case V.Boolean:
			S += it.ToType("sum", V.FloatType).(V.Float).Val
		}
	}
	return V.Float{Val: S}
}

func Avg(area []V.Value) V.Value {
	if len(area) == 0 {
		return V.Float{0}
	}
	S := Sum(area).(V.Float)
	return V.Float{Val: S.Val / float64(len(area))}
}
