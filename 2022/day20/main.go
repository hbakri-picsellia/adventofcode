package main

import (
	"adventofcode/utils"
	"fmt"
)

func GrovePositioningSystem(nodes *Nodes) {
	for index, node := range nodes.List {
		if node.number == 0 {
			continue
		} else if node.number > 0 {
			nbMoves := nodes.List[index].number % (nodes.Len() - 1)
			nodes.MoveRight(index, nbMoves)
		} else {
			nbMoves := -nodes.List[index].number % (nodes.Len() - 1)
			nodes.MoveLeft(index, nbMoves)
		}
	}
}

func step1(input string) int {
	nodes := MakeNodes(input, 1)
	GrovePositioningSystem(&nodes)
	return nodes.GetGroveCoordinates()
}

func step2(input string) int {
	nodes := MakeNodes(input, 811589153)
	for i := 0; i < 10; i++ {
		GrovePositioningSystem(&nodes)
	}
	return nodes.GetGroveCoordinates()
}

func main() {
	const title, day = "--- Day 20: Grove Positioning System ---", "2022/day20/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 3, "example step1")
	utils.AssertEqual(step2(example), 1623178306, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 872, "step1")
	utils.AssertEqual(step2(input), 5382459262696, "step2")
}
