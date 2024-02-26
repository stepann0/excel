package data

import (
	"encoding/csv"
	"os"
	"strconv"
)

func LoadCSV(t *DataTable, path string) {
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
	if x, y := t.Cols(), t.Rows(); rows > y || cols > x {
		return
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data, dtype := ConvertType(records[i][j])
			t.Put(j, i, data, dtype)
		}
	}
}

func ConvertType(record string) (any, CellType) {
	if record[0] == '=' {
		return record, Formula
	}
	if n, err := strconv.ParseFloat(record, 64); err == nil {
		return n, Number__
	}
	return record, Text
}
