package value

type String struct {
	Val string
}

func (s String) Type() ValueType { return StringType }

func (s String) String() string {
	return s.Val
}
