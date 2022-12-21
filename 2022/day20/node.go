package main

import . "adventofcode/structs"

type Node struct {
	number int
	left   *Node
	right  *Node
}

type Nodes struct {
	List[Node]
}

func (nodes *Nodes) MoveRight(index int) {
	nbMoves := nodes.List[index].number % (nodes.Len() - 1)
	currentNode := nodes.List[index]
	if nbMoves == 0 {
		return
	}
	for i := 0; i < nbMoves; i++ {
		nodes.List[index] = *nodes.List[index].right
	}
	nodes.List[index].right.left = nodes.List[index].left
	nodes.List[index].left.right = nodes.List[index].right
	currentNode.right.left = &nodes.List[index]
	nodes.List[index].right = currentNode.right
	currentNode.right = &nodes.List[index]
	nodes.List[index].left = &currentNode
}
