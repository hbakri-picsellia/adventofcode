package models

import (
	"strings"
)

type CommandType string

const (
	CD CommandType = "cd"
	LS             = "ls"
)

type Command struct {
	Type   CommandType
	Arg    string
	Result string
}

func (command *Command) Decode(input string) {
	parts := strings.SplitN(input, "\n", 2)
	commands := strings.SplitN(parts[0], " ", 2)
	command.Type = CommandType(commands[0])
	if len(commands) > 1 {
		command.Arg = commands[1]
	}
	if len(parts) > 1 {
		command.Result = parts[1]
	}
}
