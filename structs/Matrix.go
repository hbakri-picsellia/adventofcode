package structs

import (
	"adventofcode/operators"
	"fmt"
	"strings"
)

type Matrix[T any] [][]T

func MakeMatrix[T any](n, m int, defaultValue T) Matrix[T] {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i := 0; i < len(rows); i++ {
		rows[i] = defaultValue
	}
	for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
		endRow := startRow + m
		matrix[i] = rows[startRow:endRow:endRow]
	}
	return matrix
}

func (matrix *Matrix[T]) Decode(input string, rowSeparator string, columnSeparator string, f func(string) T) {
	*matrix = operators.Map(strings.Split(input, rowSeparator), func(rowInput string) []T {
		return operators.Map(operators.Filter(strings.Split(rowInput, columnSeparator),
			func(str string) bool { return len(str) > 0 },
		), f)
	})
}

func (matrix *Matrix[T]) GetRow(index int) []T {
	return (*matrix)[index]
}

func (matrix *Matrix[T]) GetColumn(index int) []T {
	return operators.Map(*matrix, func(row []T) T {
		return row[index]
	})
}

func (matrix *Matrix[T]) GetDimension() (int, int) {
	return len(*matrix), len((*matrix)[0])
}

func (matrix *Matrix[T]) Find(f func(T) bool) Position {
	for i := range *matrix {
		for j := range (*matrix)[i] {
			if f((*matrix)[i][j]) {
				return Position{X: i, Y: j}
			}
		}
	}
	return Position{X: -1, Y: -1}
}

func (matrix *Matrix[T]) Contains(position Position) bool {
	n, m := matrix.GetDimension()
	return position.X >= 0 && position.X < n && position.Y >= 0 && position.Y < m
}

func (matrix *Matrix[T]) GetNeighbors(position Position) (neighbors []Position) {
	newDirections := []Position{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, newDirection := range newDirections {
		newPosition := position.Add(newDirection)
		if matrix.Contains(newPosition) {
			neighbors = append(neighbors, newPosition)
		}
	}
	return neighbors
}

func (matrix *Matrix[T]) Display() {
	for _, row := range *matrix {
		for _, element := range row {
			fmt.Print(element)
		}
		fmt.Println()
	}
}

func Map[T, U any](matrix Matrix[T], f func(T) U) Matrix[U] {
	return operators.Map(matrix, func(row []T) []U {
		return operators.Map(row, f)
	})
}
