package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
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

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	fmt.Println(p.Data)

	return result
}
