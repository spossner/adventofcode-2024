package _1

import (
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"slices"
)

const DAY = 1

func part1(dev bool) any {
	cfg := config.NewConfig(DAY, dev, config.SplitWords, config.ParseInts)
	p := puzzle.NewPuzzle(cfg)
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
	cfg := config.NewConfig(DAY, dev)
	p := puzzle.NewPuzzle(cfg)
	cols := utils.Transpose(p.ParsedCells)

	counter := utils.Counter(cols[1])
	total := 0
	for _, n := range cols[0] {
		total += n * counter[n]
	}
	return total
}
