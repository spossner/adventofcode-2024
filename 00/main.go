package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
)

const DAY = -1

func part1(dev bool) any {
	return solve(dev, false)
}

func part2(dev bool) any {
	return solve(dev, true)
}

func solve(dev bool, isPart2 bool) any {
	if isPart2 {
	}
	cfg := config.NewConfig(DAY, dev) // split lines only
	// cfg := config.NewConfig(DAY, dev, config.SplitWords, config.ParseInts)
	p := puzzle.NewPuzzle(cfg)
	fmt.Println(p.Data)
	result := 0

	return result
}
