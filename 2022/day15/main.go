package main

import (
	"adventofcode/2022/day15/structs"
	structs2 "adventofcode/structs"
	"adventofcode/utils"
	"fmt"
)

func step1(input string, X int) int {
	sensors, minY, maxY := structs.SensorList{}.Decode(input)
	mapped := map[structs2.Position]bool{}
	for y := minY; y <= maxY; y++ {
		candidate := structs2.Position{X: X, Y: y}
		for _, sensor := range sensors {
			if sensor.Position.Distance(candidate) <= sensor.DistanceWithClosestBeacon {
				mapped[candidate] = true
			}
		}
	}
	for _, sensor := range sensors {
		delete(mapped, sensor.ClosestBeaconPosition)
	}
	return len(mapped)
}

func isValid(point structs2.Position, sensors []structs.Sensor) bool {
	for _, sensor := range sensors {
		if point.Distance(sensor.Position) <= sensor.DistanceWithClosestBeacon {
			return false
		}
	}
	return true
}

func step2(input string, maxX, maxY int) int {
	sensors, _, _ := structs.SensorList{}.Decode(input)
	for _, sensor := range sensors {
		for _, point := range sensor.GetExternalBorder() {
			if point.X < 0 || point.X > maxX || point.Y < 0 || point.Y > maxY {
				continue
			}
			if isValid(point, sensors) {
				return point.Y*4000000 + point.X
			}
		}
	}
	return 0
}

func main() {
	const title, day = "--- Day 15: Beacon Exclusion Zone ---", "2022/day15/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example, 10), 26, "example step1")
	utils.AssertEqual(step2(example, 20, 20), 56000011, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input, 2000000), 5403290, "step1") // 4420072
	utils.AssertEqual(step2(input, 4000000, 4000000), 10291582906626, "step2")
}
