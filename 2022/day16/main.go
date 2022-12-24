package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
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
	minutes        float64
	currentName    string
	candidatesName string
	isElephant     bool
}

type Func func(valves Valves, distance Matrix[float64], minutes float64, current Valve, candidates List[Valve], isElephant bool) float64

var cache = make(map[Params]float64)

func memorized(fn Func) Func {
	return func(valves Valves, distance Matrix[float64], minutes float64, current Valve, candidates List[Valve], isElephant bool) float64 {
		names := operators.Map(candidates, func(valve Valve) string {
			return valve.Name
		})
		sort.Strings(names)
		input := Params{minutes: minutes, currentName: current.Name, candidatesName: strings.Join(names, ","), isElephant: isElephant}
		if val, found := cache[input]; found {
			return val
		}

		result := fn(valves, distance, minutes, current, candidates, isElephant)
		cache[input] = result
		return result
	}
}

func ProboscideaVolcanium(valves Valves, distance Matrix[float64], minutes float64, current Valve, candidates List[Valve], isElephant bool) (result float64) {
	if isElephant {
		initialValve := *valves.Find(func(v Valve) bool { return v.Name == "AA" })
		result = memorized(ProboscideaVolcanium)(valves, distance, 26, initialValve, candidates, false)
	}
	i := valves.FindIndex(func(v Valve) bool { return v.Name == current.Name })
	for index, candidate := range candidates {
		j := valves.FindIndex(func(v Valve) bool { return v.Name == candidate.Name })
		if nextMinutes := minutes - distance[i][j] - 1; nextMinutes >= 0 {
			newCandidates := candidates.Clone()
			newCandidates.RemoveIndex(index)
			result = math.Max(result, candidate.FlowRate*nextMinutes+memorized(ProboscideaVolcanium)(valves, distance, nextMinutes, candidate, newCandidates, isElephant))
		}
	}
	return result
}

func step1(input string) int {
	valves := MakeValves(input)
	adjacencyMatrix := valves.GetAdjacencyMatrix()
	distance := FloydWarshall(adjacencyMatrix)
	initialValve := *valves.Find(func(v Valve) bool { return v.Name == "AA" })
	initialCandidates := valves.Filter(func(valve Valve) bool { return valve.FlowRate > 0 })
	return int(memorized(ProboscideaVolcanium)(valves, distance, 30, initialValve, initialCandidates, false))
}

func step2(input string) int {
	valves := MakeValves(input)
	adjacencyMatrix := valves.GetAdjacencyMatrix()
	distance := FloydWarshall(adjacencyMatrix)
	initialValve := *valves.Find(func(v Valve) bool { return v.Name == "AA" })
	initialCandidates := valves.Filter(func(valve Valve) bool { return valve.FlowRate > 0 })
	return int(memorized(ProboscideaVolcanium)(valves, distance, 26, initialValve, initialCandidates, true))
}

func main() {
	const title, day = "--- Day 16: Proboscidea Volcanium ---", "2022/day16/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 1651, "example step1")
	utils.AssertEqual(step2(example), 1707, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 1991, "step1")
	utils.AssertEqual(step2(input), 2705, "step2")
}
