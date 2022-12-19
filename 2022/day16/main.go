package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

func FloydWarshall(adjacencyMatrix Matrix[float64]) (W Matrix[float64]) {
	W = adjacencyMatrix
	n := len(W)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				W[i][j] = math.Min(W[i][j], W[i][k]+W[k][j])
			}
		}
	}
	return W
}

type Params struct {
	minutes        int
	currentName    string
	candidatesName string
}

type Func func(valves Valves, distance Matrix[float64], minutes int, current Valve, candidates List[Valve]) int

func memorized(fn Func) Func {
	cache := make(map[Params]int)

	return func(valves Valves, distance Matrix[float64], minutes int, current Valve, candidates List[Valve]) int {
		names := operators.Map(candidates, func(valve Valve) string {
			return valve.Name
		})
		sort.Strings(names)
		input := Params{minutes: minutes, currentName: current.Name, candidatesName: strings.Join(names, ",")}
		if val, found := cache[input]; found {
			log.Println("Read from cache")
			return val
		}

		result := fn(valves, distance, minutes, current, candidates)
		cache[input] = result
		return result
	}
}

func ProboscideaVolcanium(valves Valves, distance Matrix[float64], minutes int, current Valve, candidates List[Valve]) (result int) {
	i := valves.FindIndex(func(v Valve) bool { return v.Name == current.Name })
	for index, candidate := range candidates {
		j := valves.FindIndex(func(v Valve) bool { return v.Name == candidate.Name })
		minutesLeft := minutes - int(distance[i][j]) - 1
		if minutesLeft >= 0 {
			fmt.Println(candidates, index)
			candidates.RemoveIndex(index)
			result = int(math.Max(float64(result), float64(valves.List[j].FlowRate*minutesLeft+ProboscideaVolcanium(valves, distance, minutesLeft, valves.List[j], candidates))))
		}
	}
	return result
}

func step1(input string) int {
	valves := MakeValves(input)
	adjacencyMatrix := valves.GetAdjacencyMatrix()
	distance := FloydWarshall(adjacencyMatrix)
	return memorized(ProboscideaVolcanium)(valves, distance, 30, valves.List[0], valves.Filter(func(valve Valve) bool {
		return valve.FlowRate > 0
	}))
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 16: Proboscidea Volcanium ---", "2022/day16/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 1651, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
