package models

import (
	"math"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func (point *Point) IsClose(point2 Point) bool {
	return math.Abs(float64(point2.X)-float64(point.X)) <= 1 &&
		math.Abs(float64(point2.Y)-float64(point.Y)) <= 1
}

func (point *Point) ToString() string {
	return strconv.FormatInt(int64(point.X), 10) + "," + strconv.FormatInt(int64(point.Y), 10)
}
