package _0

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/interval"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/utils"
	"slices"
	"strconv"
)

type BlockType int

const (
	File = iota
	Space
)

type Block struct {
	interval interval.Interval
	id       int
	t        BlockType
}

func (b *Block) Len() int {
	return b.interval.Len()
}

func (b *Block) String() string {
	n := b.interval.Len()
	c := byte('.')
	if b.t == File {
		c = []byte(strconv.Itoa(b.id))[0]
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = c
	}
	return string(buf)
}

func NewBlock(from, to, id int, t BlockType) *Block {
	return &Block{interval.Interval{from, to}, id, t}
}

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitWords(""),
		config.WithGetInts(),
		config.WithDevFile("dev.txt"))
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve2(createConfig(dev))
}

func parseBlocks(p *puzzle.Puzzle) (blocks []*Block, left int, right int) {
	blocks = make([]*Block, 0)
	ptr := 0
	id := 0
	left = 1
	for i, n := range p.ParsedRows {
		if i%2 == 0 {
			blocks = append(blocks, NewBlock(ptr, ptr+n, id, File))
			id++
			right = i
		} else {
			blocks = append(blocks, NewBlock(ptr, ptr+n, -1, Space))
		}
		ptr += n
	}
	return blocks, left, right
}

func checksum(blocks []*Block) int {
	result := 0
	ptr := 0
	for _, block := range blocks {
		if block.t == Space {
			ptr += block.Len()
			continue
		}
		for range block.Len() {
			result += (ptr * block.id)
			ptr++
		}
	}
	return result
}

func solve2(cfg *config.Config) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	fmt.Println("PART2")

	//fmt.Println(p.ParsedRows)

	blocks, _, right := parseBlocks(p)
Outer:
	for ; right > 0; right-- {
		if blocks[right].t == Space {
			continue
		}
		//fmt.Println(blocks)
		lenRight := blocks[right].Len()
		for i := 0; i < right; i++ {
			if blocks[i].t == File {
				continue
			}
			if lenRight == blocks[i].Len() { // matches exactly (easy)
				blocks[i].id = blocks[right].id
				blocks[i].t = File

				blocks[right].id = -1
				blocks[right].t = Space
				continue Outer
			} else if blocks[i].Len() >= lenRight {
				newBlock := NewBlock(blocks[i].interval.From, blocks[i].interval.From+lenRight, blocks[right].id, File)
				blocks[i].interval.From = blocks[i].interval.From + lenRight
				blocks[right].id = -1
				blocks[right].t = Space
				blocks = slices.Insert(blocks, i, newBlock)
				continue Outer
			}
		}
	}
	//fmt.Println(blocks)
	result = checksum(blocks)
	return result
}

func solve(cfg *config.Config, isPart2 bool) any {
	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	fmt.Println(p.ParsedRows)

	blocks, left, right := parseBlocks(p)

	for left < right {
		//fmt.Println(blocks)
		lenLeft := blocks[left].interval.Len()
		lenRight := blocks[right].interval.Len()

		if lenRight == lenLeft {
			blocks[left].id = blocks[right].id
			blocks[left].t = File

			blocks[right].id = -1
			blocks[right].t = Space

			left += 2
			right -= 2

		} else if lenRight < lenLeft {
			newBlock := NewBlock(blocks[left].interval.From, blocks[left].interval.From+lenRight, blocks[right].id, File)
			blocks[left].interval.From = blocks[left].interval.From + lenRight
			blocks[right].id = -1
			blocks[right].t = Space
			blocks = slices.Insert(blocks, left, newBlock)
			left += 1
			right -= 1
		} else {
			blocks[left].id = blocks[right].id
			blocks[left].t = File
			l := blocks[left].interval.Len()

			newBlock := NewBlock(blocks[right].interval.To-l, blocks[right].interval.To, -1, Space)
			blocks[right].interval = interval.Interval{blocks[right].interval.From, blocks[right].interval.To - l}
			blocks = slices.Insert(blocks, right+1, newBlock)
			left += 2 // next free block
		}
	}

	//fmt.Println(blocks)

	result = checksum(blocks)

	return result
}
