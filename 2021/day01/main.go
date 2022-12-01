package main

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseIntListFromString(text string, separator string) []int {
	return utils.Map(strings.Split(text, separator), func(s string) int {
		num, _ := strconv.ParseInt(s, 10, 0)
		return int(num)
	})
}

func step1(input string) (increment int) {
	increment = 0
	list := parseIntListFromString(input, "\n")
	for index, value := range list {
		if index != 0 && value > list[index-1] {
			increment += 1
		}
	}
	return increment
}

func main() {
	example := utils.ParseTxtFile("2021/day01/example.txt")
	input := utils.ParseTxtFile("2021/day01/input.txt")
	fmt.Println("step1 example: ", step1(example))
	fmt.Println("step1: ", step1(input))
	//fmt.Println("step2 example: ", step2(example))
	//fmt.Println("step2: ", step2(input))
}
