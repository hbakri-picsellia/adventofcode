package models

import (
	"math"
	"strings"
)

type Segment struct {
	Start Point
	End   Point
}

func (segment *Segment) Decode(segmentInput string) {
	edges := strings.Split(segmentInput, " -> ")
	segment.Start.Decode(edges[0])
	segment.End.Decode(edges[1])
}

func (segment *Segment) GetCoveredPointsWithoutDiagonals() (points []Point) {
	if segment.Start.X == segment.End.X {
		min := int(math.Min(float64(segment.Start.Y), float64(segment.End.Y)))
		max := int(math.Max(float64(segment.Start.Y), float64(segment.End.Y)))
		for i := min; i <= max; i++ {
			points = append(points, Point{X: segment.Start.X, Y: i})
		}
	} else if segment.Start.Y == segment.End.Y {
		min := int(math.Min(float64(segment.Start.X), float64(segment.End.X)))
		max := int(math.Max(float64(segment.Start.X), float64(segment.End.X)))
		for i := min; i <= max; i++ {
			points = append(points, Point{X: i, Y: segment.Start.Y})
		}
	}
	return points
}

func (segment *Segment) GetCoveredPointsWithDiagonals() (points []Point) {
	points = segment.GetCoveredPointsWithoutDiagonals()
	diffX := float64(segment.End.X - segment.Start.X)
	diffY := float64(segment.End.Y - segment.Start.Y)
	if math.Abs(diffX) == math.Abs(diffY) {
		coeffX := int(math.Abs(diffX) / diffX)
		coeffY := int(math.Abs(diffY) / diffY)
		for i := 0; i <= int(math.Abs(diffX)); i++ {
			points = append(points, Point{X: segment.Start.X + coeffX*i, Y: segment.Start.Y + coeffY*i})
		}
	}
	return points
}
