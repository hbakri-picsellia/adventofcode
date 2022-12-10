package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

func signalStrength(cycle, X int) int {
	if (cycle-20)%40 == 0 {
		return cycle * X
	}
	return 0
}

func step1(input string) (score int) {
	commandList := strings.Split(input, "\n")
	cycle, X := 0, 1
	for _, value := range commandList {
		parts := strings.Split(value, " ")
		switch parts[0] {
		case "noop":
			cycle += 1
			score += signalStrength(cycle, X)
		case "addx":
			cycle += 1
			score += signalStrength(cycle, X)
			cycle += 1
			score += signalStrength(cycle, X)
			X += utils.ParseStringToInt(parts[1])
		}
	}
	return score
}

func draw(cycle, X int) (pixel string) {
	if math.Abs(float64(cycle%40-X)) <= 1 {
		pixel = "#"
	} else {
		pixel = "."
	}
	if cycle%40 == 0 {
		return "\n" + pixel
	} else {
		return pixel
	}
}

func step2(input string) (text string) {
	commandList := strings.Split(input, "\n")
	cycle, X := 0, 1
	for _, value := range commandList {
		parts := strings.Split(value, " ")
		switch parts[0] {
		case "noop":
			text += draw(cycle, X)
			cycle += 1
		case "addx":
			text += draw(cycle, X)
			cycle += 1
			text += draw(cycle, X)
			cycle += 1
			X += utils.ParseStringToInt(parts[1])
		}
	}
	return text
}

func main() {
	const title, day = "--- Day 10: Cathode-Ray Tube ---", "2022/day10/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 13140, "example step1")
	fmt.Println(step2(example))

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 11960, "step1")
	fmt.Println(step2(input))
}
