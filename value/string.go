package value

type String struct {
	Val string
}

func (s String) Type() ValueType { return StringType }

func (s String) ToType(fn string, toT ValueType) Value {
	if toT == StringType {
		return s
	}
	TypeError()
	return nil
}

func (s String) String() string {
	return s.Val
}
