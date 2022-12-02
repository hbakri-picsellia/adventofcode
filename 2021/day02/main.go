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

type SubmarinePosition struct {
	horizontal int
	depth      int
	aim        int
}

func Dive(input string, f func(string, int, *SubmarinePosition)) int {
	instructions := strings.Split(input, "\n")
	position := SubmarinePosition{}
	for _, value := range instructions {
		instruction := strings.Split(value, " ")
		direction := instruction[0]
		number := utils.ParseStringToInt(instruction[1])
		f(direction, number, &position)
	}
	return position.horizontal * position.depth
}

func step1(input string) int {
	return Dive(input, func(direction string, number int, position *SubmarinePosition) {
		switch direction {
		case Forward:
			position.horizontal += number
		case Down:
			position.depth += number
		case Up:
			position.depth -= number
		}
	})
}

func step2(input string) int {
	return Dive(input, func(direction string, number int, position *SubmarinePosition) {
		switch direction {
		case Forward:
			position.horizontal += number
			position.depth += position.aim * number
		case Down:
			position.aim += number
		case Up:
			position.aim -= number
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
