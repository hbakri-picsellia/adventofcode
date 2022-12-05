package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func Lanternfish(arr []int, nbDays int) []int {
	if nbDays == 0 {
		return arr
	}
	nbNewElements := 0
	for index, _ := range arr {
		if arr[index] == 0 {
			nbNewElements += 1
			arr[index] = 6
		} else {
			arr[index] -= 1
		}
	}
	for i := 0; i < nbNewElements; i++ {
		arr = append(arr, 8)
	}
	return Lanternfish(arr, nbDays-1)
}

func step1(input string) int {
	return len(Lanternfish(
		operators.Map(strings.Split(input, ","), utils.ParseStringToInt), 80),
	)
}

func step2(input string) int {
	return len(Lanternfish(
		operators.Map(strings.Split(input, ","), utils.ParseStringToInt), 256),
	)
}

func main() {
	const title, day = "--- Day 6: Lanternfish ---", "2021/day06/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 5934, "example step1")
	utils.AssertEqual(step2(example), 26984457539, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 371379, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
