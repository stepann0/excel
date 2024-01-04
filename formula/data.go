package formula

import (
	"math"
	"strconv"
)

type CellType int

const (
	Text CellType = iota
	Number
	Formula
)

func ConvertType(record string) (any, CellType) {
	if record[0] == '=' {
		return record, Formula
	}
	if n, err := strconv.ParseFloat(record, 64); err == nil {
		return n, Number
	}
	return record, Text
}

func getFunc(name string) func(...float64) float64 {
	switch name {
	case "sum":
		return func(f ...float64) float64 {
			res := 0.0
			for _, i := range f {
				res += i
			}
			return res
		}
	case "sin":
		return func(f ...float64) float64 {
			if len(f) != 1 {
				panic("expected only one argument")
			}
			return math.Sin(f[0])
		}
	}
	return nil
}
