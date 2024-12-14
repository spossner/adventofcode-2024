package rectangle

import (
	"github.com/spossner/aoc2024/internal/point"
	"iter"
	"log"
)

type Rectangle struct {
	X, Y, Width, Height int
}

func NewRectangle(x, y, width, height int) Rectangle {
	return Rectangle{x, y, width, height}
}

func NewBounds[T any](matrix [][]T) Rectangle {
	if len(matrix) == 0 {
		return Rectangle{}
	}
	return Rectangle{0, 0, len(matrix[0]), len(matrix)}
}

func (r Rectangle) Extends(points ...point.Point) Rectangle {
	bounds := Rectangle{r.X, r.Y, r.Width, r.Height}
	for _, p := range points {
		if p.X < bounds.X {
			bounds.Width = bounds.Width + bounds.X - p.X
			bounds.X = p.X
		} else if p.X >= bounds.X+bounds.Width {
			bounds.Width = p.X - bounds.X + 1
		}

		if p.Y < bounds.Y {
			bounds.Height = bounds.Height + bounds.Y - p.Y
			bounds.Y = p.Y
		} else if p.Y >= bounds.Y+bounds.Height {
			bounds.Height = p.Y - bounds.Y + 1
		}
	}
	return bounds
}

func (r Rectangle) Contains(p point.Point) bool {
	if p.X >= r.X && p.Y >= r.Y && p.X < r.X+r.Width && p.Y < r.Y+r.Height {
		return true
	}
	return false
}

func (r Rectangle) Translate(offset point.Point) Rectangle {
	return Rectangle{r.X + offset.X, r.Y + offset.Y, r.Width, r.Height}
}

func (r Rectangle) MoveTo(p point.Point) Rectangle {
	return Rectangle{X: p.X, Y: p.Y, Width: r.Width, Height: r.Height}
}

func (r Rectangle) Grow(i int) Rectangle {
	newWidth := r.Width + (i << 1)
	newHeight := r.Height + (i << 1)
	if newWidth < 0 || newHeight < 0 {
		log.Fatalf("shrinked rectangle %v below zero width or height: %vx%v\n", r, newWidth, newHeight)
	}
	return Rectangle{X: r.X - i, Y: r.Y - i, Width: newWidth, Height: newHeight}
}

func (r Rectangle) All() iter.Seq2[int, point.Point] {
	return func(yield func(int, point.Point) bool) {
		i := 0
		for y := r.Y; y < r.Y+r.Height; y++ {
			for x := r.X; x < r.X+r.Width; x++ {
				if (!yield(i, point.Point{X: x, Y: y})) {
					return
				}
				i++
			}
		}
	}
}

func (r Rectangle) Center() point.Point {
	return point.Point{r.X + (r.Width >> 1), r.Y + (r.Height >> 1)}
}
