package main

import (
	"adventofcode/2022/day11/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func parseMonkeys(input string) []models.Monkey {
	return operators.Map(strings.Split(input, "\n\n"), func(monkeyInput string) (monkey models.Monkey) {
		monkey.Decode(monkeyInput)
		return monkey
	})
}

func step1(input string) int {
	monkeys := parseMonkeys(input)
	for round := 0; round < 20; round++ {
		for index := range monkeys {
			for {
				if monkeys[index].Items.IsEmpty() {
					break
				}
				currentItem, _ := monkeys[index].Items.Shift()
				currentItem = monkeys[index].Operation(currentItem)
				currentItem = int(math.Floor(float64(currentItem) / float64(3)))
				if monkeys[index].Test(currentItem) {
					monkeys[monkeys[index].SuccessDestination].Items.Push(currentItem)
				} else {
					monkeys[monkeys[index].FailureDestination].Items.Push(currentItem)
				}
				monkeys[index].NbItemsInspected += 1
			}
		}
	}
	nbItemsInspected := operators.Map(monkeys, func(monkey models.Monkey) int {
		return monkey.NbItemsInspected
	})
	sort.Ints(nbItemsInspected)
	return nbItemsInspected[len(nbItemsInspected)-1] * nbItemsInspected[len(nbItemsInspected)-2]
}

func step2(input string) int {
	monkeys := parseMonkeys(input)
	for round := 0; round < 10000; round++ {
		for index := range monkeys {
			monkeys[index].NbItemsInspected += len(monkeys[index].Items)
			for {
				if monkeys[index].Items.IsEmpty() {
					break
				}
				currentItem, _ := monkeys[index].Items.Shift()
				currentItem = monkeys[index].Operation(currentItem)
				currentItem %= operators.Multiply(operators.Map(monkeys, func(monkey models.Monkey) int {
					return monkey.TestDivider
				}))
				if monkeys[index].Test(currentItem) {
					monkeys[monkeys[index].SuccessDestination].Items.Push(currentItem)
				} else {
					monkeys[monkeys[index].FailureDestination].Items.Push(currentItem)
				}
			}
		}
	}
	nbItemsInspected := operators.Map(monkeys, func(monkey models.Monkey) int {
		return monkey.NbItemsInspected
	})
	sort.Ints(nbItemsInspected)
	return nbItemsInspected[len(nbItemsInspected)-1] * nbItemsInspected[len(nbItemsInspected)-2]
}

func main() {
	const title, day = "--- Day 11: Monkey in the Middle ---", "2022/day11/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), 10605, "example step1")
	utils.AssertEqual(step2(example), 2713310158, "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), 120756, "step1")
	utils.AssertEqual(step2(input), 39109444654, "step2")
}
