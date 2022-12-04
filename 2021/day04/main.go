package main

import (
	"adventofcode/2021/day04/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func getNumbersAndBoards(input string) ([]int, []models.Board) {
	inputs := strings.Split(input, "\n\n")
	numbers := operators.Map(strings.Split(inputs[0], ","), utils.ParseStringToInt)
	boards := make([]models.Board, len(inputs)-1)
	for index, boardInput := range inputs[1:] {
		boards[index].Init(boardInput)
	}
	return numbers, boards
}

func GiantSquid(input string) (scores []int) {
	numbers, boards := getNumbersAndBoards(input)
	for _, number := range numbers {
		for index := range boards {
			if !boards[index].IsFinished {
				boards[index].MarkNumber(number)
				if boards[index].HasBingo() {
					scores = append(scores, number*boards[index].SumUnmarkedNumbers())
					boards[index].MarkAsFinished()
				}
			}
		}
	}
	return scores
}

func step1(input string) int {
	boardScores := GiantSquid(input)
	return boardScores[0]
}

func step2(input string) int {
	boardScores := GiantSquid(input)
	return boardScores[len(boardScores)-1]
}

func main() {
	const title, day = "--- Day 4: Giant Squid ---", "2021/day04/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 4512, "example step1")
	utils.AssertEqual(step2(example), 1924, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 10374, "step1")
	utils.AssertEqual(step2(input), 24742, "step2")
}
