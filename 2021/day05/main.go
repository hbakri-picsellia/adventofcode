package main

import (
	"adventofcode/2021/day05/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func getSegments(input string) []models.Segment {
	return operators.Map(strings.Split(input, "\n"), func(segmentInput string) (segment models.Segment) {
		segment.Decode(segmentInput)
		return segment
	})
}

func HydrothermalVenture(input string, f func(segment models.Segment) []models.Point) (count int) {
	mapping := make(map[models.Point]int)
	operators.ForEach(getSegments(input), func(segment models.Segment) {
		operators.ForEach(f(segment), func(point models.Point) {
			mapping[point] += 1
			if mapping[point] == 2 {
				count += 1
			}
		})
	})
	return count
}

func step1(input string) int {
	return HydrothermalVenture(input, func(segment models.Segment) []models.Point {
		return segment.GetCoveredPointsWithoutDiagonals()
	})
}

func step2(input string) int {
	return HydrothermalVenture(input, func(segment models.Segment) []models.Point {
		return segment.GetCoveredPointsWithDiagonals()
	})
}

func main() {
	const title, day = "--- Day 5: Hydrothermal Venture ---", "2021/day05/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 5, "example step1")
	utils.AssertEqual(step2(example), 12, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 6572, "step1")
	utils.AssertEqual(step2(input), 21466, "step2")
}
