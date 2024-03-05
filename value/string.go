package value

type String string

func (s String) Type() ValueType { return StringType }

func (s String) ToType(fn string, toT ValueType, abort bool) Value {
	switch toT {
	case StringType:
		return s
	case BooleanType:
		return Boolean(len(s) > 0)
	}
	if abort {
		TypeError()
	}
	return nil
}

func (s String) String() string {
	return string(s)
}
