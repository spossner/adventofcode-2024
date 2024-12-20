package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/pair"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"strings"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitWords("\n\n"),
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

	count, result := 0, 0
	p := puzzle.NewPuzzle(cfg)
	if isPart2 {
		fmt.Println("PART2")
	}

	towels := strings.Split(utils.PickFirst(p.Rows), ", ")
	patterns := strings.Split(utils.PickLast(p.Rows), "\n")

	towelMap := make(map[string][]string)
	for _, t := range towels {
		c := t[:1]
		if !utils.Contains(towelMap, c) {
			towelMap[c] = make([]string, 0)
		}
		towelMap[c] = append(towelMap[c], t)
	}

	for _, pattern := range patterns {
		i := countSolutions(pattern, towelMap)
		result += i
		if i > 0 {
			count++
		}
	}

	return pair.NewIntPair(count, result)
}

func countSolutions(pattern string, towelMap map[string][]string) int {
	n := len(pattern)
	dp := make([]int, n+1)

	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		p := pattern[i:]
		c := p[:1]
		if candidates, ok := towelMap[c]; ok {
			for _, candidate := range candidates {
				if len(candidate) > n-i {
					continue // too long
				}
				if !strings.HasPrefix(p, candidate) {
					continue // no match
				}
				dp[i] += dp[i+len(candidate)]
			}
		}
	}
	return dp[0]
}
