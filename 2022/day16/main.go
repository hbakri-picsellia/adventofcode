package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
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

func ProboscideaVolcanium(valves Valves, distance Matrix[float64], currentValve Valve, minutes int, candidatesName map[string]bool) (result int) {
	i := valves.FindIndex(func(v Valve) bool { return v.Name == currentValve.Name })
	for candidateName := range candidatesName {
		j := valves.FindIndex(func(v Valve) bool { return v.Name == candidateName })
		cost := minutes - int(distance[i][j]) - 1
		if cost >= 0 {
			delete(candidatesName, candidateName)
			result = int(math.Max(float64(result), float64(valves.List[j].FlowRate*cost+ProboscideaVolcanium(valves, distance, valves.List[j], cost, candidatesName))))
		}
	}
	return result
}

func step1(input string) int {
	valves := MakeValves(input)
	adjacencyMatrix := valves.GetAdjacencyMatrix()
	distance := FloydWarshall(adjacencyMatrix)

	candidatesMap := make(map[string]bool)
	candidatesName := operators.Map(valves.Filter(func(valve Valve) bool {
		return valve.FlowRate > 0
	}), func(valve Valve) string { return valve.Name })
	for _, value := range candidatesName {
		candidatesMap[value] = true
	}
	return ProboscideaVolcanium(valves, distance, valves.List[0], 30, candidatesMap)
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 16: Proboscidea Volcanium ---", "2022/day16/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
