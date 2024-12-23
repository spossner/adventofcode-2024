package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/counter"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithSplitWords(""),
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
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	g := grid.AsGrid(p.Cells)
	start, _ := g.FindMarker("S")
	end, _ := g.FindMarker("E")
	cheatsAllowed := utils.If(isPart2, 20, 2)

	path := g.Bfs(start, end)
	distancesToGoal := make(map[point.Point]int)
	for i, pos := range path {
		distancesToGoal[pos] = len(path) - i - 1
	}

	found := countCheats(g, path, distancesToGoal, cheatsAllowed)
	if cfg.Debug {
		fmt.Println(found)
	}

	for k, v := range found {
		if k >= utils.If(cfg.Dev, 50, 100) {
			if cfg.Debug {
				fmt.Println(v, "cheats save", k)
			}
			result += v
		}
	}

	return result
}

func countCheats(g grid.Grid, path []point.Point, distancesToGoal map[point.Point]int, allowed int) counter.Counter[int] {
	center := point.Point{0, 0}
	cheatAdj := set.NewSet[point.Point]()
	q := queue.NewQueue[point.Point](center)
	for range allowed {
		for range q.Len() {
			pos, _ := q.PopLeft()
			for _, adj := range pos.DirectAdjacents() {
				if adj == center || cheatAdj.Contains(adj) {
					continue
				}
				cheatAdj.Add(adj)
				q.Append(adj)
			}
		}
	}

	uncheatedCosts := len(path) - 1
	found := counter.NewCounter[int]()
	total := 0
	for costs, item := range path {
		for delta := range cheatAdj {
			cheated := item.Translate(delta.X, delta.Y)
			if !g.Contains(cheated) {
				continue
			}
			if g.Get(cheated.X, cheated.Y) == g.Wall() {
				continue
			}
			if restCosts, ok := distancesToGoal[cheated]; ok {
				cheatCosts := costs + item.Manhatten(cheated) + restCosts
				if cheatCosts < uncheatedCosts {
					total++
					found.Add(uncheatedCosts - cheatCosts)
				}
			}
		}
	}
	return found
}
