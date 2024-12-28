package point

import (
	"fmt"
	"iter"
)

type Point struct {
	X, Y int
}

var (
	NORTH = Point{0, -1}
	SOUTH = Point{0, 1}
	WEST  = Point{-1, 0}
	EAST  = Point{1, 0}

	NORTH_EAST = Point{1, -1}
	SOUTH_EAST = Point{1, 1}
	NORTH_WEST = Point{-1, -1}
	SOUTH_WEST = Point{-1, 1}

	DIRECT_ADJACENT_POINTS = []Point{
		NORTH,
		EAST,
		SOUTH,
		WEST,
	}

	ADJACENT_POINTS = []Point{
		NORTH,
		NORTH_EAST,
		EAST,
		SOUTH_EAST,
		SOUTH,
		SOUTH_WEST,
		WEST,
		NORTH_WEST,
	}

	OPPOSITE_DIRECTION = map[Point]Point{
		NORTH: SOUTH,
		SOUTH: NORTH,
		EAST:  WEST,
		WEST:  EAST,
	}

	DIRECTIONS = map[string]Point{
		"N": NORTH, "n": NORTH, "U": NORTH, "u": NORTH, "^": NORTH,
		"S": SOUTH, "s": SOUTH, "D": SOUTH, "d": SOUTH, "v": SOUTH,
		"E": EAST, "e": EAST, "R": EAST, "r": EAST, ">": EAST,
		"W": WEST, "w": WEST, "L": WEST, "l": WEST, "<": WEST,
	}

	NAMES = map[Point]string{
		NORTH:      "N",
		SOUTH:      "S",
		WEST:       "W",
		EAST:       "E",
		NORTH_WEST: "NW",
		NORTH_EAST: "NE",
		SOUTH_EAST: "SE",
		SOUTH_WEST: "SW",
	}

	ICONS = map[Point]string{
		NORTH: "^",
		SOUTH: "v",
		WEST:  "<",
		EAST:  ">",
	}

	UP    = NORTH
	DOWN  = SOUTH
	LEFT  = WEST
	RIGHT = EAST
)

func FromValues(values ...int) Point {
	if len(values) < 2 {
		panic("too few values to create a point")
	}
	return Point{values[0], values[1]}
}

func FromDirection(direction string) Point {
	return DIRECTIONS[direction]
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func (p Point) Name() string {
	if name, ok := NAMES[p]; ok {
		return name
	}
	return "UNKNWON"
}

func (p Point) Icon() string {
	if name, ok := ICONS[p]; ok {
		return name
	}
	return "?"
}

func (p Point) Translate(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func (p Point) Subtract(other Point) Point {
	return Point{p.X - other.X, p.Y - other.Y}
}

func (p Point) Mul(scale float64) Point {
	return Point{int(float64(p.X) * scale), int(float64(p.Y) * scale)}
}

func (p Point) RotateRight() Point {
	return Point{-p.Y, p.X}
}

func (p Point) RotateLeft() Point {
	return Point{p.Y, -p.X}
}

// DirectAdjacents gives a range func through the direct adjacents points.
// It yields the direction as first value and the position of the adjacent as Point as second value.
func (p Point) DirectAdjacents() iter.Seq2[Point, Point] {
	return func(yield func(Point, Point) bool) {
		for _, delta := range DIRECT_ADJACENT_POINTS {
			if !(yield(delta, Point{X: p.X + delta.X, Y: p.Y + delta.Y})) {
				break
			}
		}
	}
}

func (p Point) Adjacents() iter.Seq2[Point, Point] {
	return func(yield func(Point, Point) bool) {
		for _, delta := range ADJACENT_POINTS {
			if !(yield(delta, Point{X: p.X + delta.X, Y: p.Y + delta.Y})) {
				break
			}
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (p Point) Manhatten(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}
