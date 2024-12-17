package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/point"
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

func findRobot(area [][]string) point.Point {
	for pos, value := range grid.IterateGrid(area) {
		if value == "@" {
			return pos
		}
	}
	panic("no robot found")
}

func move(area [][]string, start, direction point.Point) bool {
	adj := start.Add(direction)
	v := grid.PickFrom(area, adj)
	if v == "#" {
		return false // can not move in that direction
	}

	if v == "O" && !move(area, adj, direction) {
		return false // can not clear the adjacent position
	}

	area[adj.Y][adj.X] = area[start.Y][start.X]
	area[start.Y][start.X] = "."

	return true
}

func move2(area [][]string, start, direction point.Point, justCheck bool) bool {
	if grid.PickFrom(area, start) == "#" { // check starting at wall
		return false
	}

	adj := start.Add(direction)
	v := grid.PickFrom(area, adj)
	if v == "#" {
		return false // can not move in that direction
	}

	if v == "[" {
		if direction == point.WEST || direction == point.EAST {
			if !move2(area, adj, direction, justCheck) {
				return false
			}
		} else {
			if !move2(area, adj, direction, justCheck) || !move2(area, adj.Add(point.RIGHT), direction, justCheck) {
				return false
			}
		}
	}

	if v == "]" {
		if direction == point.WEST || direction == point.EAST {
			if !move2(area, adj, direction, justCheck) {
				return false
			}
		} else {
			if !move2(area, adj, direction, justCheck) || !move2(area, adj.Add(point.LEFT), direction, justCheck) {
				return false
			}
		}
	}

	if !justCheck {
		area[adj.Y][adj.X] = area[start.Y][start.X]
		area[start.Y][start.X] = "."
	}

	return true
}

func countGPS(area [][]string) int {
	result := 0
	for pos, v := range grid.IterateGrid(area) {
		if v == "O" || v == "[" {
			result += pos.Y*100 + pos.X
		}
	}
	return result
}

//func dumpGrid(area [][]string) {
//	for _, row := range area {
//		fmt.Println(strings.Join(row, ""))
//	}
//	fmt.Println()
//}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	var area [][]string
	if isPart2 {
		area, _ = utils.Map(strings.Split(p.Rows[0], "\n"), func(row string) ([]string, error) {
			cells := make([]string, len(row)<<1)
			for i, v := range strings.Split(row, "") {
				c1, c2 := ".", "."
				switch v {
				case "O":
					c1, c2 = "[", "]"
				case "#":
					c1, c2 = "#", "#"
				case "@":
					c1, c2 = "@", "."
				}
				cells[i<<1], cells[i<<1+1] = c1, c2
			}
			return cells, nil
		})
	} else {
		area, _ = utils.Map(strings.Split(p.Rows[0], "\n"), func(row string) ([]string, error) {
			return strings.Split(row, ""), nil
		})
	}
	cmds := strings.Split(p.Rows[1], "")

	robot := findRobot(area)
	//fmt.Printf("Initial state:\n")
	//dumpGrid(area)

	for _, cmd := range cmds {
		if isPart2 {
			if !move2(area, robot, point.FromDirection(cmd), true) {
				continue
			}
			if move2(area, robot, point.FromDirection(cmd), false) {
				robot = robot.Add(point.FromDirection(cmd))
			} else {
				panic("Hähh.. check works but for real not? :-O")
			}
		} else {
			if move(area, robot, point.FromDirection(cmd)) {
				robot = robot.Add(point.FromDirection(cmd))
			}
		}
		//fmt.Printf("Move %s:\n", cmd)
		//dumpGrid(area)
	}

	result = countGPS(area)
	return result
}
