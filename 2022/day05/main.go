package main

import (
	"adventofcode/2022/day05/models"
	"adventofcode/operators"
	"adventofcode/utils"
	"fmt"
	"strings"
)

func parseProcedures(input string) []models.Procedure {
	return operators.Map(strings.Split(input, "\n"), func(procedureInput string) (procedure models.Procedure) {
		procedure.Decode(procedureInput)
		return procedure
	})
}

func parseTitleStacks(input string) []int {
	return operators.Map(operators.Filter(strings.Split(input, " "),
		func(str string) bool { return len(str) > 0 },
	), utils.ParseInt)
}

func parseStacks(input string) (stacks []models.Stack) {
	list := strings.Split(input, "\n")
	titles := parseTitleStacks(list[len(list)-1])
	stacks = make([]models.Stack, len(titles))
	for i := len(list) - 2; i >= 0; i-- {
		chunks := operators.Map(operators.Chunk([]rune(list[i]), 4), func(chars []rune) string {
			return strings.Trim(string(chars), " []")
		})
		for index, value := range chunks {
			if len(value) > 0 {
				stacks[index].Push(value)
			}
		}
	}
	return stacks
}

func SupplyStacks(input string, operator func(models.Procedure, *models.Stack, *models.Stack)) string {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(parts[0])
	procedures := parseProcedures(parts[1])
	operators.ForEach(procedures, func(procedure models.Procedure) {
		operator(procedure, &stacks[procedure.Source-1], &stacks[procedure.Destination-1])
	})
	return strings.Join(operators.Map(stacks, func(stack models.Stack) (lastElement string) {
		lastElement, _ = stack.Pop()
		return lastElement
	}), "")
}

func step1(input string) string {
	return SupplyStacks(input, func(procedure models.Procedure, stack1 *models.Stack, stack2 *models.Stack) {
		var element string
		for i := 0; i < procedure.Number; i++ {
			element, _ = stack1.Pop()
			stack2.Push(element)
		}
	})
}

func step2(input string) string {
	return SupplyStacks(input, func(procedure models.Procedure, stack1 *models.Stack, stack2 *models.Stack) {
		tmpStack := models.Stack{}
		var element string
		for i := 0; i < procedure.Number; i++ {
			element, _ = stack1.Pop()
			tmpStack.Push(element)
		}
		for i := 0; i < procedure.Number; i++ {
			element, _ = tmpStack.Pop()
			stack2.Push(element)
		}
	})
}

func main() {
	const title, day = "--- Day 5: Supply Stacks ---", "2022/day05/"
	fmt.Println(title)

	example := utils.ParseFileToString(day + "example.txt")
	utils.AssertEqual(step1(example), "CMZ", "example step1")
	utils.AssertEqual(step2(example), "MCD", "example step2")

	input := utils.ParseFileToString(day + "input.txt")
	utils.AssertEqual(step1(input), "VWLCWGSDQ", "step1")
	utils.AssertEqual(step2(input), "TCGLQSLPW", "step2")
}
