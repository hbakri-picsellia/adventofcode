package main

import (
	"adventofcode/utils"
	"fmt"
)

func SonarSweep(input string, n int) (increment int) {
	list := utils.ParseStringToIntList(input, "\n")
	increment = 0
	for index, _ := range list {
		if index > 0 && index < len(list)-n+1 {
			if utils.Sum(list[index:index+n]) > utils.Sum(list[index-1:index+n-1]) {
				increment += 1
			}
		}
	}
	return increment
}

func step1(input string) int {
	return SonarSweep(input, 1)
}

func step2(input string) int {
	return SonarSweep(input, 3)
}

func main() {
	fmt.Println("--- Day 1: Calorie Counting ---")

	example := utils.ParseFileToString("2021/day01/example.txt")
	utils.AssertEqual(step1(example), 7, "example step1")
	utils.AssertEqual(step2(example), 5, "example step2")

	input := utils.ParseFileToString("2021/day01/input.txt")
	utils.AssertEqual(step1(input), 1387, "step1")
	utils.AssertEqual(step2(input), 1362, "step2")
}
