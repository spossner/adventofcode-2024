package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"strconv"
)

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev, config.NoLineSplit, config.GetInts)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func splitStone(v int) []int {
	if v == 0 {
		return []int{1}
	}

	s := strconv.Itoa(v)
	n := len(s)
	if n%2 == 0 {
		v1, _ := strconv.Atoi(s[:n>>1])
		v2, _ := strconv.Atoi(s[n>>1:])
		return []int{v1, v2}
	}
	return []int{v * 2024}
}

func countStones(dp map[int]map[int]int, v int, n int) int {
	//fmt.Println("Count", v, "after", n, "steps...")
	counts, ok := dp[v] // is v already calculated?
	if !ok {
		counts = make(map[int]int)
		counts[0] = 1
		dp[v] = counts
	}
	sum, ok := counts[n]
	if ok {
		return sum
	}

	for _, subValue := range splitStone(v) {
		sum += countStones(dp, subValue, n-1)
	}
	counts[n] = sum
	//fmt.Printf("%d[%d] := %d\n", v, n, sum)
	return sum
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	dp := make(map[int]map[int]int)
	n := 25
	if isPart2 {
		n = 75
	}
	for _, v := range p.ParsedRows {
		result += countStones(dp, v, n)
	}

	return result
}
