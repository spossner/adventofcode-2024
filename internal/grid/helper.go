package grid

import (
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/rectangle"
	"iter"
)

func GetBounds[T any](grid [][]T) rectangle.Rectangle {
	if len(grid) == 0 {
		return rectangle.Rectangle{}
	}
	return rectangle.Rectangle{0, 0, len(grid[0]), len(grid)}
}

func IterateGrid[T any](grid [][]T) iter.Seq2[point.Point, T] {
	return func(yield func(point.Point, T) bool) {
	Outer:
		for y, row := range grid {
			for x, cell := range row {
				if (!yield(point.Point{X: x, Y: y}, cell)) {
					break Outer
				}
			}
		}
	}
}

func PickFrom[T any](grid [][]T, pos point.Point) T {
	return grid[pos.Y][pos.X]
}
