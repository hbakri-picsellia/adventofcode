package models

import "strings"

type Stack struct {
	title   int
	content []string
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.content) == 0
}

func (stack *Stack) Push(str string) {
	stack.content = append(stack.content, str) // Simply append the new value to the end of the stack
}

func (stack *Stack) Pop() (string, bool) {
	if stack.IsEmpty() {
		return "", false
	} else {
		index := len(stack.content) - 1         // Get the index of the top most element.
		element := (stack.content)[index]       // Index into the slice and obtain the element.
		stack.content = (stack.content)[:index] // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (stack *Stack) Decode(input string) {
	list := strings.Split(input, "\n")
	for i := len(list) - 1; i >= 0; i-- {
		if i == len(list)-1 {

		} else {

		}
	}
}
