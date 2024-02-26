package formula

import (
	"fmt"
	"testing"
)

var inputs []string = []string{
	"-(-(-(-4)))",
	"----4",
	"(0.31415*3+(2/77-1) * 100)",
	"((-14.98+34.241*0.4)/(-(201.2+33.241)*(0.05))-(11)+1852.098)",
	"-(3053450.352463)/-(-123346)*+(0.00053524)",
	"3+2*4", "100", "+++200", "+(+(+(+(10.0))))",
	"(200 / (4 * 0.5)) + 50 * 1",
	"(1+2*3/4-5) -  (0 * (2+1))",
	"()()(", "(1+2*3/4) ()",

	"A2+A4*4", "A4:B23",
	"TRUE", "  FALSE  ", "TRUE()",
	" abs(-4) * 2 - (-1 + 3*2) + true() - sum(2, -3) * pi()",
	"sum(-3.1314, 0, 50, 60) + (-3)*3",
	"sum()", "max(TRUE, FALSE, FALSE, TRUE)",
	"sum(10, 10+0.0, 20*20/40, 100-(30+60))",
	"sin(0)", "sin(3.1415)", "avg TRUE",
}

func TestLex(t *testing.T) {
	expr := "TRUE, FALSE, (-102030.9876 + \t \t 0.0) A2, W30:Z40 sum(9.8, 7.002, 4)"
	l := NewLexer(expr)
	fmt.Printf("%#v\n", expr)

	tok := l.NextToken()
	for !tok.EOF() {
		t.Log(fmt.Printf("%s, ", tok))
		tok = l.NextToken()
	}
}

func TestParse(t *testing.T) {
	for _, expr := range inputs {
		p := NewParser(expr, nil)
		node := p.Parse()
		fmt.Printf("%#v\nTree: %s\nResult: %#v\n\n", expr, node.inspect(0), node.Eval())
	}
}
