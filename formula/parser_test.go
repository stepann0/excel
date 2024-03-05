package formula

import (
	"fmt"
	"math"
	"testing"

	V "github.com/stepann0/excel/value"
)

var parseTests []string = []string{
	"100 >= 10",
	"90*80 > 100-3",
	"5 >= 4-2 = 3*2 <> 10 < 0",
	"10 =< 90",
}

func TestParse(t *testing.T) {
	fmt.Println((4 > 5) == (10 < 5))
	for _, test := range parseTests {
		p := NewParser(test, nil)
		node := p.Parse()
		fmt.Printf("%#v\nTree: %s\n\n", test, node.inspect(0))
	}
}

func TestLex(t *testing.T) {
	// expr := "TRUE, FALSE, (-102030.9876 + \t \t 0.0) A2, W30:Z40 sum(9.8, 7.002, 4)"
	expr := "1>23 4<0  <=  500.2 >=  0=0 10 <>100 <="
	l := NewLexer(expr)
	fmt.Printf("%#v\n", expr)

	tok := l.NextToken()
	for !tok.EOF() {
		t.Log(fmt.Printf("%s, ", tok))
		tok = l.NextToken()
	}
	fmt.Println()
}

var evalTest = []struct {
	expr string
	res  V.Value
}{
	{
		"((-14.98+34.241*0.4)/(-(201.2+33.241)*(0.05))-(11)+1852.098)",
		V.Float(1841.207503031),
	},
	{
		"1000 - (50 + 10 - 2000 + 9876) * 2 * (300 - 1200 + 4*120)",
		V.Int(1000 - (50+10-2000+9876)*2*(300-1200+4*120)),
	},
	{"100", V.Int(100)},
	{"+-+200", V.Int(-200)},
	{
		"sum(abs(-3), abs(+4), exp(5), 100)",
		V.Float(3 + 4 + math.Exp(5) + 100),
	},
	{"sin(150)", V.Float(math.Sin(150))},
	{"true()", V.Boolean(true)},
	{"false()", V.Boolean(false)},
	{"and(FALSE, TRUE)", V.Boolean(false)},
	{"or(FALSE, TRUE)", V.Boolean(true)},
	{"xor(FALSE, TRUE)", V.Boolean(true)},
	{"xor(1, 1, 0, 0, 1.01, TRUE, FALSE, 1>0)", V.Boolean(true)},
	{"not(FALSE)", V.Boolean(true)},
	{
		"0 < 1.0001 = (10.0 / 2 < 100/2)",
		V.Boolean(0 < 1.0001 == (10.0/2 < 100/2)),
	},
}

func ValEq(a, b V.Value) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	switch a := a.(type) {
	case V.Float:
		// almost equal
		if b, ok := b.(V.Float); !ok || math.Abs(float64(a)-float64(b)) > 0.0001 {
			return false
		}
	case V.Int:
		if b, ok := b.(V.Int); !ok || int64(a) != int64(b) {
			return false
		}
	case V.String:
		if b, ok := b.(V.String); !ok || string(a) != string(b) {
			return false
		}
	case V.Boolean:
		if b, ok := b.(V.Boolean); !ok || bool(a) != bool(b) {
			return false
		}
	case V.Error:
		if b, ok := b.(V.Error); !ok || a.Msg != b.Msg {
			return false
		}
	case V.Area:
	}
	return true
}

func TestEval(t *testing.T) {
	for _, test := range evalTest {
		fmt.Println(test.expr)
		p := NewParser(test.expr, nil)
		node := p.Parse()
		fmt.Printf("%#v\nTree: %s\nResult: %#v\n\n", test.expr, node.inspect(0), node.Eval())
		if got := node.Eval(); !ValEq(got, test.res) {
			t.Errorf("%s = %s, expected %s", test.expr, got, test.res)
		}
	}
}
