package models

import (
	"adventofcode/operators"
	"adventofcode/structs"
	"adventofcode/utils"
	"strings"
)

type Monkey struct {
	Items              structs.List[int]
	NbItemsInspected   int
	rawOperation       []string
	TestDivider        int
	SuccessDestination int
	FailureDestination int
}

func getStringBetween(s string, sep1 string, sep2 string) string {
	return strings.TrimSpace(strings.SplitN(strings.SplitAfterN(s, sep1, 2)[1], sep2, 2)[0])
}

func (monkey *Monkey) Decode(input string) {
	rawItems := getStringBetween(input, "Starting items:", "\n")
	monkey.Items = operators.Map(strings.Split(rawItems, ", "), utils.ParseInt)
	monkey.rawOperation = strings.Split(getStringBetween(input, "Operation: new =", "\n"), " ")
	monkey.TestDivider = utils.ParseInt(getStringBetween(input, "Test: divisible by", "\n"))
	monkey.SuccessDestination = utils.ParseInt(getStringBetween(input, "If true: throw to monkey", "\n"))
	monkey.FailureDestination = utils.ParseInt(getStringBetween(input, "If false: throw to monkey", "\n"))
}

func (monkey *Monkey) Operation(value int) (result int) {
	evaluate := func(s string) int {
		if s == "old" {
			return value
		} else {
			return utils.ParseInt(s)
		}
	}
	switch monkey.rawOperation[1] {
	case "+":
		return evaluate(monkey.rawOperation[0]) + evaluate(monkey.rawOperation[2])
	case "*":
		return evaluate(monkey.rawOperation[0]) * evaluate(monkey.rawOperation[2])
	default:
		panic("Operator not implemented")
	}
}

func (monkey *Monkey) Test(value int) bool {
	return value%monkey.TestDivider == 0
}
