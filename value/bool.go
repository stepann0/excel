package value

const (
	TRUE_LITERAL  = "TRUE"
	FALSE_LITERAL = "FALSE"
)

type Boolean struct {
	Val bool
}

// in Excel TRUE and FALSE are just 1 and 0, so it is IntType
func (b Boolean) Type() ValueType { return IntType }

func (b Boolean) ToType(fn string, toT ValueType) Value {
	switch toT {
	case BooleanType:
		return b
	case IntType:
		if b.Val {
			return Int{1}
		}
		return Int{0}
	case FloatType:
		if b.Val {
			return Float{1.0}
		}
		return Float{0.0}
	}
	TypeError()
	return nil
}

func (b Boolean) String() string {
	if b.Val {
		return TRUE_LITERAL
	}
	return FALSE_LITERAL
}
