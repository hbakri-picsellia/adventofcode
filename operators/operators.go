package operators

import "math"

func ForEach[T any](list []T, f func(int, T)) {
	for index, value := range list {
		f(index, value)
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

func Max(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return int(math.Max(float64(acc), float64(current)))
	}, 0)
}

func Sum(list []int) int {
	return Reduce(list, func(acc, current int) int {
		return acc + current
	}, 0)
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

func Intersection[T comparable](list1 []T, list2 []T) (sharedElements []T) {
	m := make(map[T]bool)
	ForEach(list1, func(_ int, value T) {
		m[value] = true
	})
	ForEach(list2, func(_ int, value T) {
		if m[value] {
			sharedElements = append(sharedElements, value)
		}
	})
	return sharedElements
}
