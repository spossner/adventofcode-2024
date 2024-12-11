package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/utils"
	"strconv"
)

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev, config.GetInts)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func testMath(goal int, numbers []int, isPart2 bool) bool {
	type Step struct {
		ptr   int
		total int
	}
	n := len(numbers)

	q := queue.NewQueue(Step{1, numbers[0]})
	for {
		step, ok := q.PopLeft()
		if !ok {
			break
		}
		if goal != 0 && step.total > goal {
			continue // already too high
		}

		if step.ptr == n {
			if step.total == goal {
				return true
			}
			continue
		}
		j := numbers[step.ptr]
		q.Append(Step{step.ptr + 1, step.total + j})
		q.Append(Step{step.ptr + 1, step.total * j})
		if isPart2 {
			combined, err := strconv.Atoi(fmt.Sprintf("%d%d", step.total, j))
			utils.AssertNil(err, "error combining %v and %v", step.total, j)
			q.Append(Step{step.ptr + 1, combined})
		}
	}
	return false
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	for _, row := range p.ParsedCells {
		if testMath(row[0], row[1:], isPart2) {
			result += row[0]
		}
	}

	return result
}
