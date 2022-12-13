package models

type Stack[T comparable] []T

func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack[T]) Push(newElement T) {
	*stack = append(*stack, newElement)
}

func (stack *Stack[T]) Shift() (result T, exists bool) {
	if stack.IsEmpty() {
		return result, false
	} else {
		element := (*stack)[0]
		*stack = (*stack)[1:]
		return element, true
	}
}

func (stack *Stack[T]) Pop() (result T, exists bool) {
	if stack.IsEmpty() {
		return result, false
	} else {
		index := len(*stack) - 1
		element := (*stack)[index]
		*stack = (*stack)[:index]
		return element, true
	}
}

func (stack *Stack[T]) RemoveIndex(index int) {
	*stack = append((*stack)[:index], (*stack)[index+1:]...)
}

func (stack *Stack[T]) Find(f func(T) bool) int {
	for index, value := range *stack {
		if f(value) {
			return index
		}
	}
	return -1
}

func (stack *Stack[T]) Remove(value T) {
	index := stack.Find(func(v T) bool { return v == value })
	stack.RemoveIndex(index)
}
