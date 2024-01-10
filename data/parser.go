package data

import (
	"fmt"
)

const (
	ResError int = iota
	ResNil
	ResNumber
	ResRange
)

type Result struct {
	typ int
	val any
}

func result(val any) Result {
	switch val.(type) {
	case nil:
		return Result{}
	case []any:
		return Result{ResRange, val}
	case int, int32, int64, float32:
		return Result{ResNumber, float64(val.(int))}
	case float64:
		return Result{ResNumber, val}
	}
	return Result{ResError, fmt.Errorf("unknown type of value: %v, type %T", val, val)}
}

func (r Result) String() string {
	var t string
	switch r.typ {
	case ResError:
		t = "err"
	case ResNil:
		t = "nil"
	case ResNumber:
		t = "num"
	case ResRange:
		t = "rng"
	}
	return fmt.Sprintf("{%s %v}", t, r.val)
}

func (r *Result) Type() int {
	return r.typ
}

func (r *Result) Value() any {
	return r.val
}

func (a *Result) add(b Result) {
	if a.typ == ResNumber && b.typ == ResNumber {
		n1, n2 := a.val.(float64), b.val.(float64)
		a.val = n1 + n2
		return
	}
	*a = result(fmt.Errorf("can't add %v and %v", a, b))
}

func (a *Result) sub(b Result) {
	if a.typ == ResNumber && b.typ == ResNumber {
		n1, n2 := a.val.(float64), b.val.(float64)
		a.val = n1 - n2
		return
	}
	*a = result(fmt.Errorf("can't sub %v and %v", a, b))
}

func (a *Result) mul(b Result) {
	if a.typ == ResNumber && b.typ == ResNumber {
		n1, n2 := a.val.(float64), b.val.(float64)
		a.val = n1 * n2
		return
	}
	*a = result(fmt.Errorf("can't multiply %v and %v", a, b))
}

func (a *Result) div(b Result) {
	if a.typ == ResNumber && b.typ == ResNumber {
		n1, n2 := a.val.(float64), b.val.(float64)
		a.val = n1 / n2
		return
	}
	*a = result(fmt.Errorf("can't divide %v and %v", a, b))
}

func (a *Result) neg() {
	if a.typ == ResNumber {
		a.val = -a.val.(float64)
		return
	}
	*a = result(fmt.Errorf("can't negate %v", a))
}

// <expr>    = <term> {("+" | "-") <term>}
// <term>    = <factor> {("*" | "/" | "%") <factor>}
// <factor>  = ("-" | "+") <factor> | <base>
// <base>    = <constant> | <ref> | <ref>:<ref> |
// <function> "(" <expr> {"," <expr>} ")" | "(" <expr> ")"
type Parser struct {
	lex *Lexer
	tok Token
	dt  *DataTable
}

func NewParser(expr string, table *DataTable) *Parser {
	p := &Parser{
		lex: NewLexer(expr),
		dt:  table,
	}
	p.tok = p.lex.NextToken()
	return p
}

func (p *Parser) Eval() Result {
	return p.expr()
}

func (p *Parser) eat() {
	p.tok = p.lex.NextToken()
}

func (p *Parser) expect(tok, msg string) error {
	if p.tok.is(tok) {
		p.eat()
		return nil
	}
	return fmt.Errorf("invalid syntax: %s", msg)
}

func (p *Parser) expr() Result {
	t1 := p.term()
	for p.tok.oneOf("+", "-") {
		if p.tok.is("+") {
			p.eat()
			t1.add(p.term())
		} else if p.tok.is("-") {
			p.eat()
			t1.sub(p.term())
		}
	}
	return t1
}

func (p *Parser) term() Result {
	t1 := p.factor()
	for p.tok.oneOf("*", "/") {
		if p.tok.is("*") {
			p.eat()
			t1.mul(p.factor())
		} else if p.tok.is("/") {
			p.eat()
			t1.div(p.factor())
		}
	}
	return t1
}

func (p *Parser) factor() Result {
	if p.tok.is("-") {
		p.eat()
		f := p.factor()
		f.neg()
		return f
	} else if p.tok.is("+") {
		p.eat()
		return p.factor()
	}
	return p.base()
}

func (p *Parser) base() Result {
	if p.tok.is("num") {
		num := p.tok.value.(float64)
		p.eat()
		return result(num)
	} else if p.tok.is("(") {
		p.eat()
		res := p.expr()
		p.expect(")", "missing ')'")
		return res
	} else if p.tok.is("func") {
		func_name := p.tok.value.(string)
		f := getFunc(func_name)
		if f == nil {
			return result(fmt.Errorf("function '%s' not implemented", func_name))
		}
		p.eat()
		p.expect("(", "missing '('")
		args := p.argList()
		res := f(args...)
		p.expect(")", "missing ')'")
		return res
	} else if p.tok.is("ref") {
		ref1 := p.tok
		p.eat()
		if p.tok.is(":") {
			p.eat()
			ref2 := p.tok
			if err := p.expect("ref", "missing reference after ':'"); err != nil {
				return result(err)
			}
			rng := p.dt.GetRange(ref1.value.(string), ref2.value.(string))
			return result(rng)
		}
		cell := p.dt.AtRef(ref1.value.(string))
		if cell == nil {
			return result(fmt.Errorf("no such cell"))
		}

		cell_data := cell.Data()
		if cell_data == nil {
			return result(nil)
		}
		if num, ok := cell_data.(float64); ok {
			return result(num)
		}
		return result(fmt.Errorf("formula can accept only numeric data"))
	}
	return result(nil)
}

func (p *Parser) argList() []Result {
	list := []Result{p.expr()}
	for p.tok.is(",") {
		p.eat()
		list = append(list, p.expr())
	}
	return list
}
