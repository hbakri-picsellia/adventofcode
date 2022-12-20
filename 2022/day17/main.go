package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type Input struct {
	rockIndex int
	windIndex int
	restXList string
}

type Output struct {
	rockIndex int
	windIndex int
	restMaxX  int
}

type Func func(rest *ListComparable[Position], wind *[]int, rockIndex, windIndex, total int) (int, int)

func memorized(fn Func) Func {
	var cache = make(map[Input]Output)

	return func(rest *ListComparable[Position], wind *[]int, rockIndex, windIndex, total int) (int, int) {
		newRockIndex, newWindIndex := fn(rest, wind, rockIndex, windIndex, total)

		maxX := getMaxX(rest.List)
		restXList := strings.Join(operators.Map(getXList(rest.List, maxX), utils.ParseIntToString), ",")
		input := Input{rockIndex: rockIndex % len(ROCKS), windIndex: windIndex % len(*wind), restXList: restXList}
		if output, found := cache[input]; found {
			rep := total - newRockIndex
			rem := rep / (newRockIndex - output.rockIndex)
			rest.List = rest.Map(func(position Position) Position {
				return position.Add(Position{X: rem * (maxX - output.restMaxX)})
			})
			return newRockIndex + rem*(newRockIndex-output.rockIndex), output.windIndex
		}

		cache[input] = Output{rockIndex: newRockIndex, windIndex: newWindIndex, restMaxX: maxX}
		return newRockIndex, newWindIndex
	}
}

var ROCKS = [][]Position{
	{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
	{{X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 2, Y: 1}},
	{{X: 2, Y: 2}, {X: 1, Y: 2}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
	{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1}},
}
var DIRECTIONS = map[rune]int{
	'>': 1,
	'<': -1,
}

func display(rest ListComparable[Position]) {
	matrix := MakeMatrix[string](getMaxX(rest.List)+2, 7, ".")
	for _, position := range rest.List {
		matrix[getMaxX(rest.List)-position.X][position.Y] = "#"
	}
	matrix.Display()
}

func getMaxX(positions List[Position]) int {
	return positions.Reduce(func(previousValue, currentValue Position) Position {
		return Position{X: int(math.Max(float64(previousValue.X), float64(currentValue.X)))}
	}, positions[0]).X
}

func getXList(positions List[Position], maxX int) (result []int) {
	for y := 0; y < 7; y++ {
		result = append(result, maxX-getMaxX(positions.Filter(func(position Position) bool { return position.Y == y })))
	}
	return result
}

func PyroclasticFlow(rest *ListComparable[Position], wind *[]int, rockIndex, windIndex, total int) (int, int) {
	var rock, newRock ListComparable[Position]
	start := Position{X: 4 + getMaxX(rest.List), Y: 2}
	rock.List = ROCKS[rockIndex%len(ROCKS)]
	rock.List = rock.Map(func(position Position) Position { return position.Add(start) })
	for {
		newRock.List = rock.Map(func(position Position) Position { return position.Add(Position{Y: (*wind)[windIndex%len(*wind)]}) })
		windIndex++
		if !newRock.Intersects(*rest) && newRock.All(func(position Position) bool { return 0 <= position.Y && position.Y < 7 }) {
			rock = newRock
		}

		newRock.List = rock.Map(func(position Position) Position { return position.Add(Position{X: -1}) })
		if newRock.Intersects(*rest) {
			rest.Push(rock.List...)
			break
		}
		rock = newRock
	}
	rockIndex++
	return rockIndex, windIndex
}

func step1(input string) int {
	windIndex, wind := 0, operators.Map([]rune(input), func(char rune) int { return DIRECTIONS[char] })
	var rest ListComparable[Position]
	for y := 0; y < 7; y++ {
		rest.Push(Position{X: -1, Y: y})
	}

	rockIndex := 0
	MemoPyroclasticFlow := memorized(PyroclasticFlow)
	total := 2022
	for rockIndex < total {
		rockIndex, windIndex = MemoPyroclasticFlow(&rest, &wind, rockIndex, windIndex, total)
	}
	return getMaxX(rest.List) + 1
}

func step2(input string) int {
	windIndex, wind := 0, operators.Map([]rune(input), func(char rune) int { return DIRECTIONS[char] })
	var rest ListComparable[Position]
	for y := 0; y < 7; y++ {
		rest.Push(Position{X: -1, Y: y})
	}

	rockIndex := 0
	MemoPyroclasticFlow := memorized(PyroclasticFlow)
	total := 1000000000000
	for rockIndex < total {
		rockIndex, windIndex = MemoPyroclasticFlow(&rest, &wind, rockIndex, windIndex, total)
	}
	return getMaxX(rest.List) + 1
}

func main() {
	const title, day = "--- Day 17: Pyroclastic Flow ---", "2022/day17/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 3068, "example step1")
	utils.AssertEqual(step2(example), 1514285714288, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 3159, "step1")
	utils.AssertEqual(step2(input), 1566272189352, "step2")
}
