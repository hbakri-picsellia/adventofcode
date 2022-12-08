package main

import (
	"adventofcode/2022/day07/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func parseCommands(input string) []models.Command {
	inputCommands := strings.Split(input, "$")[1:]
	return operators.Map(inputCommands, func(commandInput string) (command models.Command) {
		command.Decode(strings.TrimSpace(commandInput))
		return command
	})
}

func parseTree(input string) *models.Folder {
	commands := parseCommands(input)[1:]
	currentFolder := &models.Folder{Name: "/"}
	operators.ForEach(commands, func(command models.Command) {
		switch command.Type {
		case models.CD:
			if command.Arg == ".." {
				currentFolder = currentFolder.Parent
			} else {
				currentFolder = currentFolder.GetChild(command.Arg)
			}
		case models.LS:
			currentFolder.DecodeChildren(command.Result)
		}
	})
	currentFolder = currentFolder.GetRoot()
	return currentFolder
}

func step1(input string) int {
	root := parseTree(input)
	return root.Step1()
}

func step2(input string) int {
	root := parseTree(input)
	rootSize := root.GetSize()
	targetSize := 30000000 - (70000000 - rootSize)
	return root.Step2(targetSize)
}

func main() {
	const title, day = "--- Day 7: No Space Left On Device ---", "2022/day07/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 95437, "example step1")
	utils.AssertEqual(step2(example), 24933642, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 1443806, "step1")
	utils.AssertEqual(step2(input), 942298, "step2")
}
