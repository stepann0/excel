package data

type Node interface {
	tokenLiteral() string
	inspect(int) string
	Eval() Value
}

// --- Numbers ---
type NumberLit[NumType float64 | int] struct {
	token Token // number literal
	val   NumType
}

func (n *NumberLit[NumType]) tokenLiteral() string { return n.token.literal }

// --- String ---
type StringLit struct {
	token Token // number literal
	val   string
}

func (s *StringLit) tokenLiteral() string { return s.token.literal }

// --- Function ---
type FuncCall struct {
	token Token // func name
	args  []Node
}

func (f *FuncCall) tokenLiteral() string { return f.token.literal }

// --- BoolLit ---
type BoolLit bool

func (b *BoolLit) tokenLiteral() string {
	if *b {
		return "TRUE"
	}
	return "FALSE"
}

// --- Operators ---
type BiOperator struct {
	token       Token
	left, right Node
}

func (o *BiOperator) tokenLiteral() string { return o.token.literal }

type UnOperator struct {
	token Token
	right Node
}

func (o *UnOperator) tokenLiteral() string { return o.token.literal }

// --- Reference ---
type ReferenceLit struct {
	token Token
	table *DataTable
}

func (r *ReferenceLit) tokenLiteral() string { return r.token.literal }

// --- Parsing errors ---
type ParseErrorNode struct {
	body error
}

func (e *ParseErrorNode) tokenLiteral() string { return "no_literal" }
