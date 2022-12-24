package models

import (
	"adventofcode/utils"
	"strings"
)

type File struct {
	Size int
	Name string
}

func (file *File) Decode(input string) {
	parts := strings.SplitN(input, " ", 2)
	file.Size = utils.ParseInt(parts[0])
	file.Name = parts[1]
}
