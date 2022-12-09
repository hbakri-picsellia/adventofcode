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

func (point *Point) Follow(point2 Point) {
	if !point.IsClose(point2) {
		if point.X == point2.X {
			directionY := (point2.Y - point.Y) / int(math.Abs(float64(point2.Y-point.Y)))
			point.Y += directionY
		} else if point.Y == point2.Y {
			directionX := (point2.X - point.X) / int(math.Abs(float64(point2.X-point.X)))
			point.X += directionX
		} else {
			directionX := (point2.X - point.X) / int(math.Abs(float64(point2.X-point.X)))
			directionY := (point2.Y - point.Y) / int(math.Abs(float64(point2.Y-point.Y)))
			point.Y += directionY
			point.X += directionX
		}
	}
}

func (point *Point) Move(command Command) {
	switch command.Type {
	case Up:
		point.X += 1
	case Down:
		point.X -= 1
	case Right:
		point.Y += 1
	case Left:
		point.Y -= 1
	}
}

func (point *Point) ToString() string {
	return strconv.FormatInt(int64(point.X), 10) + "," + strconv.FormatInt(int64(point.Y), 10)
}
