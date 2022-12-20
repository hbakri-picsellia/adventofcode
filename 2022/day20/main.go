package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

type Node struct {
	number int
	left   *Node
	right  *Node
}

func step1(input string) int {
	numbers := operators.Map(strings.Split(input, "\n"), utils.ParseStringToInt)
	nodes := operators.Map(numbers, func(number int) Node {
		return Node{number: number}
	})
	for index := range nodes {
		nodes[index].left = &nodes[(len(nodes)+index-1)%len(nodes)]
		nodes[index].right = &nodes[(index+1)%len(nodes)]
	}
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 20: Grove Positioning System ---", "2022/day20/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	utils.AssertEqual(step2(example), -1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), -1, "step1")
	utils.AssertEqual(step2(input), -1, "step2")
}
