package functions

import V "github.com/stepann0/excel/value"

func True(_ []V.Value) V.Value {
	return V.Boolean{true}
}

func False(_ []V.Value) V.Value {
	return V.Boolean{false}
}

func Not(args []V.Value) V.Value {
	a := args[0].(V.Boolean)
	return V.Boolean{!a.Val}
}

func And(args []V.Value) V.Value {
	a, b := args[0].(V.Boolean), args[1].(V.Boolean)
	return V.Boolean{a.Val && b.Val}
}

func Or(args []V.Value) V.Value {
	a, b := args[0].(V.Boolean), args[1].(V.Boolean)
	return V.Boolean{a.Val || b.Val}
}

func Xor(args []V.Value) V.Value {
	a, b := args[0].(V.Boolean), args[1].(V.Boolean)
	return V.Boolean{(a.Val || b.Val) && !(a.Val && b.Val)}
}
