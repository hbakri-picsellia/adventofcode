package main

import (
	"adventofcode/models"
	"adventofcode/utils"
	"fmt"
)

type Node struct {
	position  models.Position
	Parent    *Node
	cost      int
	heuristic float64
}

func (node *Node) getHeight(matrix models.Matrix[rune]) int {
	return getLetterHeight(matrix[node.position.I][node.position.J])
}

func getLetterHeight(char rune) int {
	if char == 'E' {
		return getLetterHeight('z')
	} else if char == 'S' {
		return getLetterHeight('a')
	} else {
		return int(char - 'a' + 1)
	}
}

func AStar(matrix models.Matrix[rune], startNode, endNode Node) *Node {
	queue := models.Stack[Node]{startNode}
	visited := map[models.Position]bool{}
	for !queue.IsEmpty() {
		currentNode, _ := queue.Shift()
		visited[currentNode.position] = true

		if currentNode.position.Equals(endNode.position) {
			return &currentNode
		}

		currentHeight := currentNode.getHeight(matrix)
		for _, direction := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			neighbor := models.Position{I: currentNode.position.I + direction[0], J: currentNode.position.J + direction[1]}
			neighborNode := Node{position: neighbor, Parent: &currentNode}
			if matrix.Contains(neighbor) && !visited[neighborNode.position] && neighborNode.getHeight(matrix)-currentHeight <= 1 {
				neighborNode.cost = currentNode.cost + 1
				queue.Push(neighborNode)
			}
		}
	}
	return nil
}

func step1(input string) int {
	matrix := models.Matrix[rune]{}
	matrix.Decode(input, "\n", "", func(s string) rune { return []rune(s)[0] })
	startNode := Node{position: matrix.Find(func(s rune) bool { return s == 'S' })}
	endNode := Node{position: matrix.Find(func(s rune) bool { return s == 'E' })}
	result := AStar(matrix, startNode, endNode)
	return result.cost
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 12: Hill Climbing Algorithm ---", "2022/day12/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 31, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
