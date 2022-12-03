package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
	"unicode"
)

func getLetterPriority(char rune) int {
	if unicode.IsLower(char) {
		return int(char - 'a' + 1)
	} else {
		return int(char - 'A' + 27)
	}
}

func step1(input string) int {
	rucksacks := strings.Split(input, "\n")
	return operators.Sum(operators.Map(rucksacks, func(rucksack string) int {
		compartments := operators.Chunk([]rune(rucksack), len(rucksack)/2)
		sharedCharacters := operators.Intersection(compartments[0], compartments[1])
		return getLetterPriority(sharedCharacters[0])
	}))
}

func step2(input string) int {
	groups := operators.Chunk(strings.Split(input, "\n"), 3)
	return operators.Sum(operators.Map(groups, func(group []string) int {
		sharedCharacters := operators.Reduce(group, func(acc []rune, current string) []rune {
			return operators.Intersection(acc, []rune(current))
		}, []rune(group[0]))
		return getLetterPriority(sharedCharacters[0])
	}))
}

func main() {
	fmt.Println("--- Day 3: Rucksack Reorganization ---")

	example := utils.ParseFileToString("2022/day03/example.txt")
	utils.AssertEqual(step1(example), 157, "example step1")
	utils.AssertEqual(step2(example), 70, "example step2")

	input := utils.ParseFileToString("2022/day03/input.txt")
	utils.AssertEqual(step1(input), 7908, "step1")
	utils.AssertEqual(step2(input), 2838, "step2")
}
