package formula

import (
	"fmt"
	"strconv"

	V "github.com/stepann0/excel/value"
)

type Parser struct {
	lex       *Lexer
	curTok    Token
	dataTable *V.DataTable
}

func NewParser(expr string, table *V.DataTable) *Parser {
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
	result := p.equality()
	p.expect("eof", fmt.Sprintf("unexpected %s", p.curTok))
	return result
}

func (p *Parser) equality() Node {
	t1 := p.comparison()
	for p.curTok.oneOf("=", "<>") {
		t1 = &BiOperator{p.eat(), t1, p.comparison()}
	}
	return t1
}

func (p *Parser) comparison() Node {
	t1 := p.sum()
	for p.curTok.oneOf(">", "<", ">=", "<=") {
		t1 = &BiOperator{p.eat(), t1, p.sum()}
	}
	return t1
}

func (p *Parser) sum() Node {
	t1 := p.prod()
	for p.curTok.oneOf("+", "-") {
		t1 = &BiOperator{p.eat(), t1, p.prod()}
	}
	return t1
}

func (p *Parser) prod() Node {
	t1 := p.unary()
	for p.curTok.oneOf("*", "/") {
		t1 = &BiOperator{p.eat(), t1, p.unary()}
	}
	return t1
}

func (p *Parser) unary() Node {
	if p.curTok.is("-") {
		return &UnOperator{p.eat(), p.unary()}
	} else if p.curTok.is("+") {
		p.eat()
		return p.unary()
	}
	return p.atom()
}

func (p *Parser) atom() Node {
	if p.curTok.is("num") {
		if int_num, err := strconv.ParseInt(p.curTok.literal, 10, 32); err == nil {
			return &NumberLit[int]{p.eat(), int(int_num)}
		}
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
		expr := p.sum()
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
	arg_list := []Node{p.sum()}
	for p.curTok.is(",") {
		p.eat()
		arg_list = append(arg_list, p.sum())
	}
	return arg_list
}
