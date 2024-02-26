package formula

import "strings"

const (
	OPEN  = "{"
	CLOSE = "}"
)

func shift(level int) string {
	return strings.Repeat("  ", level)
}

func node_header(node Node) string {
	switch node.(type) {
	case *FuncCall:
		return "fn[" + node.tokenLiteral() + "]" + OPEN
	case *BiOperator, *UnOperator:
		return "op[" + node.tokenLiteral() + "]" + OPEN
	}
	return node.tokenLiteral()
}

func (n *NumberLit[ValType]) inspect(level int) string { return shift(level) + n.tokenLiteral() }

func (s *StringLit) inspect(level int) string { return shift(level) + "str" + s.tokenLiteral() }

func (f *FuncCall) inspect(level int) string {
	repr := shift(level) + node_header(f)
	repr += "\n"
	for _, arg := range f.args {
		repr += arg.inspect(level+1) + "\n"
	}
	repr += shift(level) + CLOSE
	return repr
}

func (b *BoolLit) inspect(level int) string { return shift(level) + b.tokenLiteral() }

func (o *BiOperator) inspect(level int) string {
	repr := shift(level) + node_header(o)
	repr += "\n"
	repr += o.left.inspect(level+1) + "\n"
	repr += o.right.inspect(level+1) + "\n"
	repr += shift(level) + CLOSE
	return repr
}

func (o *UnOperator) inspect(level int) string {
	repr := shift(level) + node_header(o)
	repr += "\n"
	repr += o.right.inspect(level+1) + "\n"
	repr += shift(level) + CLOSE
	return repr
}

func (r *ReferenceLit) inspect(level int) string {
	return shift(level) + "&" + r.tokenLiteral()
}

func (e *ParseErrorNode) inspect(level int) string {
	return shift(level) + "ERROR(" + e.body.Error() + ")"
}
