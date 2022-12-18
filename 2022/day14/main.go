package main

import (
	"adventofcode/2022/day14/structs"
	"adventofcode/operators"
	. "adventofcode/structs"
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

func drawCave(cave Matrix[Material]) {
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

func RegolithReservoir(cave *Matrix[Material], sandSource Position) Position {
	n, m := (*cave).GetDimension()
	nextMaterialIndex := operators.FindIndex(cave.GetColumn(sandSource.Y)[sandSource.X:], func(material Material) bool {
		return material != Air
	})
	if nextMaterialIndex < 0 {
		return Position{X: -1, Y: sandSource.Y}
	}
	sandPosition := Position{X: nextMaterialIndex + sandSource.X - 1, Y: sandSource.Y}
	if sandPosition.X+1 >= n || sandPosition.Y-1 < 0 {
		return Position{X: sandPosition.X + 1, Y: sandPosition.Y - 1}
	} else if (*cave)[sandPosition.X+1][sandPosition.Y-1] == Air {
		return RegolithReservoir(cave, Position{X: sandPosition.X + 1, Y: sandPosition.Y - 1})
	} else if sandPosition.Y+1 >= m {
		return Position{X: sandPosition.X + 1, Y: sandPosition.Y + 1}
	} else if (*cave)[sandPosition.X+1][sandPosition.Y+1] == Air {
		return RegolithReservoir(cave, Position{X: sandPosition.X + 1, Y: sandPosition.Y + 1})
	} else {
		return sandPosition
	}
}

func step1(input string) int {
	var points []Position
	minJ, maxI, maxJ := math.MaxInt, math.MinInt, math.MinInt
	for _, rowPath := range strings.Split(input, "\n") {
		path := structs.Path{}
		path.Decode(rowPath)
		for _, point := range path.GetCoveredPoints() {
			points = append(points, Position{X: point.Y, Y: point.X})
			maxI = int(math.Max(float64(point.Y), float64(maxI)))
			minJ = int(math.Min(float64(point.X), float64(minJ)))
			maxJ = int(math.Max(float64(point.X), float64(maxJ)))
		}
	}
	var cave Matrix[Material] = MakeMatrix[Material](maxI+1, maxJ+1-minJ)
	for _, point := range points {
		cave[point.X][point.Y-minJ] = Rock
	}
	sandSource := Position{X: 0, Y: 500 - minJ}

	iteration := 0
	for {
		newSandPosition := RegolithReservoir(&cave, sandSource)
		if !cave.Contains(newSandPosition) {
			break
		}
		cave[newSandPosition.X][newSandPosition.Y] = Sand
		iteration++
	}
	//drawCave(cave)
	return iteration
}

func step2(input string) int {
	var points []Position
	minJ, maxI, maxJ := math.MaxInt, math.MinInt, math.MinInt
	for _, rowPath := range strings.Split(input, "\n") {
		path := structs.Path{}
		path.Decode(rowPath)
		for _, point := range path.GetCoveredPoints() {
			points = append(points, Position{X: point.Y, Y: point.X})
			maxI = int(math.Max(float64(point.Y), float64(maxI)))
			minJ = int(math.Min(float64(point.X), float64(minJ)))
			maxJ = int(math.Max(float64(point.X), float64(maxJ)))
		}
	}
	var cave Matrix[Material] = MakeMatrix[Material](maxI+3, maxJ+1+maxI+1)
	for _, point := range points {
		cave[point.X][point.Y] = Rock
	}
	for j := 0; j < len(cave[0]); j++ {
		cave[maxI+2][j] = Rock
	}
	sandSource := Position{X: 0, Y: 500}

	iteration := 0
	for {
		newSandPosition := RegolithReservoir(&cave, sandSource)
		if newSandPosition == sandSource {
			iteration++
			break
		}
		cave[newSandPosition.X][newSandPosition.Y] = Sand
		iteration++
	}
	//drawCave(cave)
	return iteration
}

func main() {
	const title, day = "--- Day 14: Regolith Reservoir ---", "2022/day14/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 24, "example step1")
	utils.AssertEqual(step2(example), 93, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 774, "step1")
	utils.AssertEqual(step2(input), 22499, "step2")
}
