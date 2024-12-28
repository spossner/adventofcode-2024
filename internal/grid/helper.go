package grid

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/rectangle"
	"iter"
	"slices"
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

func buildPath(item point.Point, previous map[point.Point]point.Point) Path {
	var path Path
	var ok bool
	for {
		path = append(path, item)
		item, ok = previous[item]
		if !ok {
			panic(fmt.Sprintf("missing previous node of %v", item))
		}
		if item == root {
			slices.Reverse(path)
			return path
		}
	}
}

func buildMultiPath(end point.Point, previous map[point.Point][]point.Point, directions bool) []Path {
	type pathItem struct {
		pos  point.Point
		path []point.Point
	}
	var q *queue.Queue[pathItem]
	if directions {
		q = queue.NewQueue(pathItem{pos: end, path: []point.Point{}})
	} else {
		q = queue.NewQueue(pathItem{pos: end, path: []point.Point{end}})
	}

	solutions := make([]Path, 0)
	for !q.Empty() {
		item, _ := q.PopLeft()
		v, ok := previous[item.pos]
		if !ok {
			panic(fmt.Sprintf("missing previous node of %v", item))
		}
		if v == nil { // root found
			slices.Reverse(item.path)
			solutions = append(solutions, item.path)
			continue
		}
		for _, p := range v {
			if directions {
				q.Append(pathItem{pos: p, path: append(slices.Clone(item.path), item.pos.Subtract(p))})
			} else {
				q.Append(pathItem{pos: p, path: append(slices.Clone(item.path), p)})
			}
		}
	}
	return solutions
}
