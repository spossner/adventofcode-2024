package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
	"regexp"
	"slices"
	"strings"
)

const DEV_FILE = "dev2.txt"

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

type gateType int8

const (
	UNKNOWN_GATE gateType = iota
	AND_GATE
	OR_GATE
	XOR_GATE
)

var GATE_TYPES = map[string]gateType{
	"AND": AND_GATE,
	"OR":  OR_GATE,
	"XOR": XOR_GATE,
}

var TYPE_MAP = map[gateType]string{
	AND_GATE: "AND",
	OR_GATE:  "OR",
	XOR_GATE: "XOR",
}

type gate struct {
	in1, in2 string
	out      string
	typ      gateType
	value    int
	ready    bool
}

func (g *gate) reset() {
	g.value = 0
	g.ready = false
}

func (g *gate) String() string {
	if g.ready {
		return fmt.Sprintf("%s %s %s -> %s (%d)", g.in1, TYPE_MAP[g.typ], g.in2, g.out, g.value)
	}
	return fmt.Sprintf("%s %s %s -> %s", g.in1, TYPE_MAP[g.typ], g.in2, g.out)
}

func (g *gate) step(wires map[string]int) (int, bool) {
	if g.ready {
		return g.value, true
	}
	v1, ok := wires[g.in1]
	if !ok {
		return 0, false
	}
	v2, ok := wires[g.in2]
	if !ok {
		return 0, false
	}
	switch g.typ {
	case AND_GATE:
		g.value = v1 & v2
	case OR_GATE:
		g.value = v1 | v2
	case XOR_GATE:
		g.value = v1 ^ v2
	case UNKNOWN_GATE:
		panic("unknown gate type")
	}
	g.ready = true
	return g.value, true
}

func eval(gates []*gate, wires map[string]int) int {
	result := 0
	active := true
	for active {
		active = false
		for _, g := range gates {
			if g.ready {
				continue
			}
			if v, ok := g.step(wires); ok {
				active = true
				wires[g.out] = v
			}
		}
	}

	for k, v := range wires {
		if v == 0 || !strings.HasPrefix(k, "z") {
			continue
		}
		bit := utils.MustAtoi(k[1:])
		result |= (1 << bit)
	}
	return result
}

func parseData(p *puzzle.Puzzle) (map[string]int, []*gate) {
	wires := make(map[string]int)
	gates := make([]*gate, 0)
	for _, cable := range strings.Split(p.Rows[0], "\n") {
		parts := strings.Split(cable, ": ")
		wires[parts[0]] = utils.MustAtoi(parts[1])
	}
	re := regexp.MustCompile(`(.{3}) (AND|OR|XOR) (.{3}) -> (.{3})`)
	for _, wiring := range strings.Split(p.Rows[1], "\n") {
		matches := re.FindStringSubmatch(wiring)
		if len(matches) == 0 {
			panic("illegal wiring specification " + wiring)
		}
		gates = append(gates, &gate{
			in1: matches[1],
			in2: matches[3],
			out: matches[4],
			typ: GATE_TYPES[matches[2]],
		})
	}
	return wires, gates
}

func solve2(cfg *config.Config) any {
	p := puzzle.NewPuzzle(cfg)
	_, gates := parseData(p)
	gatesMap := make(map[string]*gate)
	for _, g := range gates {
		gatesMap[g.out] = g
	}

	wrongOutputs := set.NewSet[string]()
	for _, g := range gates {
		if g.out == "z45" {
			if g.typ != OR_GATE {
				fmt.Println("wrong last gate to", g.out, g)
				wrongOutputs.Add(g.out)
			}
		} else if g.out[0] == 'z' {
			if g.typ != XOR_GATE {
				fmt.Println("wrong gate connected to", g.out, g)
				wrongOutputs.Add(g.out)
			}
		} else {
			switch g.typ {
			case OR_GATE:
				if gatesMap[g.in1].typ != AND_GATE {
					fmt.Println("input 1: AND gate into OR gate is illegal", g)
					wrongOutputs.Add(g.in1)
				}
				if gatesMap[g.in2].typ != AND_GATE {
					fmt.Println("input 2: AND gate into OR gate is illegal", g)
					wrongOutputs.Add(g.in2)
				}
			case AND_GATE:
				if prevGate, ok := gatesMap[g.in1]; ok && prevGate.typ == AND_GATE && prevGate.in1 != "x00" && prevGate.in2 != "x00" {
					fmt.Println("input 1: two AND gates in a row:", prevGate, "=>", g)
					wrongOutputs.Add(g.in1)
				}
				if prevGate, ok := gatesMap[g.in2]; ok && prevGate.typ == AND_GATE && prevGate.in1 != "x00" && prevGate.in2 != "x00" {
					fmt.Println("input 2: two AND gates in a row:", prevGate, "=>", g)
					wrongOutputs.Add(g.in2)
				}
			case XOR_GATE:
				prev1, ok1 := gatesMap[g.in1]
				prev2, ok2 := gatesMap[g.in2]

				if ok1 && ok2 && (prev1.typ == XOR_GATE && prev2.typ == OR_GATE || prev2.typ == XOR_GATE && prev1.typ == OR_GATE) {
					if g.out[0] != 'z' {
						fmt.Println("invalid input gate to XOR:", prev1, "+", prev2, "=>", g)
						wrongOutputs.Add(g.out)
					}
				}
			default:

			}
		}
	}
	result := wrongOutputs.List()
	slices.Sort(result)
	return strings.Join(result, ",")
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	if isPart2 {
		fmt.Println("PART2")
		return solve2(cfg)
	}

	p := puzzle.NewPuzzle(cfg)
	wires, gates := parseData(p)
	return eval(gates, wires)
}
