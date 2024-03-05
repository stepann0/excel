package value

import "fmt"

type Area []Value

func (a Area) Type() ValueType { return AreaType }

func (a Area) ToType(fn string, toT ValueType, abort bool) Value {
	if toT == AreaType {
		return a
	}
	if abort {
		TypeError()
	}
	return nil
}

func (a Area) String() string {
	return fmt.Sprintf("area: %v", a)
}
