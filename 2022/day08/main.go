package main

import (
	"adventofcode/2022/day08/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
)

func remove(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func step1(input string) int {
	matrix := models.Matrix{}
	matrix.Decode(input, "\n", "")
	fmt.Println(matrix.GetRow(0))
	fmt.Println(matrix.GetColumn(0))
	for i, row := range matrix {
		for j := range row {
			// condition i>0 and i<nbColumns and same for j, with :
			anyRow := operators.Any(remove(row, j), func(value int) bool {
				return value >= matrix[i][j]
			})
			//anyColumn :=
			//	fmt.Println(remove(row, j), "(", i, j, ")", anyRow)
		}
	}
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 8: Treetop Tree House ---", "2022/day08/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
