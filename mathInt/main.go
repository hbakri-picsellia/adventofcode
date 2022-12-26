package mathInt

import "math"

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func Max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func Min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
