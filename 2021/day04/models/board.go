package models

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"strings"
)

const boardSize = 5

type Board struct {
	content    [][]int
	IsFinished bool
}

func (board *Board) Init(input string) {
	board.content = operators.Map(strings.Split(input, "\n"), func(rowInput string) []int {
		return operators.Map(operators.Filter(strings.Split(rowInput, " "),
			func(str string) bool { return len(str) > 0 },
		), utils.ParseStringToInt,
		)
	})
	board.IsFinished = false
}

func (board *Board) MarkNumber(number int) {
	for rowIndex, rowValue := range board.content {
		for columnIndex, columnValue := range rowValue {
			if columnValue == number {
				board.content[rowIndex][columnIndex] = -1
			}
		}
	}
}

func (board *Board) SumUnmarkedNumbers() int {
	return operators.Sum(operators.Map(board.content, func(row []int) int {
		return operators.Sum(operators.Filter(row, func(value int) bool {
			return value >= 0
		}))
	}))
}

func (board *Board) hasHorizontalBingo() bool {
	return operators.Any(
		operators.Map(board.content, operators.Sum),
		func(sum int) bool { return sum == -boardSize },
	)
}

func (board *Board) hasVerticalBingo() bool {
	return operators.Any(
		operators.Reduce(board.content, func(acc []int, current []int) []int {
			return operators.Add(acc, current)
		}, make([]int, boardSize)),
		func(sum int) bool { return sum == -boardSize },
	)
}

func (board *Board) HasBingo() bool {
	return board.hasHorizontalBingo() || board.hasVerticalBingo()
}

func (board *Board) MarkAsFinished() {
	board.IsFinished = true
}
