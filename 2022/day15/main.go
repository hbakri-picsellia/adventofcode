package main

import (
	"adventofcode/2022/day15/structs"
	"adventofcode/models"
	"adventofcode/utils"
	"fmt"
)

func step1(input string, i int) int {
	sensors, minJ, maxJ := structs.SensorList{}.Decode(input)
	mapped := map[models.Position]bool{}
	for j := minJ; j <= maxJ; j++ {
		candidate := models.Position{I: i, J: j}
		for _, sensor := range sensors {
			if sensor.Position.Distance(candidate) <= sensor.DistanceWithClosestBeacon {
				mapped[candidate] = true
			}
		}
	}
	for _, sensor := range sensors {
		delete(mapped, sensor.Position)
		delete(mapped, sensor.ClosestBeaconPosition)
	}
	return len(mapped)
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 15: Beacon Exclusion Zone ---", "2022/day15/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example, 10), 26, "example step1")
	//utils.AssertEqual(step2(example), 56000011, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input, 2000000), 5403290, "step1") // 4420072
	//utils.AssertEqual(step2(input), -1, "step2")
}
