package main

import (
	"adventofcode/2022/day09/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func parseCommands(input string) []models.Command {
	inputCommands := strings.Split(input, "\n")
	return operators.Map(inputCommands, func(commandInput string) (command models.Command) {
		command.Decode(strings.TrimSpace(commandInput))
		return command
	})
}

func RopeBridge(input string, ropeSize int) int {
	commands := parseCommands(input)
	knots := make([]models.Point, ropeSize)
	results := make(map[string]int, 0)
	for _, command := range commands {
		for step := 0; step < command.Number; step++ {
			knots[0].Move(command)
			for i := 0; i < len(knots)-1; i++ {
				knots[i+1].Follow(knots[i])
			}
			results[knots[len(knots)-1].ToString()] += 1
		}
	}
	return len(results)
}

func step1(input string) int {
	return RopeBridge(input, 2)
}

func step2(input string) int {
	return RopeBridge(input, 10)
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
