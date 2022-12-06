package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
)

func TuningTrouble(input string, windowSize int) int {
	for i := windowSize; i < len(input); i++ {
		if !operators.HasDuplicates([]rune(input[i-windowSize : i])) {
			return i
		}
	}
	return 0
}

func step1(input string) int {
	return TuningTrouble(input, 4)
}

func step2(input string) int {
	return TuningTrouble(input, 14)
}

func main() {
	const title, day = "--- Day 6: Tuning Trouble ---", "2022/day06/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 7, "example step1")
	utils.AssertEqual(step2(example), 19, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 1833, "step1")
	utils.AssertEqual(step2(input), 3425, "step2")
}
