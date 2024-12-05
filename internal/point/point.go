package point

type Point struct {
	X, Y int
}

var DIRECT_ADJACENT_POINTS = []Point{
	Point{0, -1},
	Point{-1, 0},
	Point{1, 0},
	Point{0, 1},
}

var ADJACENT_POINTS = []Point{
	Point{-1, -1},
	Point{0, -1},
	Point{1, -1},
	Point{-1, 0},
	Point{1, 0},
	Point{-1, 1},
	Point{0, 1},
	Point{1, 1},
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

func (v Point) Translate(dx, dy int) Point {
	return Point{v.X + dx, v.Y + dy}
}
