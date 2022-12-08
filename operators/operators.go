package operators

import (
	"math"
)

func ForEach[T any](list []T, f func(T)) {
	for index := range list {
		f(list[index])
	}
}

func Map[T, U any](list []T, f func(T) U) []U {
	mapped := make([]U, len(list))
	for index, value := range list {
		mapped[index] = f(value)
	}
	return mapped
}

func Reduce[T, U any](acc []T, f func(U, T) U, initValue U) U {
	reduced := initValue
	for _, value := range acc {
		reduced = f(reduced, value)
	}
	return reduced
}

func Filter[T any](list []T, f func(T) bool) []T {
	return Reduce(list, func(acc []T, current T) []T {
		if f(current) {
			return append(acc, current)
		} else {
			return acc
		}
	}, []T{})
}

func FindIndex[T any](list []T, f func(T) bool) int {
	for index, value := range list {
		if f(value) {
			return index
		}
	}
	return -1
}

func Any[T any](list []T, f func(T) bool) bool {
	return FindIndex(list, f) >= 0
}

func All[T any](list []T, f func(T) bool) bool {
	return len(Filter(list, f)) == len(list)
}

func Max(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return int(math.Max(float64(acc), float64(current)))
	}, math.MinInt)
}

func Min(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return int(math.Min(float64(acc), float64(current)))
	}, math.MaxInt)
}

func Sum(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return acc + current
	}, 0)
}

func Multiply(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return acc * current
	}, 1)
}

func Chunk[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T
	for {
		if len(slice) == 0 {
			break
		}
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}
		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}
	return chunks
}

func Intersection[T comparable](list1, list2 []T) (sharedElements []T) {
	m := make(map[T]bool)
	ForEach(list1, func(value T) {
		m[value] = true
	})
	ForEach(list2, func(value T) {
		if m[value] {
			sharedElements = append(sharedElements, value)
		}
	})
	return sharedElements
}

type numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Add[T numeric](list1, list2 []T) []T {
	sum := make([]T, len(list1))
	for index, _ := range list1 {
		sum[index] = list1[index] + list2[index]
	}
	return sum
}

func HasDuplicates[T comparable](arr []T) bool {
	visited := make(map[T]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return true
		} else {
			visited[arr[i]] = true
		}
	}
	return false
}
