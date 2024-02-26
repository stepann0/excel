package data

import "fmt"

var (
	ParseError          = fmt.Errorf("parse error")
	TypeError           = fmt.Errorf("type error")
	NotImplementedError = fmt.Errorf("not implemented")
)

func (n *NumberLit[NumType]) Eval() Value { return Number[NumType]{n.val} }
func (s *StringLit) Eval() Value          { return String{s.val} }
func (b *BoolLit) Eval() Value            { return Boolean{bool(*b)} }

func (op *BiOperator) Eval() Value {
	if op.token.is(":") {
		// ref1, ref2 := op.left.tokenLiteral(), op.right.tokenLiteral()
		return Error{NotImplementedError}
	}

	a, ok_a := op.left.Eval().(Number[float64])
	b, ok_b := op.right.Eval().(Number[float64])
	if !(ok_a || ok_b) {
		return Error{TypeError}
	}

	switch {
	case op.token.is("+"):
		return Number[float64]{a.Val + b.Val}
	case op.token.is("-"):
		return Number[float64]{a.Val - b.Val}
	case op.token.is("*"):
		return Number[float64]{a.Val * b.Val}
	case op.token.is("/"):
		return Number[float64]{a.Val / b.Val}
	}
	return nil
}

func (op *UnOperator) Eval() Value {
	if op.token.is("-") {
		a, ok := op.right.Eval().(Number[float64])
		if !ok {
			return Error{TypeError}
		}
		return Number[float64]{-a.Val}
	}
	return nil
}

func (fc *FuncCall) Eval() Value {
	return nil
}

func (ref *ReferenceLit) Eval() Value {
	// return ref.table.At(ref.tokenLiteral())
	return nil
}

func (e *ParseErrorNode) Eval() Value {
	return Error{e.body}
}
