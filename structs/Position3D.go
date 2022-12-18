package structs

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"strings"
)

type Position3D struct {
	X, Y, Z int
}

func MakePosition3D(s string) Position3D {
	coordinates := operators.Map(strings.Split(s, ","), utils.ParseStringToInt)
	return Position3D{X: coordinates[0], Y: coordinates[1], Z: coordinates[2]}
}

func (position *Position3D) Add(direction [3]int) Position3D {
	return Position3D{X: position.X + direction[0], Y: position.Y + direction[1], Z: position.Z + direction[2]}
}

func (position *Position3D) IsInferior(position2 Position3D) bool {
	return position.X <= position2.X && position.Y <= position2.Y && position.Z <= position2.Z
}
