package models

import (
	"adventofcode/operators"
	"strings"
)

type Matrix[T any] [][]T

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
				return Position{I: i, J: j}
			}
		}
	}
	return Position{I: -1, J: -1}
}

func (matrix *Matrix[T]) Contains(position Position) bool {
	n, m := matrix.GetDimension()
	return position.I >= 0 && position.I < n && position.J >= 0 && position.J < m
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

func Map[T, U any](matrix Matrix[T], f func(T) U) Matrix[U] {
	return operators.Map(matrix, func(row []T) []U {
		return operators.Map(row, f)
	})
}
