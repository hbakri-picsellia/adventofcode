package main

import (
	"adventofcode/operators"
	"adventofcode/structs"
	"adventofcode/utils"
	"strings"
)

type Blueprint struct {
	Items              structs.List[int]
	NbItemsInspected   int
	rawOperation       []string
	TestDivider        int
	SuccessDestination int
	FailureDestination int
}

func (blueprint *Blueprint) Decode(input string) {
	rawItem := utils.GetStringBetween(input, "Starting items:", "\n")
	blueprint.Items = operators.Map(strings.Split(rawItem, ", "), utils.ParseStringToInt)
	blueprint.rawOperation = strings.Split(utils.GetStringBetween(input, "Operation: new =", "\n"), " ")
	blueprint.TestDivider = utils.ParseStringToInt(utils.GetStringBetween(input, "Test: divisible by", "\n"))
	blueprint.SuccessDestination = utils.ParseStringToInt(utils.GetStringBetween(input, "If true: throw to blueprint", "\n"))
	blueprint.FailureDestination = utils.ParseStringToInt(utils.GetStringBetween(input, "If false: throw to blueprint", "\n"))
}
