package utils

import (
	"os"
	"sort"
)

func ParseTxtFile(filename string) string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(fileContent)
}

func Max(list []int) (max int) {
	max = list[0]
	for _, value := range list {
		if value > max {
			max = value
		}
	}
	return max
}

func MaxN(list []int, n int) (maxN []int) {
	sort.Ints(list)
	return list[len(list)-n:]
}

func Sum(list []int) (sum int) {
	return Reduce(list, func(acc, current int) int {
		return acc + current
	}, 0)
}

func Map[T, U any](list []T, f func(T) U) []U {
	mapped := make([]U, len(list))
	for index, value := range list {
		mapped[index] = f(value)
	}
	return mapped
}

func Reduce[T, U any](list []T, f func(U, T) U, initValue U) U {
	reduced := initValue
	for _, value := range list {
		reduced = f(reduced, value)
	}
	return reduced
}
