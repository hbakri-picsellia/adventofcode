package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

var DIRECTIONS = [6][3]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {0, 0, -1}, {0, -1, 0}, {-1, 0, 0}}

func step1(input string) int {
	positions := ListComparable[Position3D]{List: operators.Map(strings.Split(input, "\n"), MakePosition3D)}
	nbSides := 6 * len(positions.List)
	for _, position := range positions.List {
		for _, direction := range DIRECTIONS {
			if positions.Contains(position.Add(direction)) {
				nbSides--
			}
		}
	}
	return nbSides
}

func step2(input string) int {
	positions := ListComparable[Position3D]{List: operators.Map(strings.Split(input, "\n"), MakePosition3D)}

	maxPosition := operators.Reduce(positions.List, func(acc, value Position3D) Position3D {
		return Position3D{
			X: int(math.Max(float64(acc.X), float64(value.X))),
			Y: int(math.Max(float64(acc.Y), float64(value.Y))),
			Z: int(math.Max(float64(acc.Z), float64(value.Z))),
		}
	}, Position3D{X: math.MinInt, Y: math.MinInt, Z: math.MinInt})
	maxPosition = maxPosition.Add([3]int{1, 1, 1})
	minPosition := operators.Reduce(positions.List, func(acc, value Position3D) Position3D {
		return Position3D{
			X: int(math.Min(float64(acc.X), float64(value.X))),
			Y: int(math.Min(float64(acc.Y), float64(value.Y))),
			Z: int(math.Min(float64(acc.Z), float64(value.Z))),
		}
	}, Position3D{X: math.MaxInt, Y: math.MaxInt, Z: math.MaxInt})
	minPosition = minPosition.Add([3]int{-1, -1, -1})

	queue := ListComparable[Position3D]{List: []Position3D{minPosition}}
	visited := map[Position3D]bool{}
	nbSides := 0
	for !queue.IsEmpty() {
		currentPosition, _ := queue.Shift()
		if visited[currentPosition] {
			continue
		}
		visited[currentPosition] = true

		for _, direction := range DIRECTIONS {
			neighbor := currentPosition.Add(direction)
			if minPosition.IsInferior(neighbor) && neighbor.IsInferior(maxPosition) {
				if positions.Contains(neighbor) {
					nbSides++
				} else {
					queue.Push(neighbor)
				}
			}
		}
	}
	return nbSides
}

func main() {
	const title, day = "--- Day 18: Boiling Boulders ---", "2022/day18/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 64, "example step1")
	utils.AssertEqual(step2(example), 58, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 3522, "step1")
	utils.AssertEqual(step2(input), 2074, "step2")
}
