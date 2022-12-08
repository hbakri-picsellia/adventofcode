package models

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"strings"
)

type Matrix [][]int

func (matrix *Matrix) Decode(input string, rowSeparator string, columnSeparator string) {
	*matrix = operators.Map(strings.Split(input, rowSeparator), func(rowInput string) []int {
		return operators.Map(operators.Filter(strings.Split(rowInput, columnSeparator),
			func(str string) bool { return len(str) > 0 },
		), utils.ParseStringToInt)
	})
}

func (matrix *Matrix) GetRow(index int) []int {
	return (*matrix)[index]
}

func (matrix *Matrix) GetColumn(index int) []int {
	return operators.Map(*matrix, func(row []int) int {
		return row[index]
	})
}
