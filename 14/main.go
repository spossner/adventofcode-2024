package main

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/counter"
	"github.com/spossner/aoc2024/internal/pair"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/utils"
)

const DEV_FILE = "dev.txt"

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

func dumpRobots(robots counter.Counter[point.Point], bounds rectangle.Rectangle) {
	for y := range bounds.Height {
		for x := range bounds.Width {
			if _, ok := robots[point.Point{x, y}]; ok {
				//fmt.Printf("%d", v)
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func moveRobots(cells [][]int, moves int, bounds rectangle.Rectangle) counter.Counter[point.Point] {
	robots := counter.NewCounter[point.Point]()
	for _, robot := range cells {
		x, y, vx, vy := utils.Pick4From(robot)
		nx := utils.Mod(x+vx*moves, bounds.Width)
		ny := utils.Mod(y+vy*moves, bounds.Height)
		robots.Add(point.Point{nx, ny})
	}
	return robots
}

func calculateQuadrants(robots counter.Counter[point.Point], bounds rectangle.Rectangle) [4]int {
	var quadrants [4]int
	mid := point.Point{bounds.Width >> 1, bounds.Height >> 1}

	for pos, c := range robots {
		if pos.X < mid.X {
			if pos.Y < mid.Y {
				quadrants[0] += c
			} else if pos.Y > mid.Y {
				quadrants[2] += c
			}
		} else if pos.X > mid.X {
			if pos.Y < mid.Y {
				quadrants[1] += c
			} else if pos.Y > mid.Y {
				quadrants[3] += c
			}
		}
	}
	return quadrants
}

func calculateTogetherness(robots counter.Counter[point.Point]) int {
	togetherness := 0
	for p := range robots {
		for _, adj := range p.Adjacents() {
			if utils.Contains(robots, adj) {
				togetherness++
			}
		}
	}
	return togetherness
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	bounds := rectangle.NewRectangle(0, 0, 101, 103)
	if cfg.Dev {
		bounds = rectangle.NewRectangle(0, 0, 11, 7)
	}

	if isPart2 {
		var maxRobots counter.Counter[point.Point]
		maxTogetherness := pair.NewIntPair(0, 0)
		for i := 0; i < 10000; i++ {
			robots := moveRobots(p.ParsedCells, i, bounds)
			togetherness := calculateTogetherness(robots)
			if togetherness > maxTogetherness.A {
				maxTogetherness = pair.NewIntPair(togetherness, i)
				maxRobots = robots

			}
		}
		dumpRobots(maxRobots, bounds)
		return maxTogetherness.B
	}

	robots := moveRobots(p.ParsedCells, 100, bounds)
	quadrants := calculateQuadrants(robots, bounds)

	result = utils.Product(quadrants[:])

	return result
}
