package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
)

func part1(dev bool) any {
	return solve(dev, false)
}

func part2(dev bool) any {
	return solve(dev, true)
}

func solve(dev bool, isPart2 bool) any {
	cfg := config.NewConfig(utils.GetPackageDir(), dev, config.SplitFields, config.GetInts)
	p := puzzle.NewPuzzle(cfg)
	result := 0

	if isPart2 {
		fmt.Println("PART2")
	}

	for _, levels := range p.ParsedCells {
		for skip := range len(levels) {
			lvls := utils.Cut(levels, skip)
			level := lvls[0]
			inc := lvls[1] > level
			safe := true

			for _, nextLevel := range lvls[1:] {
				if inc && nextLevel <= level || !inc && nextLevel >= level {
					safe = false
					break
				}
				step := utils.AbsInt(nextLevel - level)
				if step < 1 || step > 3 {
					safe = false
					break
				}
				level = nextLevel
			}
			if safe {
				fmt.Println(levels, "Safe")
				result++
				break // also break the skip loop
			}
		}
	}

	return result
}
