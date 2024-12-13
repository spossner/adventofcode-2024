package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev, config.WithSplitWords("\n\n"), config.WithGetInts(), config.WithDevFile(DEV_FILE))
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	for _, row := range p.ParsedCells {
		A := point.FromValues(row...)
		B := point.FromValues(row[2:]...)
		P := point.FromValues(row[4:]...)

		if isPart2 {
			P.X += 10000000000000
			P.Y += 10000000000000
		}

		// cramer - thanks to reddit discussions... https://www.youtube.com/watch?v=Yr9hTPvl8Ng
		D := A.X*B.Y - B.X*A.Y
		Dx := P.X*B.Y - B.X*P.Y
		Dy := A.X*P.Y - P.X*A.Y

		// let panic when D == 0
		pressA := Dx / D
		pressB := Dy / D

		// only accept solutions with whole number presses (integers) for pressA and pressB
		if float64(pressA) != float64(Dx)/float64(D) || float64(pressB) != float64(Dy)/float64(D) {
			continue
		}

		result += (Dx/D)*3 + (Dy / D)
	}

	return result
}
