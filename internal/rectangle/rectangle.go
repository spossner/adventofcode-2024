package rectangle

import "github.com/spossner/aoc2024/internal/point"

type Rectangle struct {
	X, Y, Width, Height int
}

func NewRectangle(x, y, width, height int) Rectangle {
	return Rectangle{x, y, width, height}
}

func (r Rectangle) Contains(p point.Point) bool {
	if p.X < r.X || p.Y < r.Y || p.X >= r.X+r.Width || p.Y >= r.Y+r.Height {
		return false
	}
	return true
}