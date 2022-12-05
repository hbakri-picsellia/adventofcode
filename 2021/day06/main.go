package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
	"time"
)

func RecursiveLanternfish(arr []int, nbDays int) []int {
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
	return RecursiveLanternfish(arr, nbDays-1)
}

func IterativeLanternfish(arr []int, nbDays int) []int {
	for {
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
		nbDays--
	}
}

func step1(input string) int {
	start := time.Now()
	value := len(RecursiveLanternfish(
		operators.Map(strings.Split(input, ","), utils.ParseStringToInt), 80),
	)
	fmt.Println("step1", time.Since(start))
	return value
}

func step2(input string) int {
	start := time.Now()
	value := len(IterativeLanternfish(
		operators.Map(strings.Split(input, ","), utils.ParseStringToInt), 80),
	)
	fmt.Println("step2", time.Since(start))
	return value
}

func main() {
	const title, day = "--- Day 6: Lanternfish ---", "2021/day06/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 5934, "example step1")
	utils.AssertEqual(step2(example), 26984457539, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 371379, "step1")
	utils.AssertEqual(step2(input), -1, "step2")
}
