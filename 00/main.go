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
	cfg := config.NewConfig(DAY, dev) // split lines only
	if isPart2 {
		// e.g. update dev sample file
		// cfg.DevFile = "dev-2.txt"
	}
	p := puzzle.NewPuzzle(cfg)
	result := 0

	fmt.Println(p.Data)

	return result
}
