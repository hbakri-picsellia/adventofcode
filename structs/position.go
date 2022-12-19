package structs

import "math"

type Position struct {
	X, Y int
}

func (position *Position) Equals(other Position) bool {
	return position.X == other.X && position.Y == other.Y
}

func (position *Position) Add(other Position) Position {
	return Position{X: position.X + other.X, Y: position.Y + other.Y}
}

func (position *Position) Distance(other Position) float64 {
	return math.Abs(float64(position.X-other.X)) + math.Abs(float64(position.Y-other.Y))
}

func (position *Position) EuclideanDistance(other Position) float64 {
	return math.Sqrt(math.Pow(float64(position.X-other.X), 2) + math.Pow(float64(position.Y-other.Y), 2))
}

type PositionList struct {
	List[Position]
}
