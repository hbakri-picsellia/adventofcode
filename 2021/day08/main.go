package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func areSamePattern(string1 string, string2 string) bool {
	intersection := operators.Intersection([]rune(string1), []rune(string2))
	return len(intersection) == len(string1) && len(intersection) == len(string2)
}

func isEasyDigit(input string) bool {
	return (len(input) >= 2 && len(input) <= 4) || len(input) == 7
}

func getNbEasyDigits(input string) (value int) {
	parts := strings.Split(input, " | ")
	inputDigits := strings.Split(parts[0], " ")
	outputDigits := strings.Split(parts[1], " ")
	for _, outputDigit := range outputDigits {
		value += utils.BoolToInt(operators.Any(inputDigits, func(inputDigit string) bool {
			return isEasyDigit(outputDigit) && areSamePattern(inputDigit, outputDigit)
		}))
	}
	return value
}

func getDigits(input string) (value int) {
	parts := strings.Split(input, " | ")
	inputDigits := strings.Split(parts[0], " ")
	outputDigits := strings.Split(parts[1], " ")
	for _, outputDigit := range outputDigits {
		value += utils.BoolToInt(operators.Any(inputDigits, func(inputDigit string) bool {
			return isEasyDigit(outputDigit) && areSamePattern(inputDigit, outputDigit)
		}))
	}
	return value
}

func step1(input string) int {
	return operators.Sum(operators.Map(strings.Split(input, "\n"), func(rowInput string) int {
		return getNbEasyDigits(rowInput)
	}))
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 8: Seven Segment Search ---", "2021/day08/"
	fmt.Println(title)

	//example := utils.ParseFileToString(day + "example.txt")
	//utils.AssertEqual(step1(example), 26, "example step1")
	//utils.AssertEqual(step2(example), 5353, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), 288, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")

	example := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	parts := strings.Split(example, " | ")
	inputDigits := strings.Split(parts[0], " ")
	outputDigits := strings.Split(parts[1], " ")
	fmt.Println(inputDigits, outputDigits)
}
