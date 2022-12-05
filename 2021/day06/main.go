package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func RecursiveLanternfish(arr []int, nbDays int) int {
	if nbDays == 0 {
		return len(arr)
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
	return RecursiveLanternfish(arr, nbDays-1)
}

func IterativeLanternfish(arr []int, nbDays int) (sum int) {
	lanternFishes := make(map[int]int, 9)
	for _, value := range arr {
		lanternFishes[value] += 1
	}

	nbNewLanternFishes := 0
	for day := 0; day < nbDays; day++ {
		nbNewLanternFishes = lanternFishes[0]
		for index := 1; index <= 8; index++ {
			lanternFishes[index-1] = lanternFishes[index]
		}
		lanternFishes[6] += nbNewLanternFishes
		lanternFishes[8] = nbNewLanternFishes
	}
	for _, value := range lanternFishes {
		sum += value
	}
	return sum
}

func step1(input string) int {
	initialState := operators.Map(strings.Split(input, ","), utils.ParseStringToInt)
	return RecursiveLanternfish(initialState, 80)
}

func step2(input string) int {
	initialState := operators.Map(strings.Split(input, ","), utils.ParseStringToInt)
	return IterativeLanternfish(initialState, 256)
}

func main() {
	const title, day = "--- Day 6: Lanternfish ---", "2021/day06/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 5934, "example step1")
	utils.AssertEqual(step2(example), 26984457539, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 371379, "step1")
	utils.AssertEqual(step2(input), 1674303997472, "step2")
}
