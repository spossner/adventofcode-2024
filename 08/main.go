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

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev, config.WithSplitLines(), config.WithSplitWords(""))
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	points := map[string][]point.Point{}
	for y, row := range p.Cells {
		for x, c := range row {
			if c == "." {
				continue
			}
			coords, ok := points[c]
			if !ok {
				coords = []point.Point{}
			}
			points[c] = append(coords, point.Point{X: x, Y: y})
		}
	}
	bounds := grid.GetBounds(p.Cells)
	antinodes := set.NewSet[point.Point]()
	for _, adj := range points {
		if len(adj) < 2 {
			continue
		}
		if isPart2 {
			antinodes.Add(adj...)
		}
		for i := range len(adj) - 1 {
			for j := i + 1; j < len(adj); j++ {
				p1 := adj[i]
				p2 := adj[j]
				d := point.Point{X: p2.X - p1.X, Y: p2.Y - p1.Y}

				for {
					a1 := p1.Add(d.Mul(-1.0))
					a2 := p2.Add(d)
					//fmt.Println(k, a1, a2)
					found := false
					if bounds.Contains(a1) {
						found = true
						antinodes.Add(a1)
					}
					if bounds.Contains(a2) {
						found = true
						antinodes.Add(a2)
					}
					if !found || !isPart2 {
						break
					}
					p1 = a1
					p2 = a2
				}
			}
		}
	}

	result += len(antinodes)

	return result
}
