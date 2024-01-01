package formula

import "strconv"

type CellType int

const (
	Text CellType = iota
	Number
	Formula
)

func ConvertType(record string) (any, CellType) {
	if record[0] == '=' {
		return record, Formula
	}
	if n, err := strconv.ParseFloat(record, 64); err == nil {
		return n, Number
	}
	return record, Text
}
