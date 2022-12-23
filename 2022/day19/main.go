package main

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strings"
)

func step1(input string) int {
	for _, rawBlueprint := range strings.Split(input, "\n") {
		r := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

		fmt.Println(r.FindAllString(rawBlueprint, -1))
	}
	return 0
}

func step2(input string) int {
	return 0
}

func main() {
	const title, day = "--- Day 19: Not Enough Minerals ---", "2022/day19/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), -1, "example step1")
	//utils.AssertEqual(step2(example), -1, "example step2")
	//
	//input := utils.ParseFileToString(day + "input.txt")
	//utils.AssertEqual(step1(input), -1, "step1")
	//utils.AssertEqual(step2(input), -1, "step2")
}
