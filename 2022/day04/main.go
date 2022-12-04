package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func getSections(assignment string) []int {
	return operators.Map(strings.Split(assignment, "-"), utils.ParseStringToInt)
}

func fullyContains(assignment1 string, assignment2 string) bool {
	sections1 := getSections(assignment1)
	sections2 := getSections(assignment2)
	return sections1[0] <= sections2[0] && sections1[1] >= sections2[1]
}

func overlapAtAll(assignment1 string, assignment2 string) bool {
	sections1 := getSections(assignment1)
	sections2 := getSections(assignment2)
	return sections1[0] <= sections2[0] && sections1[1] >= sections2[0]
}

func CampCleanup(input string, f func(string, string) bool) int {
	assignmentPairs := strings.Split(input, "\n")
	return operators.Reduce(operators.Map(assignmentPairs, func(assignmentPair string) bool {
		assignment1 := strings.Split(assignmentPair, ",")[0]
		assignment2 := strings.Split(assignmentPair, ",")[1]
		return f(assignment1, assignment2) || f(assignment2, assignment1)
	}), func(acc int, t bool) int {
		return acc + utils.BoolToInt(t)
	}, 0)
}

func step1(input string) int {
	return CampCleanup(input, fullyContains)
}

func step2(input string) int {
	return CampCleanup(input, overlapAtAll)
}

func main() {
	fmt.Println("--- Day 4: Camp Cleanup ---")

	example := utils.ParseFileToString("2022/day04/example.txt")
	utils.AssertEqual(step1(example), 2, "example step1")
	utils.AssertEqual(step2(example), 4, "example step2")

	input := utils.ParseFileToString("2022/day04/input.txt")
	utils.AssertEqual(step1(input), 503, "step1")
	utils.AssertEqual(step2(input), 827, "step2")
}
