package formula

import (
	"fmt"
	"testing"
)

var inputs []string = []string{
	" abs(-4) * 2 - (-1 + 3*2) + true() - sum(2, -3) * pi()",
	"sum(-3.1314, 0, 50, 60) + (-3)*3",
	"A2+A4*4",
	"A4:B23",
	"(1+2*3/4-5) -  (0 * (2+1))",
	"TRUE", "  FALSE  ", "TRUE()",
	"-45000",
	"sum()", "max(TRUE, FALSE, FALSE, TRUE)",
	"3+2*4",
	"(200 / (4 * 0.5)) + 50 * 1",
	"sum(10, 10+0.0, 20*20/40, 100-(30+60))",
	"sin(0)", "sin(3.1415)",
}

func TestLex(t *testing.T) {
	l := NewLexer("TRUE, FALSE, (-102030.9876 + \t \t 0.0) A2, W30:Z40 sum(9.8, 7.002, 4)")
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
		fmt.Printf("%s\n%s\nResult: %s\n\n", expr, node.inspect(0), node.Eval())
	}
}
