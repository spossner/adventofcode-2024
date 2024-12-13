package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
	"strconv"
)

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithSplitWords(""),
		config.WithDevFile("dev5.txt"),
	)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

type Step struct {
	pos    point.Point
	height int
	seen   set.Set[point.Point]
}

func NewStep(p point.Point, height int, seen set.Set[point.Point]) *Step {
	return &Step{p, height, seen}
}

func (s *Step) String() string {
	return fmt.Sprintf("(%d,%d) - %d", s.pos.X, s.pos.Y, s.height)
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	starts := make([]point.Point, 0)
	ends := make([]*Step, 0)

	for y, row := range p.Cells {
		for x, cell := range row {
			if cell == "." {
				continue
			}
			h, err := strconv.Atoi(cell)
			utils.AssertNil(err)
			switch h {
			case 0:
				starts = append(starts, point.Point{X: x, Y: y})
			case 9:
				ends = append(ends, NewStep(point.Point{X: x, Y: y}, 9, set.NewSet[point.Point]()))
			}
		}
	}

	bounds := rectangle.NewBounds(p.Cells)

	dp := make([][]int, len(p.Cells))
	for i := range len(p.Cells) {
		dp[i] = make([]int, len(p.Cells[0]))
	}

	q := queue.NewQueue[*Step]()
	for _, pos := range ends {
		q.Append(pos)
	}
	for !q.Empty() {
		s, ok := q.PopLeft()
		if !ok || s.height == 0 {
			continue
		}
		for _, newPos := range s.pos.DirectAdjacents() {
			if !bounds.Contains(newPos) {
				continue
			}
			if !isPart2 && s.seen.Contains(newPos) { // already visited in this trail
				continue
			}
			c := p.Cells[newPos.Y][newPos.X]
			if c == "." {
				continue
			}
			h, err := strconv.Atoi(c)
			utils.AssertNil(err)
			if h == s.height-1 {
				dp[newPos.Y][newPos.X]++
				s.seen.Add(newPos)
				q.Append(NewStep(newPos, s.height-1, s.seen))
			}
		}
	}

	for _, p := range starts {
		//fmt.Println(p, dp[p.Y][p.X])
		result += dp[p.Y][p.X]
	}

	return result
}
