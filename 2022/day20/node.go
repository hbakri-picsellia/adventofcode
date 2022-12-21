package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"strings"
)

type Node struct {
	id     int
	number int
	left   *Node
	right  *Node
}

type Nodes struct {
	List[Node]
}

func MakeNodes(s string, decryptionKey int) (nodes Nodes) {
	numbers := operators.Map(strings.Split(s, "\n"), utils.ParseStringToInt)
	for index, number := range numbers {
		nodes.Push(Node{id: index, number: decryptionKey * number})
	}
	for index := range nodes.List {
		nodes.List[index].left = &nodes.List[(nodes.Len()+index-1)%nodes.Len()]
		nodes.List[index].right = &nodes.List[(index+1)%nodes.Len()]
	}
	return
}

func (nodes *Nodes) MoveRight(index, nbMoves int) {
	if nbMoves == 0 {
		return
	}

	sourceNode := &nodes.List[index]
	sourceNode.right.left = sourceNode.left
	sourceNode.left.right = sourceNode.right

	destinationNode := &nodes.List[index]
	for i := 0; i < nbMoves; i++ {
		destinationNode = destinationNode.right
	}
	destinationNode.right.left = sourceNode
	sourceNode.right = destinationNode.right
	destinationNode.right = sourceNode
	sourceNode.left = destinationNode
}

func (nodes *Nodes) MoveLeft(index, nbMoves int) {
	if nbMoves == 0 {
		return
	}

	sourceNode := &nodes.List[index]
	sourceNode.right.left = sourceNode.left
	sourceNode.left.right = sourceNode.right

	destinationNode := &nodes.List[index]
	for i := 0; i < nbMoves; i++ {
		destinationNode = destinationNode.left
	}
	destinationNode.left.right = sourceNode
	sourceNode.left = destinationNode.left
	destinationNode.left = sourceNode
	sourceNode.right = destinationNode
}

func (nodes *Nodes) GetGroveCoordinates() (result int) {
	zero := nodes.Find(func(node Node) bool {
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
