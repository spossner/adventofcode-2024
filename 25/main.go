package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/counter"
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

func parseBlock(block string) counter.Counter[int] {
	data, _ := utils.Map[string, []string](strings.Split(block, "\n"), func(s string) ([]string, error) {
		return strings.Split(s, ""), nil
	})
	cnt := counter.NewCounter[int]()
	for i := 1; i < len(data)-1; i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "#" {
				cnt.Add(j) // just count no of #.. do not care about from top or from bottom
			}
		}
	}
	return cnt
}

func checkKeyLock(lock, key counter.Counter[int]) bool {
	for i := range 5 {
		// check if there are more than 5 pins in total of lock and key
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	locks := make([]counter.Counter[int], 0)
	keys := make([]counter.Counter[int], 0)
	for _, block := range p.Rows {
		if block[0] == '#' { // locks have # in first row; keys not
			locks = append(locks, parseBlock(block))
		} else {
			keys = append(keys, parseBlock(block))
		}
	}

	for _, lock := range locks {
		for _, key := range keys {
			if checkKeyLock(lock, key) {
				result += 1
			}
		}
	}

	return result
}
