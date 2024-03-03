package value

import "fmt"

type Int struct {
	Val int64
}

func (n Int) Type() ValueType { return IntType }

func (n Int) String() string {
	return fmt.Sprint(n.Val)
}

func (n Int) ToFloat() Float {
	return Float{float64(n.Val)}
}

func (n Int) ToType(fn string, toT ValueType) Value {
	switch toT {
	case IntType:
		return n
	case FloatType:
		fl := n.ToFloat()
		return fl
	case BooleanType:
		if n.Val != 0 {
			return Boolean{true}
		}
		return Boolean{false}
	}
	TypeError()
	return nil
}

type Float struct {
	Val float64
}

func (n Float) Type() ValueType { return FloatType }

func (n Float) String() string {
	return fmt.Sprint(n.Val)
}

func (n Float) ToType(fn string, toT ValueType) Value {
	switch toT {
	case FloatType:
		return n
	case BooleanType:
		if n.Val != 0 {
			return Boolean{true}
		}
		return Boolean{false}
	}
	TypeError()
	return nil
}
