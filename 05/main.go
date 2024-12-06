package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"regexp"
	"slices"
	"strings"
)

const DAY = 5

func part1(dev bool) any {
	cfg := config.NewConfig(DAY, dev) // split lines only
	return solve(cfg, false)
}

func part2(dev bool) any {
	cfg := config.NewConfig(DAY, dev, config.WithDevFile("dev.txt")) // split lines only, setting another dev file
	return solve(cfg, true)
}

func parseData(p *puzzle.Puzzle) (map[int][]int, [][]int) {
	blocks := regexp.MustCompile(`\n\s*\n`).Split(p.Data, -1)

	mustComeAfter := make(map[int][]int)
	q, err := utils.Map(strings.Split(blocks[0], "\n"), utils.GetInts)
	utils.AssertZero(err)
	for _, pair := range q {
		if l, ok := mustComeAfter[pair[0]]; ok {
			mustComeAfter[pair[0]] = append(l, pair[1])
		} else {
			l := make([]int, 1)
			l[0] = pair[1]
			mustComeAfter[pair[0]] = l
		}
	}

	updates, err := utils.Map(strings.Split(blocks[1], "\n"), utils.GetInts)
	utils.AssertZero(err)
	return mustComeAfter, updates
}

func insertOrdered(list []int, el int, rules map[int][]int) []int {
	if len(list) == 0 {
		return append(list, el)
	}

	for i, v := range list {
		if slices.Contains(rules[el], v) {
			return slices.Insert(list, i, el)
		}
	}
	return append(list, el)
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}
	mustComeAfter, updates := parseData(p)

	wrong := make([][]int, 0)

Outer:
	for _, update := range updates {
		seen := make(map[int]struct{})
		for _, el := range update {
			for _, forbidden := range mustComeAfter[el] {
				if _, ok := seen[forbidden]; ok {
					//fmt.Println("FORBIDDEN", update)
					wrong = append(wrong, update)
					continue Outer
				}
			}
			seen[el] = struct{}{}
		}
		//fmt.Println("OK", update, update[len(update)>>1])
		result += update[len(update)>>1]
	}

	result2 := 0
	for _, update := range wrong {
		fmt.Println("PROCESSING", update)
		fixed := make([]int, 0, len(update))
		for _, v := range update {
			fixed = insertOrdered(fixed, v, mustComeAfter)
			//fmt.Printf("%d: inserted %v: %v\n", i, v, fixed)
		}
		result2 += fixed[len(fixed)>>1]
	}

	return fmt.Sprintf("%v,%v", result, result2)
}