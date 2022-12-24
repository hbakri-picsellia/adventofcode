package models

import (
	"adventofcode/operators"
	"adventofcode/utils"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (point *Point) Decode(pointInput string) {
	coordinates := operators.Map(strings.Split(pointInput, ","), utils.ParseInt)
	point.X = coordinates[0]
	point.Y = coordinates[1]
}

func (point *Point) Encode() string {
	return strconv.FormatInt(int64(point.X), 10) + "," + strconv.FormatInt(int64(point.Y), 10)
}
