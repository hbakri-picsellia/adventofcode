package utils

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseFileToString(filename string) string {
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

func FirstN(list []int, n int) []int {
	return list[:len(list)-n]
}

func LastN(list []int, n int) []int {
	return list[len(list)-n:]
}

func MaxN(list []int, n int) []int {
	sort.Ints(list)
	return LastN(list, n)
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

func ParseStringToInt(s string) int {
	num, _ := strconv.ParseInt(s, 10, 0)
	return int(num)
}

func ParseStringToIntList(text string, separator string) []int {
	return Map(strings.Split(text, separator), ParseStringToInt)
}

func AssertEqual[T comparable](input1 T, input2 T, message string) {
	if input1 == input2 {
		fmt.Println(message, "succeeded !")
	} else {
		fmt.Println(message, "failed...", "output:", input1, "expected:", input2)
	}
}
