package main

import (
	"adventofcode/utils"
	"fmt"
)

func step1(input string) int {
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 22: Monkey Map ---", "2022/day22/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	utils.AssertEqual(step2(example), -1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), -1, "step1")
	utils.AssertEqual(step2(input), -1, "step2")
}
