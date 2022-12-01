package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

const (
	Forward string = "forward"
	Down           = "down"
	Up             = "up"
)

func Dive(input string, f func(string, int, *int, *int, *int)) int {
	instructions := strings.Split(input, "\n")
	horizontal, depth, aim := 0, 0, 0
	for _, value := range instructions {
		instruction := strings.Split(value, " ")
		direction := instruction[0]
		number := utils.ParseStringToInt(instruction[1])
		f(direction, number, &horizontal, &depth, &aim)
	}
	return horizontal * depth
}

func step1(input string) int {
	return Dive(input, func(direction string, number int, horizontal *int, depth *int, aim *int) {
		switch direction {
		case Forward:
			*horizontal += number
		case Down:
			*depth += number
		case Up:
			*depth -= number
		}
	})
}

func step2(input string) int {
	return Dive(input, func(direction string, number int, horizontal *int, depth *int, aim *int) {
		switch direction {
		case Forward:
			*horizontal += number
			*depth += *aim * number
		case Down:
			*aim += number
		case Up:
			*aim -= number
		}
	})
}

func main() {
	fmt.Println("--- Day 2: Dive! ---")

	example := utils.ParseFileToString("2021/day02/example.txt")
	utils.AssertEqual(step1(example), 150, "example step1")
	utils.AssertEqual(step2(example), 900, "example step2")

	input := utils.ParseFileToString("2021/day02/input.txt")
	utils.AssertEqual(step1(input), 1648020, "step1")
	utils.AssertEqual(step2(input), 1759818555, "step2")
}
