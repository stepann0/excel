package formula

import (
	"fmt"

	"github.com/stepann0/tercel/data"
)

// <expr>    = <term> {("+" | "-") <term>}
// <term>    = <factor> {("*" | "/" | "%") <factor>}
// <factor>  = {("-" | "+")} <base>
// <base>    = <constant> | <ref> | <ref>:<ref> |
// <function> "(" <expr> {"," <expr>} ")" | "(" <expr> ")"
type Parser struct {
	lex *Lexer
	tok Token
	dt  *data.DataTable
}

func NewParser(expr string, table *data.DataTable) *Parser {
	p := &Parser{
		lex: NewLexer(expr),
	}
	p.tok = p.lex.NextToken()
	return p
}

func (p *Parser) Eval() float64 {
	return p.expr()
}

func (p *Parser) eat() {
	p.tok = p.lex.NextToken()
}

func (p *Parser) expect(tok, msg string) {
	if p.tok.is(tok) {
		p.eat()
		return
	}
	panic(fmt.Errorf("invalid syntax: %s", msg))
}

func (p *Parser) expr() float64 {
	t1 := p.term()
	for p.tok.oneOf("+", "-") {
		if p.tok.is("+") {
			p.eat()
			t1 += p.term()
		} else if p.tok.is("-") {
			p.eat()
			t1 -= p.term()
		}
	}
	return t1
}

func (p *Parser) term() float64 {
	t1 := p.factor()
	for p.tok.oneOf("*", "/") {
		if p.tok.is("*") {
			p.eat()
			t1 *= p.factor()
		} else if p.tok.is("/") {
			p.eat()
			t1 /= p.factor()
		}
	}
	return t1
}

func (p *Parser) factor() float64 {
	if p.tok.is("-") {
		p.eat()
		return -p.base()
	} else if p.tok.is("+") {
		p.eat()
		return p.base()
	}
	return p.base()
}

func (p *Parser) base() float64 {
	if p.tok.is("num") {
		num := p.tok.value.(float64)
		p.eat()
		return num
	} else if p.tok.is("(") {
		p.eat()
		res := p.expr()
		p.expect(")", "missing ')'")
		return res
	} else if p.tok.is("func") {
		f := getFunc(p.tok.value.(string))
		p.eat()
		p.expect("(", "missing '('")
		args := p.argList()
		res := f(args...)
		p.expect(")", "missing ')'")
		return res
	} else if p.tok.is("ref") {
		ref1 := p.tok
		p.eat()
		return p.dt.AtRef(ref1.value.(string)).Data().(float64)
	}
	return 0
}

func (p *Parser) argList() []float64 {
	if p.tok.is("ref") {
		ref1 := p.tok
		p.eat()
		if p.tok.is(":") {
			p.eat()
			p.expect("ref", "missing reference after ':'")
			ref2 := p.tok
			ref2, ref1 = ref2, ref1
			return []float64{111111111}
		} else {

		}
	}
	list := []float64{p.expr()}
	for p.tok.is(",") {
		p.eat()
		list = append(list, p.expr())
	}
	return list
}