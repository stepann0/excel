package csv

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/stepann0/tercel/formula"
	"github.com/stepann0/tercel/ui"
)

func LoadCSV(t *ui.Table, path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	if len(records) == 0 {
		return
	}
	rows, cols := len(records), len(records[0])
	if x, y := t.Size(); rows > y || cols > x {
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data, dtype := formula.ConvertType(records[i][j])
			t.Put(j, i, data, dtype)
		}
	}
}

func ConvertType(record string) (any, formula.CellType) {
	if record[0] == '=' {
		return record, formula.Formula
	}
	if n, err := strconv.ParseFloat(record, 64); err == nil {
		return n, formula.Number
	}
	return record, formula.Text
}
