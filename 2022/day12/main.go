package main

import (
	"adventofcode/models"
	"adventofcode/utils"
	"fmt"
)

type Node struct {
	position models.Position
	cost     int
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

func HillClimbingAlgorithm(input string, start, end rune, metric func(int) bool) int {
	matrix := models.Matrix[rune]{}
	matrix.Decode(input, "\n", "", func(s string) rune { return []rune(s)[0] })
	startNode := Node{position: matrix.Find(func(s rune) bool { return s == start })}

	queue := models.Stack[Node]{startNode}
	visited := map[models.Position]bool{}
	for !queue.IsEmpty() {
		currentNode, _ := queue.Shift()
		if visited[currentNode.position] {
			continue
		}
		visited[currentNode.position] = true

		if matrix[currentNode.position.I][currentNode.position.J] == end {
			return currentNode.cost
		}

		for _, direction := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			neighbor := models.Position{I: currentNode.position.I + direction[0], J: currentNode.position.J + direction[1]}
			neighborNode := Node{position: neighbor, cost: currentNode.cost + 1}

			if matrix.Contains(neighbor) {
				distance := neighborNode.getHeight(matrix) - currentNode.getHeight(matrix)
				if metric(distance) {
					queue.Push(neighborNode)
				}
			}
		}
	}
	return -1
}

func step1(input string) int {
	return HillClimbingAlgorithm(input, 'S', 'E', func(distance int) bool {
		return distance <= 1
	})
}

func step2(input string) int {
	return HillClimbingAlgorithm(input, 'E', 'a', func(distance int) bool {
		return distance >= -1
	})
}

func main() {
	const title, day = "--- Day 12: Hill Climbing Algorithm ---", "2022/day12/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 31, "example step1")
	utils.AssertEqual(step2(example), 29, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 361, "step1")
	utils.AssertEqual(step2(input), 354, "step2")
}
