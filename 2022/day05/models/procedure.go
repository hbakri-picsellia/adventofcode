package models

import (
	"adventofcode/utils"
	"strings"
)

type Procedure struct {
	Number      int
	Source      int
	Destination int
}

func (procedure *Procedure) Decode(input string) {
	procedure.Number = utils.ParseInt(strings.Split(strings.SplitAfter(input, "move ")[1], " ")[0])
	procedure.Source = utils.ParseInt(strings.Split(strings.SplitAfter(input, "from ")[1], " ")[0])
	procedure.Destination = utils.ParseInt(strings.Split(strings.SplitAfter(input, "to ")[1], " ")[0])
}
