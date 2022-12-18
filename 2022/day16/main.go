package main

import (
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

//def dp(i: str, t: int, remaining: frozenset):
//	ans = 0
//	for j in remaining:
//		if (next_t := t - dist[i][j] - 1) >= 0:
//			ans = max(ans, valves[j] * next_t + dp(j, next_t, remaining - {j}))
//	return ans

func step1(input string) (pressure int) {
	valves := MakeValves(input)
	adjacencyMatrix := valves.GetAdjacencyMatrix()
	distance := FloydWarshall(adjacencyMatrix)

	currentValveIndex := 0
	minutes := 0
	for minutes < 30 {

		// looking for the best choice
		bestChoice := math.Inf(-1)
		var bestChoiceIndex int
		for index := range valves.List {
			if !valves.List[index].Open {
				choice := (float64(minutes) - distance[currentValveIndex][index]) * float64(valves.List[index].FlowRate)
				if choice > bestChoice {
					bestChoice = choice
					bestChoiceIndex = index
				}
			}
		}

		// go to the best choice
		for i := 0; i < int(distance[currentValveIndex][bestChoiceIndex])-1; i++ {
			minutes++
			pressure += valves.GetPressure()
		}
		minutes++
		valves.List[bestChoiceIndex].Open = true
		fmt.Println(valves.List[currentValveIndex].Name, valves.List[bestChoiceIndex].Name)

		// select the best choice as the current
		currentValveIndex = bestChoiceIndex
	}
	return pressure
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
