package formula

import (
	"fmt"
	"strconv"
)

type Parser struct {
	lex       *Lexer
	curTok    Token
	dataTable *DataTable
}

func NewParser(expr string, table *DataTable) *Parser {
	p := &Parser{
		lex:       NewLexer(expr),
		dataTable: table,
	}
	p.eat()
	return p
}

func (p *Parser) Parse() (result Node) {
	defer func() {
		if r := recover(); r != nil {
			result = &ParseErrorNode{fmt.Errorf("%v", r)}
		}
	}()
	result = p.mainFormula()
	return result
}

// Read next token and return previous
func (p *Parser) eat() Token {
	prev_tok := p.curTok
	p.curTok = p.lex.NextToken()
	return prev_tok
}

func (p *Parser) expect(tok, msg string) {
	if p.curTok.is(tok) {
		p.eat()
		return
	}
	panic(fmt.Errorf("invalid syntax: %s", msg))
}

func (p *Parser) mainFormula() Node {
	result := p.expr()
	p.expect("eof", fmt.Sprintf("unexpected %s", p.curTok))
	return result
}

func (p *Parser) expr() Node {
	t1 := p.term()
	for p.curTok.oneOf("+", "-") {
		t1 = &BiOperator{p.eat(), t1, p.term()}
	}
	return t1
}

func (p *Parser) term() Node {
	t1 := p.factor()
	for p.curTok.oneOf("*", "/") {
		t1 = &BiOperator{p.eat(), t1, p.factor()}
	}
	return t1
}

func (p *Parser) factor() Node {
	if p.curTok.is("-") {
		return &UnOperator{p.eat(), p.factor()}
	} else if p.curTok.is("+") {
		p.eat()
		return p.factor()
	}
	return p.base()
}

func (p *Parser) base() Node {
	if p.curTok.is("num") {
		num, err := strconv.ParseFloat(p.curTok.literal, 64)
		if err != nil {
			panic(err)
		}
		return &NumberLit[float64]{p.eat(), num}
	} else if p.curTok.is("bool") {
		b := BoolLit(p.eat().literal == "TRUE")
		return &b
	} else if p.curTok.is("(") {
		p.eat()
		expr := p.expr()
		p.expect(")", "missing ')'")
		return expr
	} else if p.curTok.is("func") {
		func_tok := p.eat()
		p.expect("(", "missing '('")
		args := p.argList()
		p.expect(")", "missing ')'")
		return &FuncCall{func_tok, args}
	} else if p.curTok.is("ref") {
		ref1 := &ReferenceLit{p.eat(), p.dataTable}
		if p.curTok.is(":") {
			colon := p.eat()
			ref2 := &ReferenceLit{p.curTok, p.dataTable}
			p.expect("ref", "missing reference after ':'")
			return &BiOperator{colon, ref1, ref2}
		}
		return ref1
	}
	panic("expected number, function, reference, bool or '('")
}

func (p *Parser) argList() []Node {
	if p.curTok.is(")") {
		return []Node{}
	}
	arg_list := []Node{p.expr()}
	for p.curTok.is(",") {
		p.eat()
		arg_list = append(arg_list, p.expr())
	}
	return arg_list
}
