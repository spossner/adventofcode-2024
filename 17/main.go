package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/utils"
	"math"
	"reflect"
	"slices"
	"strconv"
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

type CPU struct {
	ptr     int
	a, b, c int
	output  []int
	goal    []int
	debug   bool
	halted  bool
}

func NewCPU(a int, goal []int) *CPU {
	return &CPU{a: a, output: make([]int, 0), goal: goal}
}

func (c *CPU) setPtr(newPtr int) *CPU {
	c.ptr = newPtr
	return c
}

// incPtr increases ptr by 2 skipping the next operandType
func (c *CPU) incPtr() *CPU {
	c.ptr += 2
	return c
}

func (c *CPU) comboValue(arg int) int {
	switch arg {
	case 0, 1, 2, 3:
		return arg
	case 4:
		return c.a
	case 5:
		return c.b
	case 6:
		return c.c
	default:
		panic("invalid operandType")
	}
}

func (c *CPU) adv(arg int) {
	c.a = c.a / int(math.Pow(2.0, float64(c.comboValue(arg))))
}

func (c *CPU) bdv(arg int) {
	c.b = c.a / int(math.Pow(2.0, float64(c.comboValue(arg))))
}

func (c *CPU) cdv(arg int) {
	c.c = c.a / int(math.Pow(2.0, float64(c.comboValue(arg))))
}

func (c *CPU) bxl(arg int) {
	c.b = c.b ^ arg
}

func (c *CPU) bxc() {
	c.b ^= c.comboValue(6)
}

func (c *CPU) bst(arg int) {
	c.b = utils.Mod(c.comboValue(arg), 8)
}

func (c *CPU) jnz(arg int) bool {
	if c.a == 0 {
		return false
	}
	c.ptr = arg
	return true
}

func (c *CPU) out(arg int) {
	v := c.comboValue(arg) % 8
	c.output = append(c.output, v)
	n := len(c.output)
	if c.goal != nil && (n > len(c.goal) || v != c.goal[n-1]) {
		c.halted = true
		return
	}
	//if len(c.output) > 1 {
	//	fmt.Print(",")
	//}
	//fmt.Print(v)
	// fmt.Printf("%s  -- %s\n", strings.Join(c.output, ","), c.dump())
}

func (c *CPU) dump() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *CPU) exec(op, arg int) *CPU {
	switch op {
	case 0:
		c.adv(arg)
		c.incPtr()
	case 1:
		c.bxl(arg)
		c.incPtr()
	case 2:
		c.bst(arg)
		c.incPtr()
	case 3:
		if jumped := c.jnz(arg); !jumped {
			c.incPtr()
		}
	case 4:
		c.bxc()
		c.incPtr()
	case 5:
		c.out(arg)
		c.incPtr()
	case 6:
		c.bdv(arg)
		c.incPtr()
	case 7:
		c.cdv(arg)
		c.incPtr()
	}
	// NOOP
	return c
}

func runCode(cfg *config.Config, a int, code []int) *CPU {
	cpu := NewCPU(a, code)
	for cpu.ptr < len(code) {
		//if cpu.halted {
		//	break
		//}
		op := code[cpu.ptr]
		arg := code[cpu.ptr+1]
		cpu = cpu.exec(op, arg)
	}
	return cpu
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	if isPart2 {
		fmt.Println("PART2")
	}

	initialValues, _ := utils.GetInts(utils.PickFirst(p.Rows))
	code, _ := utils.GetInts(utils.PickSecond(p.Rows))

	if isPart2 {
		//a := utils.PickFirst(initialValues)
		disassemble(code)
		//for a := 0; a < 1322205846300570; a++ {

		// GOAL
		// 2,4,1,5,7,5,1,6,4,1,5,5,0,3,3,0
		// 3,3,5,0,3,1,1,3,4,1,7,3,3,0,5,1
		//testCPU := runCode(cfg, 112756639541658, code)
		//fmt.Println("RUN CODE", testCPU.output)
		//for a > 0 {
		//	b := (a % 8) ^ 5
		//	c := a >> b
		//	b = b ^ 6
		//	b = b ^ c
		//	fmt.Printf("%d,", b%8)
		//	a = a >> 3
		//}
		//fmt.Println()
		if isPart2 {
			q := queue.NewQueue[int](0)
			for i := len(code) - 1; i >= 0; i-- {
				search := code[i]
				for range q.Len() {
					nextA, _ := q.PopLeft()
					nextA = nextA << 3
					for candidate := nextA; candidate < nextA+8; candidate++ {
						cpu := runCode(cfg, candidate, code)
						if len(cpu.output) > 0 && utils.PickFirst(cpu.output) == search {
							q.Append(candidate)
						}
					}
				}
			}
			fmt.Println(slices.Min(q.List()), q)
			return 0
		}

		for a := 8927907000000; a < 8928907000000; a++ {
			cpu := runCode(cfg, a, code)

			//if cpu.output[0] == code[0] && cpu.output[1] == code[1] && cpu.output[2] == code[2] {
			//
			//	fmt.Printf("%d: %+v\n", a, *cpu)
			//}

			if len(cpu.output) > len(code) {
				fmt.Println("a too high")
				return -1
			}
			if reflect.DeepEqual(cpu.output, code) {
				fmt.Println(" -> FOUND MATCH", a, cpu.dump())
				return a
			}
			a++
		}
		return "NOT FOUND"
	} else {
		cpu := runCode(cfg, utils.PickFirst(initialValues), code)
		return strings.Join(utils.Must(utils.Map(cpu.output, func(i int) (string, error) {
			return strconv.Itoa(i), nil
		})), ",")
	}
}

type operandType int

const (
	NONE operandType = iota
	LITERAL
	COMBO
)

type cmd struct {
	name    string
	comment string
	argType operandType
}

var COMMANDS = map[int]cmd{
	0: {"adv", "A = A / 2^%s", COMBO},
	1: {"bxl", "B = B xor %s", LITERAL},
	2: {"bst", "B = %s %% 8", COMBO},
	3: {"jnz", "JNZ(A) to %s", LITERAL},
	4: {"bxc", "B = B xor C", NONE},
	5: {"out", "print %s %% 8", COMBO},
	6: {"bdv", "B = A / 2^%s", COMBO},
	7: {"cdv", "C = A / 2^%s", COMBO},
}

func arg2string(c cmd, arg int) string {
	if c.argType == LITERAL || (arg >= 0 && arg <= 3) {
		return strconv.Itoa(arg)
	}
	switch arg {
	case 4:
		return "A"
	case 5:
		return "B"
	case 6:
		return "C"
	default:
		panic("illegal operandType")
	}
}

func comment(c cmd, arg int) string {
	if c.argType == NONE {
		return c.comment
	}

	return fmt.Sprintf(c.comment, arg2string(c, arg))
}

func disassemble(code []int) {
	for i := 0; i < len(code); i += 2 {
		op := code[i]
		arg := code[i+1]
		c := COMMANDS[op]
		fmt.Printf("%3d: %6s  %v  // %s\n", i, c.name, arg2string(c, arg), comment(c, arg))
	}
}
