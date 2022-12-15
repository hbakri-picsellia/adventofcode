package main

import (
	"adventofcode/2022/day14/structs"
	"adventofcode/models"
	"adventofcode/operators"
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

func MakeMatrix[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func drawCave(cave models.Matrix[Material]) {
	for _, row := range cave {
		for _, rowValue := range row {
			switch rowValue {
			case Air:
				fmt.Print(".")
			case Rock:
				fmt.Print("#")
			case Sand:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

func step1(input string) int {
	var points []models.Position
	minJ, maxI, maxJ := math.MaxInt, math.MinInt, math.MinInt
	for _, rowPath := range strings.Split(input, "\n") {
		path := structs.Path{}
		path.Decode(rowPath)
		for _, point := range path.GetCoveredPoints() {
			points = append(points, models.Position{I: point.Y, J: point.X})
			maxI = int(math.Max(float64(point.Y), float64(maxI)))
			minJ = int(math.Min(float64(point.X), float64(minJ)))
			maxJ = int(math.Max(float64(point.X), float64(maxJ)))
		}
	}
	var cave models.Matrix[Material] = MakeMatrix[Material](maxI+1, maxJ+1-minJ)
	for _, point := range points {
		cave[point.I][point.J-minJ] = Rock
	}
	sandSourcePosition := models.Position{I: 0, J: 500 - minJ}
	iteration := 0
	var newSandPosition models.Position
	for {
		newSandPosition = models.Position{I: operators.FindIndex(cave.GetColumn(sandSourcePosition.J), func(material Material) bool {
			return material != Air
		}) - 1, J: sandSourcePosition.J}
		for cave[newSandPosition.I+1][newSandPosition.J-1] == Air || cave[newSandPosition.I+1][newSandPosition.J+1] == Air {
			// TODO: this hijo de puta is going in a straight diagonal line and get out of the screen
			if cave[newSandPosition.I+1][newSandPosition.J-1] == Air {
				newSandPosition.I += 1
				newSandPosition.J -= 1
			} else if cave[newSandPosition.I+1][newSandPosition.J+1] == Air {
				newSandPosition.I += 1
				newSandPosition.J += 1
			}
		}
		if !cave.Contains(newSandPosition) {
			break
		}
		cave[newSandPosition.I][newSandPosition.J] = Sand
		drawCave(cave)
		iteration++
	}
	return iteration
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

	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
