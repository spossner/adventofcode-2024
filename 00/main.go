package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
)

const DAY = -1

func part1(dev bool) any {
	cfg := config.NewConfig(DAY, dev) // split lines only
	return solve(cfg, false)
}

func part2(dev bool) any {
	cfg := config.NewConfig(DAY, dev, config.WithDevFile("dev.txt")) // split lines only, setting another dev file
	return solve(cfg, true)
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	fmt.Println(p.Data)

	return result
}
