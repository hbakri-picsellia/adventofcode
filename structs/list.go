package structs

import (
	"adventofcode/operators"
)

type List[T any] []T

func (list *List[T]) IsEmpty() bool {
	return len(*list) == 0
}

func (list *List[T]) Push(newElement T) {
	*list = append(*list, newElement)
}

func (list *List[T]) Shift() (result T, exists bool) {
	if list.IsEmpty() {
		return result, false
	} else {
		element := (*list)[0]
		*list = (*list)[1:]
		return element, true
	}
}

func (list *List[T]) Pop() (result T, exists bool) {
	if list.IsEmpty() {
		return result, false
	} else {
		index := len(*list) - 1
		element := (*list)[index]
		*list = (*list)[:index]
		return element, true
	}
}

func (list *List[T]) RemoveIndex(index int) {
	*list = append((*list)[:index], (*list)[index+1:]...)
}

func (list *List[T]) Find(f func(T) bool) *T {
	for _, value := range *list {
		if f(value) {
			return &value
		}
	}
	return nil
}

func (list *List[T]) FindIndex(f func(T) bool) int {
	for index, value := range *list {
		if f(value) {
			return index
		}
	}
	return -1
}

func (list *List[T]) ForEach(f func(T)) {
	for index := range *list {
		f((*list)[index])
	}
}

func (list *List[T]) Filter(f func(T) bool) []T {
	return operators.Reduce(*list, func(acc []T, current T) []T {
		if f(current) {
			return append(acc, current)
		} else {
			return acc
		}
	}, []T{})
}

func (list *List[T]) Clone() List[T] {
	clone := make([]T, len(*list))
	copy(clone, *list)
	return clone
}

type ListComparable[T comparable] struct {
	List[T]
}

func (list *ListComparable[T]) Equals(list2 ListComparable[T]) bool {
	if len((*list).List) != len(list2.List) {
		return false
	}
	for index := range (*list).List {
		if (*list).List[index] != list2.List[index] {
			return false
		}
	}
	return true
}

func (list *ListComparable[T]) Contains(element T) bool {
	for _, value := range list.List {
		if value == element {
			return true
		}
	}
	return false
}

func (list *ListComparable[T]) Remove(value T) {
	index := list.FindIndex(func(v T) bool { return v == value })
	list.RemoveIndex(index)
}
