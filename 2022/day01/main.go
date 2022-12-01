package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func getSums(input string) []int {
	return utils.Map(strings.Split(input, "\n\n"), func(s string) int {
		return utils.Sum(utils.Map(strings.Split(s, "\n"), func(s string) int {
			num, _ := strconv.ParseInt(s, 10, 0)
			return int(num)
		}))
	})
}

func step1(input string) int {
	return utils.Max(getSums(input))
}

func step2(input string) int {
	return utils.Sum(utils.MaxN(getSums(input), 3))
}

func main() {
	input := utils.ParseTxtFile("2022/day01/input.txt")
	fmt.Println("step1: ", step1(input))
	fmt.Println("step2: ", step2(input))
}
