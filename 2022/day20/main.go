package main

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func step1(input string) (result int) {
	numbers := operators.Map(strings.Split(input, "\n"), utils.ParseStringToInt)
	nodes := Nodes{List: operators.Map(numbers, func(number int) Node {
		return Node{number: number}
	})}
	for index := range nodes.List {
		nodes.List[index].left = &nodes.List[(nodes.Len()+index-1)%nodes.Len()]
		nodes.List[index].right = &nodes.List[(index+1)%nodes.Len()]
	}

	fmt.Println(nodes)
	for _, node := range nodes.List {
		if node.number == 0 {
			continue
		}
		currentNode := node

		if node.number > 0 {
			nodes.MoveRight()
		} else {
			//for i := 0; i > -nodes[index].number%(len(nodes)-1); i-- {
			//	currentNode = currentNode.left
			//}
			//if currentNode == nodes[index] {
			//	continue
			//}
			//nodes[index].left.right = nodes[index].right
			//nodes[index].right.left = nodes[index].left
			//currentNode.left.right = nodes[index]
			//nodes[index].left = currentNode.left
			//currentNode.left = nodes[index]
			//nodes[index].right = currentNode
		}
	}
	fmt.Println(nodes)

	zero := *nodes.Find(func(node *Node) bool {
		return node.number == 0
	})
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			zero = zero.right
		}
		result += zero.number
	}
	return result
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 20: Grove Positioning System ---", "2022/day20/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 3, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")

	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
