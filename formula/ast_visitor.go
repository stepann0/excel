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

func (n *NumberLit[NumType]) Eval() V.Value { return V.Number[NumType]{Val: n.val} }
func (s *StringLit) Eval() V.Value          { return V.String{Val: s.val} }
func (b *BoolLit) Eval() V.Value            { return V.Boolean{Val: bool(*b)} }

func (op *BiOperator) Eval() V.Value {
	if op.token.is(":") {
		// ref1, ref2 := op.left.tokenLiteral(), op.right.tokenLiteral()
		V.NotImplementedError()
	}

	n1 := op.left.Eval()
	n2 := op.right.Eval()

	isInt := func(v V.Value) bool {
		_, ok := v.(V.Number[int])
		return ok
	}

	// Add, sub and mul two integers
	if isInt(n1) && isInt(n2) && op.token.oneOf("+", "-", "*") {
		return IntMath(op.token, n1, n2)
	}
	a, b := V.ToFloat(n1), V.ToFloat(n2)
	switch {
	case op.token.is("+"):
		return V.Number[float64]{Val: a.Val + b.Val}
	case op.token.is("-"):
		return V.Number[float64]{Val: a.Val - b.Val}
	case op.token.is("*"):
		return V.Number[float64]{Val: a.Val * b.Val}
	case op.token.is("/"):
		return V.Number[float64]{Val: a.Val / b.Val}
	}
	return nil
}

func IntMath(op Token, left, right V.Value) V.Number[int] {
	a, b := left.(V.Number[int]), right.(V.Number[int])
	switch {
	case op.is("+"):
		return V.Number[int]{Val: a.Val + b.Val}
	case op.is("-"):
		return V.Number[int]{Val: a.Val - b.Val}
	case op.is("*"):
		return V.Number[int]{Val: a.Val * b.Val}
	case op.is("/"):
		V.TypeError()
	}
	return V.Number[int]{}
}

func (op *UnOperator) Eval() V.Value {
	if op.token.is("-") {
		a := op.right.Eval()
		switch a := a.(type) {
		case V.Number[float64]:
			return V.Number[float64]{Val: -a.Val}
		case V.Number[int]:
			return V.Number[int]{Val: -a.Val}
		}
		V.TypeError()
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
