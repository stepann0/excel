package value

import "fmt"

type Int int64

func (n Int) Type() ValueType { return IntType }

func (n Int) String() string {
	return fmt.Sprint(n)
}

func (n Int) ToType(fn string, toT ValueType, abort bool) Value {
	switch toT {
	case IntType:
		return n
	case FloatType:
		return Float(n)
	case BooleanType:
		return Boolean(n != 0)
	}
	if abort {
		TypeError()
	}
	return nil
}

type Float float64

func (n Float) Type() ValueType { return FloatType }

func (n Float) String() string {
	return fmt.Sprint(n)
}

func (n Float) ToType(fn string, toT ValueType, abort bool) Value {
	switch toT {
	case FloatType:
		return n
	case BooleanType:
		return Boolean(n != 0)
	}
	if abort {
		TypeError()
	}
	return nil
}
