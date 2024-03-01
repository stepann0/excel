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
		return V.Error{Msg: NotImplementedError}
	}

	a, ok_a := op.left.Eval().(V.Number[float64])
	b, ok_b := op.right.Eval().(V.Number[float64])
	if !(ok_a || ok_b) {
		return V.Error{Msg: TypeError}
	}

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

func (op *UnOperator) Eval() V.Value {
	if op.token.is("-") {
		a, ok := op.right.Eval().(V.Number[float64])
		if !ok {
			return V.Error{Msg: TypeError}
		}
		return V.Number[float64]{Val: -a.Val}
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
		return V.Error{NotImplementedError}
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
