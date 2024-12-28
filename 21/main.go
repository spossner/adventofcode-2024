package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"math"
	"os"
	"strconv"
	"strings"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithDevFile(DEV_FILE),
	)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

var NUMERIC = grid.AsGrid([][]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"#", "0", "A"},
})

var DIRECTIONAL = grid.AsGrid([][]string{
	{"#", "^", "A"},
	{"<", "v", ">"},
})

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	for _, sequence := range p.Rows {
		candidates := shortestSequences(findPaths(NUMERIC, sequence, false))

		steps := utils.If(isPart2, 25, 2)
		for i := range steps {
			newCandidates := []string{}
			for j, c := range candidates {
				fmt.Println(sequence, i, j, len(c))
				fmt.Print(".")
				newCandidates = shortestSequences(append(newCandidates, findPaths(DIRECTIONAL, c, true)...))
			}
			fmt.Println()
			candidates = newCandidates
		}

		fmt.Println(sequence, minLength(candidates))

		m := minLength(candidates)
		if code, ok := strings.CutSuffix(sequence, "A"); ok {
			no, _ := strconv.Atoi(code)
			result += no * m
			continue
		}
		os.Exit(1)
	}

	return result
}

func shortestSequences(candidates []string) []string {
	m := minLength(candidates)
	return utils.MustFilter(candidates, func(path string) (bool, error) {
		return len(path) == m, nil
	})
}

func minLength(candidates []string) int {
	return utils.Reduce[int, int](
		utils.MustMap[string, int](
			candidates,
			func(path string) (int, error) {
				return len(path), nil
			},
		),
		func(acc int, item int) int {
			return utils.If(item < acc, item, acc)
		},
		math.MaxInt,
	)
}

var cache = make(map[string][]string)

func findPaths(g grid.Grid, sequence string, withCache bool) []string {
	var candidates []string
	var s, e point.Point
	e, _ = g.FindMarker("A")
	for i := 0; i < len(sequence); i++ {
		s = e
		e, _ = g.FindMarker(string(sequence[i]))
		iconStart := g.Get(s.X, s.Y)
		iconEnd := g.Get(e.X, e.Y)

		var paths []string
		if withCache {
			if v, ok := cache[iconStart+iconEnd]; ok {
				candidates = product(candidates, v)
				continue
			}
		}

		solutions := g.BfsAll(s, e, grid.WithDirections())
		paths = utils.MustMap[grid.Path, string](solutions, func(path grid.Path) (string, error) {
			return path.String() + "A", nil
		})
		if withCache {
			cache[iconStart+iconEnd] = paths
		}
		candidates = product(candidates, paths)
	}
	return candidates
}

func product(candidates []string, nextCandidates []string) []string {
	combinedCandidates := make([]string, 0, len(candidates)*len(nextCandidates))
	if candidates == nil {
		return nextCandidates
	}
	for _, c := range candidates {
		for _, d := range nextCandidates {
			combinedCandidates = append(combinedCandidates, c+d)
		}
	}
	return combinedCandidates
}
