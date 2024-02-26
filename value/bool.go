package value

const (
	TRUE_LITERAL  = "TRUE"
	FALSE_LITERAL = "FALSE"
)

type Boolean struct {
	Val bool
}

func (b Boolean) Type() ValueType { return BooleanType }

func (b *Boolean) ToInt() Number[int] {
	if b.Val {
		return Number[int]{1}
	}
	return Number[int]{0}
}

func (b *Boolean) ToFloat() Number[float64] {
	if b.Val {
		return Number[float64]{1.0}
	}
	return Number[float64]{0.0}
}

func (b Boolean) String() string {
	if b.Val {
		return TRUE_LITERAL
	}
	return FALSE_LITERAL
}
