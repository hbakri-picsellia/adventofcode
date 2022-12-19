package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
)

//type Params struct {
//	rockIndex int
//	windIndex int
//	XList     string
//}
//
//type Func func(valves Valves, distance Matrix[float64], minutes float64, current Valve, candidates List[Valve], isElephant bool) float64
//
//var cache = make(map[Params]float64)
//
//func memorized(fn Func) Func {
//	return func(valves Valves, distance Matrix[float64], minutes float64, current Valve, candidates List[Valve], isElephant bool) float64 {
//		names := operators.Map(candidates, func(valve Valve) string {
//			return valve.Name
//		})
//		sort.Strings(names)
//		input := Params{minutes: minutes, currentName: current.Name, candidatesName: strings.Join(names, ","), isElephant: isElephant}
//		if val, found := cache[input]; found {
//			return val
//		}
//
//		result := fn(valves, distance, minutes, current, candidates, isElephant)
//		cache[input] = result
//		return result
//	}
//}

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

func getMaxX(positions ListComparable[Position]) int {
	return positions.Reduce(func(previousValue, currentValue Position) Position {
		return Position{X: int(math.Max(float64(previousValue.X), float64(currentValue.X)))}
	}, positions.List[0]).X
}

func PyroclasticFlow() {

}

func step1(input string) int {
	wind := operators.Map([]rune(input), func(char rune) int { return DIRECTIONS[char] })
	w := 0

	var rest ListComparable[Position]
	for y := 0; y < 7; y++ {
		rest.Push(Position{X: -1, Y: y})
	}

	var rock, newRock ListComparable[Position]

	for i := 0; i < 2022; i++ {
		start := Position{X: 4 + getMaxX(rest), Y: 2}
		rock.List = ROCKS[i%len(ROCKS)]
		rock.List = rock.Map(func(position Position) Position { return position.Add(start) })
		for {
			newRock.List = rock.Map(func(position Position) Position { return position.Add(Position{Y: wind[w%len(wind)]}) })
			w++
			if !newRock.Intersects(rest) && newRock.All(func(position Position) bool { return 0 <= position.Y && position.Y < 7 }) {
				rock = newRock
			}

			newRock.List = rock.Map(func(position Position) Position { return position.Add(Position{X: -1}) })
			if newRock.Intersects(rest) {
				rest.Push(rock.List...)
				break
			}
			rock = newRock
		}

		//fmt.Println()
		//fmt.Println("round", i)
		//matrix := MakeMatrix[string](getMaxX(rest)+2, 7, ".")
		//for _, position := range rest.List {
		//	matrix[getMaxX(rest)-position.X][position.Y] = "#"
		//}
		//matrix.Display()
	}
	return getMaxX(rest) + 1
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 17: Pyroclastic Flow ---", "2022/day17/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 3068, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
