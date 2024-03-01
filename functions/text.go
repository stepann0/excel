package functions

import V "github.com/stepann0/excel/value"

func Concat(text []V.Value) V.String {
	res := ""
	for _, t := range text {
		res += t.String()
	}
	return V.String{Val: res}
}
