package data

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
}

// func TestLex(t *testing.T) {
// 	l := NewLexer(form)
// 	tok := l.NextToken()
// 	for !tok.EOF() {
// 		t.Log(fmt.Println(tok))
// 		tok = l.NextToken()
// 	}
// }

func TestParse(t *testing.T) {
	for _, expr := range inputs {
		p := NewParser(expr, nil)
		node := p.Parse()
		fmt.Printf("%s\n%s\nResult: %s\n\n", expr, node.inspect(0), node.Eval())
	}
}
