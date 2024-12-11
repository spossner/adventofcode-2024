package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"regexp"
)

func part1(dev bool) any {
	return solve(dev, false)
}

func part2(dev bool) any {
	return solve(dev, true)
}

func solve(dev bool, isPart2 bool) any {
	if isPart2 {
	}
	cfg := config.NewConfig(utils.GetPackageDir(), dev, config.NoLineSplit)
	if isPart2 {
		cfg.DevFile = "dev-2.txt"
	}
	p := puzzle.NewPuzzle(cfg)
	fmt.Println(p.Data)
	result := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if isPart2 {
		re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	}
	matches := re.FindAllSubmatch([]byte(p.Data), -1)
	active := true
	for _, match := range matches {
		if isPart2 {
			if string(match[0]) == "do()" {
				active = true
				continue
			}
			if string(match[0]) == "don't()" {
				active = false
				continue
			}
		}

		if !active {
			continue
		}
		result += utils.MustAtoi(string(match[1])) * utils.MustAtoi(string(match[2]))
	}
	return result
}
