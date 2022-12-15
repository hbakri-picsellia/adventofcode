package main

import (
	"adventofcode/2022/day14/structs"
	"adventofcode/models"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type Material int

const (
	Air  Material = 0
	Rock          = 1
	Sand          = 2
)

func step1(input string) int {
	var points []models.Position
	var minI, maxI, maxJ int
	for _, rowPath := range strings.Split(input, "\n") {
		path := structs.Path{}
		path.Decode(rowPath)
		for _, point := range path.GetCoveredPoints() {
			points = append(points, models.Position{I: point.X, J: point.Y})
			minI = int(math.Min(float64(point.X), float64(minI)))
			maxI = int(math.Max(float64(point.X), float64(maxI)))
			maxJ = int(math.Max(float64(point.Y), float64(maxJ)))
		}
	}
	cave := [][]int{}
	for _, point := range points {
		cave[point.I][point.J] = Rock
	}
	fmt.Println(cave)
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 14: Regolith Reservoir ---", "2022/day14/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 24, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
