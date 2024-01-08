package data

import (
	"fmt"
	"math"
)

func oneArgFunc(name string, f ...float64) func(float64) float64 {
	if len(f) != 1 {
		panic(fmt.Errorf("function %s: expected only one argument, got %d", name, len(f)))
	}
	switch name {
	case "sin":
		return math.Sin
	case "cos":
		return math.Cos
	case "tan":
		return math.Tan
	}
	return func(f float64) float64 { return f }
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
	case "avg":
		return func(f ...float64) float64 {
			sum := 0.0
			for _, i := range f {
				sum += i
			}
			return sum / float64(len(f))
		}
	case "sin", "cos", "tan":
		return func(f ...float64) float64 {
			return oneArgFunc(name, f...)(f[0])
		}
	case "pow":
		return func(f ...float64) float64 {
			if len(f) != 2 {
				panic("expected only two arguments")
			}
			return math.Pow(f[0], f[1])
		}
	case "max":
		return func(f ...float64) float64 {
			res := math.Inf(-1)
			for _, i := range f {
				if i >= res {
					res = i
				}
			}
			return res
		}
	case "min":
		return func(f ...float64) float64 {
			res := math.Inf(1)
			for _, i := range f {
				if i <= res {
					res = i
				}
			}
			return res
		}
	}
	return nil
}

func floatRange(rng []any) []float64 {
	res := make([]float64, len(rng))
	for i := range rng {
		f, ok := rng[i].(float64)
		if !ok {
			f = 0
		}
		res[i] = f
	}
	return res
}
