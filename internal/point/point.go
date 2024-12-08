package point

type Point struct {
	X, Y int
}

var (
	UP    = Point{0, -1}
	DOWN  = Point{0, 1}
	LEFT  = Point{-1, 0}
	RIGHT = Point{1, 0}

	NORTH = UP
	SOUTH = DOWN
	WEST  = LEFT
	EAST  = RIGHT

	NORTH_EAST = Point{1, -1}
	SOUTH_EAST = Point{1, 1}
	NORTH_WEST = Point{-1, -1}
	SOUTH_WEST = Point{-1, 1}

	DIRECT_ADJACENT_POINTS = []Point{NORTH, EAST, SOUTH, WEST}

	ADJACENT_POINTS = []Point{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}
)

func (p Point) Translate(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
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
