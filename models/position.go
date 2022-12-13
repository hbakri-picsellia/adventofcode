package models

import "math"

type Position struct {
	I, J int
}

func (position *Position) Equals(other Position) bool {
	return position.I == other.I && position.J == other.J
}

func (position *Position) Add(other Position) Position {
	return Position{I: position.I + other.I, J: position.J + other.J}
}

func (position *Position) Distance(other Position) float64 {
	return math.Abs(float64(position.I-other.I)) + math.Abs(float64(position.J-other.J))
}
