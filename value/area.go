package value

import "fmt"

type Area struct {
	Val []Value
}

func (a Area) Type() ValueType { return AreaType }

func (a Area) String() string {
	return fmt.Sprintf("area: %v", a.Val)
}
