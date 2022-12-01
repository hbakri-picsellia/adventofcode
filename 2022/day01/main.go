package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func CalorieCounter(input string) []int {
	return utils.Map(strings.Split(input, "\n\n"), func(s string) int {
		return utils.Sum(utils.ParseStringToIntList(s, "\n"))
	})
}

func step1(input string) int {
	return utils.Max(CalorieCounter(input))
}

func step2(input string) int {
	return utils.Sum(utils.MaxN(CalorieCounter(input), 3))
}

func main() {
	fmt.Println("--- Day 1: Sonar Sweep ---")

	example := utils.ParseFileToString("2022/day01/example.txt")
	utils.AssertEqual(step1(example), 24000, "example step1")
	utils.AssertEqual(step2(example), 45000, "example step2")

	input := utils.ParseFileToString("2022/day01/input.txt")
	utils.AssertEqual(step1(input), 71300, "step1")
	utils.AssertEqual(step2(input), 209691, "step2")
}
