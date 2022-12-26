package main

import (
	"adventofcode/operators"
	. "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
	"math"
)

func reverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i := range result {
		result[i] = s[len(s)-i-1]
	}
	return result
}

func getEdges(matrix Matrix[int], i int, j int) [][]int {
	return [][]int{reverse(matrix.GetRow(i)[:j]), matrix.GetRow(i)[j+1:],
		reverse(matrix.GetColumn(j)[:i]), matrix.GetColumn(j)[i+1:]}
}

func IsTreeVisible(matrix Matrix[int], i int, j int) bool {
	return operators.All(getEdges(matrix, i, j), func(edge []int) bool {
		return len(edge) > 0 && operators.Any(edge, func(point int) bool {
			return point >= matrix[i][j]
		})
	})
}

func ScenicScore(matrix Matrix[int], i int, j int) int {
	return operators.Multiply(operators.Map(getEdges(matrix, i, j), func(edge []int) int {
		index := operators.FindIndex(edge, func(point int) bool {
			return point >= matrix[i][j]
		})
		if index >= 0 {
			return index + 1
		} else {
			return len(edge)
		}
	}))
}

func step1(input string) (value int) {
	matrix := Matrix[int]{}
	matrix.Decode(input, "\n", "")
	for i, row := range matrix {
		for j := range row {
			if !IsTreeVisible(matrix, i, j) {
				value++
			}
		}
	}
	return value
}

func step2(input string) (value int) {
	matrix := Matrix[int]{}
	matrix.Decode(input, "\n", "")
	for i, row := range matrix {
		for j := range row {
			value = int(math.Max(float64(value), float64(ScenicScore(matrix, i, j))))
		}
	}
	return value
}

func main() {
	const title, day = "--- Day 8: Treetop Tree House ---", "2022/day08/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 21, "example step1")
	utils.AssertEqual(step2(example), 8, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 1801, "step1")
	utils.AssertEqual(step2(input), 209880, "step2")
}
