package formula

import (
	"reflect"
	"slices"
	"strings"
	"testing"

	V "github.com/stepann0/excel/value"
)

var rngTests = []struct {
	input string
	res   []V.Value
}{
	{
		"A1:A4",
		[]V.Value{V.Int(1), V.Int(5), V.Boolean(true), V.Int(0)},
	},
	{
		"B4:D4",
		[]V.Value{V.Int(90), V.Int(80), V.Int(70)},
	},
	{
		"A6:F6",
		[]V.Value{V.Boolean(false), V.Boolean(false), V.Boolean(true), V.Boolean(true), nil, nil},
	},
	{
		"B1:B9",
		[]V.Value{V.Int(3), V.Int(-100), V.Float(2.50), V.Int(90), V.Int(22), V.Boolean(false), nil, nil, nil},
	},
}

func TestRng(t *testing.T) {
	table := NewTable(13, 13)
	table.LoadCSV("/home/stepaFedora/Документы/small.csv")
	for _, test := range rngTests {
		a := strings.Split(test.input, ":")
		ref1, ref2 := a[0], a[1]
		rng := table.GetRange(ref1, ref2)
		eq := slices.EqualFunc[[]V.Value](rng, test.res, func(a, b V.Value) bool {
			return reflect.DeepEqual(a, b)
		})
		if !eq {
			t.Errorf("GetRange(%s): got %s, expected %s", test.input, rng, test.res)
		}
	}
}
