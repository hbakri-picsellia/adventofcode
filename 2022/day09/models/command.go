package models

import (
	"adventofcode/utils"
	"strings"
)

type CommandType string

const (
	Right CommandType = "R"
	Left              = "L"
	Up                = "U"
	Down              = "D"
)

type Command struct {
	Type   CommandType
	Number int
}

func (command *Command) Decode(input string) {
	parts := strings.SplitN(input, " ", 2)
	command.Type = CommandType(parts[0])
	command.Number = utils.ParseInt(parts[1])
}
