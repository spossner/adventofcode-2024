package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"math"
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

type steps struct {
	sequence string
	robot    int
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	cache := make(map[steps]int)
	for _, sequence := range p.Rows {
		v, ok := strings.CutSuffix(sequence, "A")
		if !ok {
			panic("illegal sequence (not ending at A): " + sequence)
		}
		no := utils.MustAtoi(v)

		robots := utils.If(isPart2, 25, 2)
		result += no * calculateKeyPresses(NUMERIC, sequence, robots, cache)
	}

	return result
}

func calculateKeyPresses(g grid.Grid, sequence string, robot int, cache map[steps]int) int {
	key := steps{sequence, robot}
	if v, ok := cache[key]; ok {
		return v
	}

	var s, e point.Point
	length := 0
	e, _ = g.FindMarker("A")
	for i := 0; i < len(sequence); i++ {
		s = e
		e, _ = g.FindMarker(string(sequence[i]))

		moves := g.BfsAll(s, e, grid.WithDirections())
		paths := utils.MustMap[grid.Path, string](moves, func(path grid.Path) (string, error) {
			return path.String() + "A", nil
		})

		if robot == 0 {
			length += len(utils.PickFirst(paths))
		} else {
			best := math.MaxInt
			for _, move := range paths {
				// check key presses of all paths in subsequent directional keypads... and choose the one with min length
				best = utils.Min(best, calculateKeyPresses(DIRECTIONAL, move, robot-1, cache))
			}
			length += best
		}
	}

	cache[key] = length
	return length
}
