package formula

import (
	"fmt"

	"github.com/stepann0/excel/functions"
	V "github.com/stepann0/excel/value"
)

var (
	ParseError          = fmt.Errorf("parse error")
	TypeError           = fmt.Errorf("type error")
	NotImplementedError = fmt.Errorf("not implemented")
)

func (n *NumberLit[NumType]) Eval() V.Value {
	if fmt.Sprintf("%T", n.val) == "int" {
		return V.Int(n.val)
	}
	return V.Float(n.val)
}
func (s *StringLit) Eval() V.Value { return V.String(s.val) }
func (b *BoolLit) Eval() V.Value   { return V.Boolean(*b) }

func promoteType(a, b V.ValueType) V.ValueType {
	if a > b {
		return a
	}
	return b
}

func (op *BiOperator) Eval() V.Value {
	if op.token.is(":") {
		// ref1, ref2 := op.left.tokenLiteral(), op.right.tokenLiteral()
		V.NotImplementedError()
	}
	n1 := op.left.Eval()
	n2 := op.right.Eval()

	higherT := max(n1.Type(), n2.Type())
	a := n1.ToType(op.tokenLiteral(), higherT, true)
	b := n2.ToType(op.tokenLiteral(), higherT, true)
	return applyBinary(op.tokenLiteral(), higherT, a, b)
}

type bifunc func(V.Value, V.Value) V.Value

var Appliers = map[string]map[V.ValueType]bifunc{
	"+": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Int(a.(V.Int) + b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Float(a.(V.Float) + b.(V.Float))
		},
	},
	"-": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Int(a.(V.Int) - b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Float(a.(V.Float) - b.(V.Float))
		},
	},
	"*": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Int(a.(V.Int) * b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Float(a.(V.Float) * b.(V.Float))
		},
	},
	"/": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Int(a.(V.Int) / b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Float(a.(V.Float) / b.(V.Float))
		},
	},
	"<": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) < b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) < b.(V.Float))
		},
	},
	">": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) > b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) > b.(V.Float))
		},
	},
	"<=": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) <= b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) <= b.(V.Float))
		},
	},
	">=": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) >= b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) >= b.(V.Float))
		},
	},
	"=": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) == b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) == b.(V.Float))
		},
	},
	"<>": {
		V.IntType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Int) != b.(V.Int))
		},
		V.FloatType: func(a, b V.Value) V.Value {
			return V.Boolean(a.(V.Float) != b.(V.Float))
		},
	},
}

func applyBinary(op string, T V.ValueType, a, b V.Value) V.Value {
	fn := Appliers[op][T]
	if fn == nil {
		V.NotImplementedError()
	}
	return fn(a, b)
}

func (op *UnOperator) Eval() V.Value {
	if op.token.is("-") {
		a := op.right.Eval()
		switch a := a.(type) {
		case V.Float:
			return V.Float(-a)
		case V.Int:
			return V.Int(-a)
		default:
			V.TypeError()
		}
	}
	return nil
}

func (fc *FuncCall) Eval() V.Value {
	args := []V.Value{}
	for _, a := range fc.args {
		args = append(args, a.Eval())
	}
	callie, ok := functions.FuncList[fc.tokenLiteral()]
	if !ok {
		V.NotImplementedError()
	}
	return callie.Call(args)
}

func (ref *ReferenceLit) Eval() V.Value {
	// return ref.table.At(ref.tokenLiteral())
	return nil
}

func (e *ParseErrorNode) Eval() V.Value {
	return V.Error{Msg: e.body}
}
