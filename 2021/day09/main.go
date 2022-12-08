package main

import (
	"adventofcode/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"sort"
)

func reverse[T any](s []T) []T {
	result := make([]T, len(s))
	for i := range result {
		result[i] = s[len(s)-i-1]
	}
	return result
}

func getEdges(matrix models.Matrix, i, j int) [][]int {
	return [][]int{reverse(matrix.GetRow(i)[:j]), matrix.GetRow(i)[j+1:],
		reverse(matrix.GetColumn(j)[:i]), matrix.GetColumn(j)[i+1:]}
}

func isLowPoint(matrix models.Matrix, i, j int) bool {
	return operators.All(getEdges(matrix, i, j), func(edge []int) bool {
		return len(edge) == 0 || edge[0] > matrix[i][j]
	})
}

func recursiveBasinSearch(matrix *models.Matrix, i, j int) (value int) {
	n, m := matrix.GetDimension()
	if i < 0 || i >= n || j < 0 || j >= m {
		return 0
	} else if (*matrix)[i][j] == 9 {
		return 0
	} else {
		value += 1
		(*matrix)[i][j] = 9
	}
	value += recursiveBasinSearch(matrix, i+1, j)
	value += recursiveBasinSearch(matrix, i-1, j)
	value += recursiveBasinSearch(matrix, i, j+1)
	value += recursiveBasinSearch(matrix, i, j-1)
	return value
}

func step1(input string) (value int) {
	matrix := models.Matrix{}
	matrix.Decode(input, "\n", "")
	for i, row := range matrix {
		for j := range row {
			if isLowPoint(matrix, i, j) {
				value += matrix[i][j] + 1
			}
		}
	}
	return value
}

func step2(input string) (value int) {
	matrix := models.Matrix{}
	matrix.Decode(input, "\n", "")
	results := make([]int, 0)
	for i, row := range matrix {
		for j := range row {
			if isLowPoint(matrix, i, j) {
				results = append(results, recursiveBasinSearch(&matrix, i, j))
			}
		}
	}
	sort.Ints(results)
	return operators.Multiply(results[len(results)-3:])
}

func main() {
	const title, day = "--- Day 9: Smoke Basin ---", "2021/day09/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 15, "example step1")
	utils.AssertEqual(step2(example), 1134, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 572, "step1")
	utils.AssertEqual(step2(input), -1, "step2")
}
