package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/rectangle"
)

const DAY = 4

func part1(dev bool) any {
	cfg := config.NewConfig(DAY, dev, config.WithSplitWords(""), config.WithDevFile("dev2.txt")) // split lines only
	return solve(cfg, false)
}

func part2(dev bool) any {
	cfg := config.NewConfig(DAY, dev, config.WithSplitWords(""), config.WithDevFile("dev2.txt")) // split lines only, setting another dev file
	p := puzzle.NewPuzzle(cfg)
	result := 0

	bounds := rectangle.NewRectangle(0, 0, len(p.Cells[0]), len(p.Cells))

	for y, row := range p.Cells {
		if y == 0 || y == bounds.Height-1 {
			continue
		}
		for x, c := range row {
			if x == 0 || x == bounds.Width-1 {
				continue
			}
			if c != "A" {
				continue
			}

			tl := p.Cells[y-1][x-1]
			br := p.Cells[y+1][x+1]

			tr := p.Cells[y-1][x+1]
			bl := p.Cells[y+1][x-1]

			if (tl == "M" && br == "S" || tl == "S" && br == "M") &&
				(tr == "M" && bl == "S" || tr == "S" && bl == "M") {
				result++
			}
		}
	}

	return result
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Print("PART2")
	}

	word := "XMAS"
	bounds := rectangle.NewRectangle(0, 0, len(p.Cells[0]), len(p.Cells))

	fmt.Println(p.Cells)
	for y, row := range p.Cells {
		for x, c := range row {
			if c != string(word[0]) {
				continue
			}

		Outer:
			for _, direction := range point.ADJACENT_POINTS {
				pos := point.Point{x, y}

				for _, cNext := range word[1:] {
					pos = pos.Translate(direction.X, direction.Y)
					if !bounds.Contains(pos) {
						continue Outer
					}
					if string(cNext) != p.Cells[pos.Y][pos.X] {
						continue Outer
					}
				}
				// fmt.Println("found XMAS at", x, y, "->", pos.X, pos.Y)
				result++
			}
		}
	}

	return result
}