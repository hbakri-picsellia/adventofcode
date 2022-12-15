package structs

import (
	"adventofcode/operators"
	"math"
	"strings"
)

type Path struct {
	points []Point
}

func (path *Path) Decode(segmentInput string) {
	path.points = operators.Map(strings.Split(segmentInput, " -> "), func(s string) (point Point) {
		point.Decode(s)
		return point
	})
}

func (path *Path) GetCoveredPoints() (points []Point) {
	for i := 1; i < len(path.points); i++ {
		if path.points[i-1].X == path.points[i].X {
			min := int(math.Min(float64(path.points[i-1].Y), float64(path.points[i].Y)))
			max := int(math.Max(float64(path.points[i-1].Y), float64(path.points[i].Y)))
			for y := min; y <= max; y++ {
				points = append(points, Point{X: path.points[i].X, Y: y})
			}
		} else if path.points[i-1].Y == path.points[i].Y {
			min := int(math.Min(float64(path.points[i-1].X), float64(path.points[i].X)))
			max := int(math.Max(float64(path.points[i-1].X), float64(path.points[i].X)))
			for x := min; x <= max; x++ {
				points = append(points, Point{X: x, Y: path.points[i].Y})
			}
		}
	}
	return points
}
