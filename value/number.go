package value

import "fmt"

type Number[T float64 | int] struct {
	Val T
}

func (n Number[T]) Type() ValueType { return NumberType }

func (n Number[T]) String() string {
	return fmt.Sprint(n.Val)
}

func FromFloat(n float64) Number[float64] {
	return Number[float64]{n}
}
