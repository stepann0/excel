package value

import "fmt"

type Area struct {
	Val []Value
}

func (a Area) Type() ValueType { return AreaType }

func (a Area) ToType(fn string, which ValueType) Value {
	if which == AreaType {
		return a
	}
	TypeError()
	return nil
}

func (a Area) String() string {
	return fmt.Sprintf("area: %v", a.Val)
}
