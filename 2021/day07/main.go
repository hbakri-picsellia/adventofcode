package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func getMedian(list []int) int {
	sort.Ints(list)
	return list[len(list)/2]
}

func getMean(list []int) int {
	return int(math.Round(float64(operators.Sum(list)) / float64(len(list))))
}

func FuelCalculation(list []int, position int, metric func(int) int) (fuel int) {
	for _, value := range list {
		fuel += metric(position - value)
	}
	return fuel
}

func step1(input string) int {
	list := operators.Map(strings.Split(input, ","), utils.ParseInt)
	median := getMedian(list)
	candidates := []int{median - 1, median, median + 1}
	return operators.Min(operators.Map(candidates, func(candidate int) int {
		return FuelCalculation(list, candidate, func(value int) int {
			return int(math.Abs(float64(value)))
		})
	}))
}

func step2(input string) int {
	list := operators.Map(strings.Split(input, ","), utils.ParseInt)
	mean := getMean(list)
	candidates := []int{mean - 1, mean, mean + 1}
	return operators.Min(operators.Map(candidates, func(candidate int) int {
		return FuelCalculation(list, candidate, func(value int) int {
			return int(math.Abs(float64(value))) * (int(math.Abs(float64(value))) + 1) / 2
		})
	}))
}

func main() {
	const title, day = "--- Day 7: The Treachery of Whales ---", "2021/day07/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 37, "example step1")
	utils.AssertEqual(step2(example), 168, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 331067, "step1")
	utils.AssertEqual(step2(input), 92881128, "step2")
}
