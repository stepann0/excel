package data

import (
	"fmt"
	"math"
)

func oneArgFunc(name string) func(float64) float64 {
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

func getFunc(name string) func(...Result) Result {
	switch name {
	case "sum":
		return func(args ...Result) Result {
			res := result(0.0)
			for _, i := range args {
				if i.typ == ResRange {
					// unpack range
					for _, j := range floatRange(i.val.([]any)) {
						res.add(result(j))
					}
					continue
				}
				res.add(i)
			}
			return res
		}
	case "avg":
		return func(args ...Result) Result {
			sum := result(0.0)
			for _, i := range args {
				if i.typ == ResRange {
					// unpack range
					for _, j := range floatRange(i.val.([]any)) {
						sum.add(result(j))
					}
					continue
				}
				sum.add(i)
			}
			sum.div(result(len(args)))
			return sum
		}
	case "sin", "cos", "tan":
		return func(f ...Result) Result {
			if len(f) != 1 {
				panic(fmt.Errorf("function %s: expected one argument, got %d", name, len(f)))
			}
			if f[0].typ != ResNumber {
				panic(fmt.Errorf("function %s: argument must be a number", name))
			}
			arg := f[0].val.(float64)
			return result(oneArgFunc(name)(arg))
		}
	case "pow":
		return func(f ...Result) Result {
			if len(f) != 2 {
				panic(fmt.Errorf("function pow: expected two arguments, got %d", len(f)))
			}
			if f[0].typ != ResNumber || f[1].typ != ResNumber {
				panic("function pow: both arguments must be numbers")
			}
			arg1 := f[0].val.(float64)
			arg2 := f[1].val.(float64)
			return result(math.Pow(arg1, arg2))
		}
	case "max":
		return func(f ...Result) Result {
			res := math.Inf(-1)
			for _, i := range f {
				if i.typ != ResNumber {
					panic("function max: all arguments must be a numbers")
				}
				num := i.val.(float64)
				if num >= res {
					res = num
				}
			}
			return result(res)
		}
	case "min":
		return func(f ...Result) Result {
			res := math.Inf(1)
			for _, i := range f {
				if i.typ != ResNumber {
					panic("function min: all arguments must be a numbers")
				}
				num := i.val.(float64)
				if num <= res {
					res = num
				}
			}
			return result(res)
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
