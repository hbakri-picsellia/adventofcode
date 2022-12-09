package main

import (
	"adventofcode/2022/day09/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

func parseCommands(input string) []models.Command {
	inputCommands := strings.Split(input, "\n")
	return operators.Map(inputCommands, func(commandInput string) (command models.Command) {
		command.Decode(strings.TrimSpace(commandInput))
		return command
	})
}

func step1(input string) int {
	commands := parseCommands(input)
	HeadPosition := models.Point{}
	TailPosition := models.Point{}
	results := make(map[string]int, 0)
	for _, command := range commands {
		for step := 0; step < command.Number; step++ {
			switch command.Type {
			case models.Up:
				HeadPosition.X += 1
			case models.Down:
				HeadPosition.X -= 1
			case models.Right:
				HeadPosition.Y += 1
			case models.Left:
				HeadPosition.Y -= 1
			}

			if !TailPosition.IsClose(HeadPosition) {
				if TailPosition.X == HeadPosition.X {
					directionY := (HeadPosition.Y - TailPosition.Y) / int(math.Abs(float64(TailPosition.Y-HeadPosition.Y)))
					TailPosition.Y += directionY
				} else if TailPosition.Y == HeadPosition.Y {
					directionX := (HeadPosition.X - TailPosition.X) / int(math.Abs(float64(TailPosition.X-HeadPosition.X)))
					TailPosition.X += directionX
				} else {
					directionX := (HeadPosition.X - TailPosition.X) / int(math.Abs(float64(TailPosition.X-HeadPosition.X)))
					directionY := (HeadPosition.Y - TailPosition.Y) / int(math.Abs(float64(TailPosition.Y-HeadPosition.Y)))
					TailPosition.Y += directionY
					TailPosition.X += directionX
				}
			}
			results[TailPosition.ToString()] += 1
		}
	}
	return len(results)
}

func nextStep(head, tail *models.Point) {
	if !tail.IsClose(*head) {
		if tail.X == head.X {
			directionY := (head.Y - tail.Y) / int(math.Abs(float64(tail.Y-head.Y)))
			tail.Y += directionY
		} else if tail.Y == head.Y {
			directionX := (head.X - tail.X) / int(math.Abs(float64(tail.X-head.X)))
			tail.X += directionX
		} else {
			directionX := (head.X - tail.X) / int(math.Abs(float64(tail.X-head.X)))
			directionY := (head.Y - tail.Y) / int(math.Abs(float64(tail.Y-head.Y)))
			tail.Y += directionY
			tail.X += directionX
		}
	}
}

func step2(input string) int {
	commands := parseCommands(input)
	HeadPosition := models.Point{}
	SnakePositions := make([]models.Point, 9)
	results := make(map[string]int, 0)
	for _, command := range commands {
		for step := 0; step < command.Number; step++ {
			switch command.Type {
			case models.Up:
				HeadPosition.X += 1
			case models.Down:
				HeadPosition.X -= 1
			case models.Right:
				HeadPosition.Y += 1
			case models.Left:
				HeadPosition.Y -= 1
			}
			nextStep(&HeadPosition, &SnakePositions[0])
			for index := range SnakePositions {
				if index < len(SnakePositions)-1 {
					nextStep(&SnakePositions[index], &SnakePositions[index+1])
				}
			}
			results[SnakePositions[8].ToString()] += 1
		}
	}
	return len(results)
}

func main() {
	const title, day = "--- Day 9: Rope Bridge ---", "2022/day09/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 13, "example step1")
	utils.AssertEqual(step2(example), 1, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 6209, "step1")
	utils.AssertEqual(step2(input), 2460, "step2")
}
