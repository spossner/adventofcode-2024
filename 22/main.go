package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/counter"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"iter"
)

const DEV_FILE = "dev2.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithGetInts(),
		config.WithDevFile(DEV_FILE),
	)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

var MOD = 1 << 24

func mixAndPrune(secret, value int) int {
	return (secret ^ value) % MOD
}

func getPrices(secret, count int) []int {
	prices := make([]int, count+1)

	prices[0] = secret % 10

	for i := range count {
		v := secret << 6
		secret = mixAndPrune(secret, v)
		v = secret >> 5
		secret = mixAndPrune(secret, v)
		v = secret << 11
		secret = mixAndPrune(secret, v)
		prices[i+1] = secret % 10
	}
	return prices
}

type seller struct {
	secret int
}

func (s seller) randomNumbers(count int) iter.Seq2[int, int] {
	i := 0

	return func(yield func(int, int) bool) {
		for range count {
			v := s.secret << 6
			s.secret = mixAndPrune(s.secret, v)
			v = s.secret >> 5
			s.secret = mixAndPrune(s.secret, v)
			v = s.secret << 11
			s.secret = mixAndPrune(s.secret, v)
			yield(i, s.secret)
			i++
		}
	}
}

type priceChanges [4]int

func solve2(cfg *config.Config) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0

	cnt := make([]counter.Counter[priceChanges], len(p.ParsedRows))
	for no, secret := range p.ParsedRows {
		cnt[no] = counter.NewCounter[priceChanges]()

		prices := getPrices(secret, 2000)
		for i := 4; i < len(prices); i++ {
			change := [4]int{
				prices[i-3] - prices[i-4],
				prices[i-2] - prices[i-3],
				prices[i-1] - prices[i-2],
				prices[i] - prices[i-1],
			}
			if _, ok := cnt[no][change]; !ok {
				cnt[no][change] = prices[i] // store the first time the sequence appears
			}
		}
	}

	combined := counter.NewCounter[priceChanges]()
	var maxSequence priceChanges

	for no := range len(p.ParsedRows) {
		for c, v := range cnt[no] {
			combined[c] += v
			if combined[c] > result {
				result = combined[c]
				maxSequence = c
			}
		}
	}

	fmt.Println(maxSequence, result)
	return result
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	if isPart2 {
		fmt.Println("PART2")
		return solve2(cfg)
	}

	p := puzzle.NewPuzzle(cfg)
	result := 0

	for _, secret := range p.ParsedRows {
		s := seller{secret: secret}
		last := secret
		for _, r := range s.randomNumbers(2000) {
			last = r
		}
		result += last
	}

	return result
}
