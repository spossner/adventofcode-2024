package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/utils"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithGetInts(),
		config.WithDevFile(DEV_FILE),
	)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	if isPart2 {
		fmt.Println("PART2")
	}

	bounds := rectangle.NewRectangle(0, 0, 71, 71)
	n := 1024
	if cfg.Dev {
		bounds = rectangle.NewRectangle(0, 0, 7, 7)
		n = 12
	}

	if isPart2 {
		lower := 0
		upper := len(p.ParsedCells) - 1
		// binary find number of falling bytes where no path exists anymore
		for lower < upper {
			m := (lower + upper) >> 1
			g := grid.NewGrid(bounds.Width, bounds.Height)
			for i := range m {
				x, y := utils.Pick2From(p.ParsedCells[i])
				g.Set(x, y, "#")
			}
			costs, _ := g.Dijkstra(bounds.TopLeft(), bounds.BottomRight())
			if costs == 0 {
				upper = m - 1
			} else {
				lower = m + 1
			}
		}
		return p.ParsedCells[lower-1] // -1 because e.g. last falling byte of 1024 fallen bytes in total is in parsedCells[1023]
	}

	g := grid.NewGrid(bounds.Width, bounds.Height)
	for i := range n {
		x, y := utils.Pick2From(p.ParsedCells[i])
		g.Set(x, y, "#")
	}
	costs, _ := g.Dijkstra(bounds.TopLeft(), bounds.BottomRight())
	return costs
}
