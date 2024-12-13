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
	"slices"
)

const DEV_FILE = "dev236.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithSplitWords(""),
		config.WithDevFile(DEV_FILE),
		config.WithDebug())
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

type Side struct {
	pos       point.Point
	direction point.Point
}

func (s Side) Translate(dx, dy int) Side {
	return Side{s.pos.Translate(dx, dy), s.direction}
}

type Plot struct {
	area      int
	perimeter int
	sides     set.Set[Side]
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	areas := []Plot{}
	seen := set.NewSet[point.Point]()
	for pos, v := range utils.IterateMatrix(p.Cells) {
		if seen.Contains(pos) {
			continue
		}
		if cfg.Debug {
			fmt.Printf("flooding %v %s ... ", pos, v)
		}
		area := findPlot(p, pos, v, seen)
		if cfg.Debug {
			fmt.Printf("%vqm, %vm, %v sides\n", area.area, area.perimeter, len(area.sides))

			sides := slices.Collect(area.sides.All())
			slices.SortStableFunc(sides, func(a, b Side) int {
				if cmpY := a.pos.Y - b.pos.Y; cmpY != 0 {
					return cmpY
				}
				return a.pos.X - b.pos.X
			})

			sideStrings, _ := utils.Map[Side, string](sides, func(s Side) (string, error) {
				return fmt.Sprintf("%v%s", s.pos, s.direction.Icon()), nil
			})
			fmt.Println(sideStrings)
		}
		areas = append(areas, area)
	}

	result = utils.Reduce(areas, func(acc int, item Plot) int {
		if isPart2 {
			return acc + (item.area * len(item.sides))
		}
		return acc + (item.area * item.perimeter)
	}, 0)

	return result
}

func findPlot(p *puzzle.Puzzle, start point.Point, v string, seen set.Set[point.Point]) Plot {
	q := queue.NewQueue[point.Point](start)
	area := 0
	sides := set.NewSet[Side]()
	borders := map[Side]struct{}{}
	bounds := rectangle.NewBounds(p.Cells)
	for !q.Empty() {
		pos, ok := q.PopLeft()
		if !ok {
			continue
		}
		if seen.Contains(pos) {
			continue
		}
		seen.Add(pos)
		area++

		for direction, adj := range pos.DirectAdjacents() {
			if !bounds.Contains(adj) || p.Cells[adj.Y][adj.X] != v {
				newSide := Side{pos: pos, direction: direction}
				borders[newSide] = struct{}{}
			} else {
				q.Append(adj)
			}
		}
	}

	for side := range borders {
		delta := point.WEST
		if side.direction == point.WEST || side.direction == point.EAST {
			delta = point.NORTH
		}
		for {
			checkSide := side.Translate(delta.X, delta.Y)
			if !utils.Contains(borders, checkSide) {
				break
			}
			side = checkSide
		}
		sides.Add(side)
	}

	return Plot{area: area, perimeter: len(borders), sides: sides}
}
