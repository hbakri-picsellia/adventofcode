package structs

import (
	"adventofcode/operators"
	"strconv"
)

type List[T any] []T

func (list *List[T]) Len() int {
	return len(*list)
}

func (list *List[T]) IsEmpty() bool {
	return len(*list) == 0
}

func (list *List[T]) Push(newElements ...T) {
	*list = append(*list, newElements...)
}

func (list *List[T]) Insert(index int, element T) {
	if index == 0 {
		*list = append([]T{element}, *list...)
	} else if index == len(*list) {
		list.Push(element)
	} else if index > 0 && index < len(*list) {
		*list = append((*list)[:index+1], (*list)[index:]...)
		(*list)[index] = element
	} else {
		panic("index " + strconv.FormatInt(int64(index), 10) + " out of range " + strconv.FormatInt(int64(len(*list)), 10))
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

func (list *List[T]) Shift() (result T, exists bool) {
	if list.IsEmpty() {
		return result, false
	} else {
		element := (*list)[0]
		*list = (*list)[1:]
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

func (list *List[T]) All(f func(T) bool) bool {
	return len(list.Filter(f)) == len(*list)
}

func (list *List[T]) Clone() List[T] {
	clone := make([]T, len(*list))
	copy(clone, *list)
	return clone
}

func (list *List[T]) Reduce(f func(T, T) T, initValue T) T {
	reduced := initValue
	for _, value := range *list {
		reduced = f(reduced, value)
	}
	return reduced
}

func (list *List[T]) Map(f func(T) T) []T {
	result := make([]T, len(*list))
	for index, value := range *list {
		result[index] = f(value)
	}
	return result
}

func (list *List[T]) Reverse() []T {
	result := make([]T, len(*list))
	for index := range result {
		result[index] = (*list)[len(*list)-index-1]
	}
	return result
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

func (list *ListComparable[T]) Intersection(list2 ListComparable[T]) (sharedElements []T) {
	m := make(map[T]bool)
	list.ForEach(func(value T) {
		m[value] = true
	})
	list2.ForEach(func(value T) {
		if m[value] {
			sharedElements = append(sharedElements, value)
		}
	})
	return sharedElements
}

func (list *ListComparable[T]) Intersects(list2 ListComparable[T]) bool {
	m := make(map[T]bool)
	list.ForEach(func(value T) {
		m[value] = true
	})
	for _, value := range list2.List {
		if m[value] {
			return true
		}
	}
	return false
}
