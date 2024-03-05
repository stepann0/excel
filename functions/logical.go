package functions

import V "github.com/stepann0/excel/value"

func True(_ []V.Value) V.Value {
	return V.Boolean(true)
}

func False(_ []V.Value) V.Value {
	return V.Boolean(false)
}

func Not(args []V.Value) V.Value {
	a := args[0].(V.Boolean)
	return V.Boolean(!a)
}

func logicalHelper(op string, args []V.Value) V.Boolean {
	A := false
	if a := args[0].ToType(op, V.BooleanType, false); a != nil {
		A = bool(a.(V.Boolean))
	}
	for _, v := range args[1:] {
		if b := v.ToType(op, V.BooleanType, false); b != nil {
			B := bool(b.(V.Boolean))
			switch op {
			case "and":
				A = A && B
			case "or":
				if A {
					return V.Boolean(A)
				}
				A = A || B
			case "xor":
				A = (A || B) && !(A && B)
			}
		}
	}
	return V.Boolean(A)
}

func And(args []V.Value) V.Value {
	return logicalHelper("and", args)
}

func Or(args []V.Value) V.Value {
	return logicalHelper("or", args)
}

func Xor(args []V.Value) V.Value {
	return logicalHelper("xor", args)
}
