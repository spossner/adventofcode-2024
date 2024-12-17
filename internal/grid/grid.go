package grid

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/set"
	"iter"
)

type GridConfig struct {
	marker string
}
type GridConfigFunc func(cfg *GridConfig) *GridConfig

func NewGridConfig() *GridConfig {
	return &GridConfig{marker: "0"}
}

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

func FindMarker(grid [][]string, marker string) (point.Point, bool) {
	for pos, value := range IterateGrid(grid) {
		if value == marker {
			return pos, true
		}
	}
	return point.Point{}, false
}

func DumpGrid[T any](grid [][]T, options ...GridConfigFunc) {
	DumpGridWithMarker(grid, set.NewSet[point.Point](), options...)
}

func DumpGridWithMarker[T any](grid [][]T, markers set.Set[point.Point], options ...GridConfigFunc) {
	cfg := NewGridConfig()
	for _, fn := range options {
		cfg = fn(cfg)
	}
	for y, row := range grid {
		for x, cell := range row {
			if markers.Contains(point.Point{x, y}) {
				fmt.Print(cfg.marker)
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}
