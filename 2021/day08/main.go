package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"math"
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

func getDigitMapping(input string) map[int]string {
	parts := strings.Split(input, " | ")
	inputDigits := strings.Split(parts[0], " ")

	contains := func(digit1, digit2 string) bool {
		return len(operators.Intersection([]rune(digit1), []rune(digit2))) == len([]rune(digit2))
	}

	numberMap := make(map[int]string, 10)
	numberMap[1] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 2
	})]
	numberMap[4] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 4
	})]
	numberMap[7] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 3
	})]
	numberMap[8] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 7
	})]
	numberMap[3] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 5 && contains(inputDigit, numberMap[1])
	})]
	numberMap[9] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 6 && contains(inputDigit, numberMap[4])
	})]
	numberMap[0] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 6 && contains(inputDigit, numberMap[1]) && inputDigit != numberMap[9]
	})]
	numberMap[6] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 6 && inputDigit != numberMap[0] && inputDigit != numberMap[9]
	})]
	numberMap[3] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 5 && contains(inputDigit, numberMap[1])
	})]
	numberMap[5] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 5 && contains(numberMap[6], inputDigit)
	})]
	numberMap[2] = inputDigits[operators.FindIndex(inputDigits, func(inputDigit string) bool {
		return len(inputDigit) == 5 && inputDigit != numberMap[3] && inputDigit != numberMap[5]
	})]
	return numberMap
}

func findKey(m map[int]string, v string) int {
	for key := range m {
		if areSamePattern(v, m[key]) {
			return key
		}
	}
	return 0
}

func getDigits(input string) (value int) {
	parts := strings.Split(input, " | ")
	outputDigits := strings.Split(parts[1], " ")
	numberMap := getDigitMapping(input)
	digits := operators.Map(outputDigits, func(outputDigit string) int {
		return findKey(numberMap, outputDigit)
	})
	for index, digit := range digits {
		value += int(float64(digit) * math.Pow(10, 3-float64(index)))
	}
	return value
}

func step1(input string) int {
	return operators.Sum(operators.Map(strings.Split(input, "\n"), func(rowInput string) int {
		return getNbEasyDigits(rowInput)
	}))
}

func step2(input string) int {
	return operators.Sum(operators.Map(strings.Split(input, "\n"), func(rowInput string) int {
		return getDigits(rowInput)
	}))
}

func main() {
	const title, day = "--- Day 8: Seven Segment Search ---", "2021/day08/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 26, "example step1")
	utils.AssertEqual(step2(example), 61229, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 288, "step1")
	utils.AssertEqual(step2(input), 940724, "step2")
}
