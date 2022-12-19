package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
)

var ROCKS = [][]Position{
	{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
	{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}},
	{{X: 0, Y: 2}, {X: 1, Y: 2}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}},
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
	{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}},
}
var DIRECTIONS = map[rune]int{
	'>': 1,
	'<': -1,
}

func getMaxY(positions List[Position]) int {
	return positions.Reduce(func(previousValue, currentValue Position) Position {
		return Position{Y: int(math.Max(float64(previousValue.Y), float64(currentValue.Y)))}
	}, positions[0]).Y
}

func step1(input string) int {
	wind := operators.Map([]rune(input), func(char rune) int { return DIRECTIONS[char] })
	var initialState []Position
	for y := 0; y < 7; y++ {
		initialState = append(initialState, Position{X: -1, Y: y})
	}

	for i := 0; i < 2022; i++ {

	}
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 17: Pyroclastic Flow ---", "2022/day17/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
