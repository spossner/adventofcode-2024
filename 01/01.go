package _1

import (
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/counter"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"slices"
)

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithGetInts(),
	)
}

func part1(dev bool) any {
	p := puzzle.NewPuzzle(createConfig(dev))
	total := 0
	cols := utils.Transpose(p.ParsedCells)
	slices.Sort(cols[0])
	slices.Sort(cols[1])
	for i := range len(cols[0]) {
		total += utils.AbsInt(cols[0][i] - cols[1][i])
	}
	return total
}

func part2(dev bool) any {
	p := puzzle.NewPuzzle(createConfig(dev))
	cols := utils.Transpose(p.ParsedCells)

	cnt := counter.NewCounter(cols[1]...)
	total := 0
	for _, n := range cols[0] {
		total += n * cnt[n]
	}
	return total
}
