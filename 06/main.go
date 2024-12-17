package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
)

var DIRECTIONS = map[string]point.Point{
	"^": point.UP,
	"v": point.DOWN,
	"<": point.LEFT,
	">": point.RIGHT,
}

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev, config.WithSplitLines(), config.WithSplitWords(""))
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func findGuard(cells [][]string) (pos point.Point, direction point.Point, ok bool) {
	for y, row := range cells {
		for x, cell := range row {
			if direction, ok := DIRECTIONS[cell]; ok {
				return point.Point{X: x, Y: y}, direction, true
			}
		}
	}
	return point.Point{}, point.Point{}, false
}

type Guard struct {
	pos       point.Point
	direction point.Point
}

func findPath(pos, direction point.Point, cells [][]string) (set.Set[Guard], bool) {
	seen := set.NewSet(Guard{pos, direction})
	bounds := grid.GetBounds(cells)
	for {
		newPos := pos.Translate(direction.X, direction.Y)
		if !bounds.Contains(newPos) {
			break
		}
		if cells[newPos.Y][newPos.X] == "#" {
			direction = direction.RotateRight()
			continue
		}
		newGuard := Guard{newPos, direction}
		if seen.Contains(newGuard) {
			return set.Set[Guard]{}, false
		}
		seen.Add(newGuard)
		pos = newPos
	}
	return seen, true
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")

	}

	pos, direction, ok := findGuard(p.Cells)
	utils.AssertTrue(ok)

	if isPart2 {
		for y, row := range p.Cells {
			for x, cell := range row {
				if x == pos.X && y == pos.Y {
					continue
				}
				if cell == "#" {
					continue
				}

				p.Cells[y][x] = "#"
				_, ok := findPath(pos, direction, p.Cells)
				if !ok {
					result++
				}
				p.Cells[y][x] = "."
			}
		}
	} else {
		seen, _ := findPath(pos, direction, p.Cells)
		fmt.Println(pos, seen)
		result = len(seen)
	}

	return result
}
