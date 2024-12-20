package _0

import (
	"container/heap"
	"fmt"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/grid"
	"github.com/spossner/aoc2024/internal/pair"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
	"math"
	"slices"
)

const DEV_FILE = "dev2.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithSplitWords(""),
		config.WithDevFile(DEV_FILE),
	)
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
	if isPart2 {
		fmt.Println("PART2")
	}

	g := grid.AsGrid(p.Cells)
	start, _ := g.FindMarker("S")
	end, _ := g.FindMarker("E")
	solutions, costs := dijkstra(p.Cells, start, point.EAST, end)

	markers := set.NewSet[point.Point]()
	for _, solution := range solutions {
		for _, el := range solution {
			markers.Add(el.location)
		}
	}
	//grid.DumpGridWithMarker(p.Cells, markers)

	return pair.NewIntPair(costs, len(markers))
}

type Pos struct {
	location, direction point.Point
}

type Step struct {
	Pos
	costs int
}

type Item struct {
	Step
	path  []Pos
	index int // used by heap
}

func NewItem(location, direction point.Point, path []Pos, costs int) *Item {
	return &Item{Step: Step{Pos: Pos{location, direction}, costs: costs}, path: path}
}

func roationCosts(startDirection, endDirection point.Point) int {
	if startDirection == endDirection {
		return 0
	}
	if startDirection.RotateRight() == endDirection || startDirection.RotateLeft() == endDirection {
		return 1000
	}
	return 2000
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].costs < pq[j].costs
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func dijkstra(field [][]string, start, direction, end point.Point) ([][]Pos, int) {
	bounds := grid.GetBounds(field)
	distances := make(map[Pos]int)
	solutions := make([][]Pos, 0)
	costsSoFar := math.MaxInt

	q := make(PriorityQueue, 0)
	q.Push(NewItem(start, direction, []Pos{{start, direction}}, 0))
	heap.Init(&q)
	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)

		// check global solution costs already passed
		if item.costs > costsSoFar {
			continue
		}

		// check if there is already a cheaper path to this position + direction
		if v, ok := distances[item.Pos]; ok && item.costs > v {
			continue
		}

		if item.location == end {
			if item.costs > costsSoFar {
				continue
			}
			if item.costs < costsSoFar {
				solutions = make([][]Pos, 0)
			}
			costsSoFar = item.costs
			solutions = append(solutions, item.path)
			continue
		}

		for d, adj := range item.location.DirectAdjacents() {
			if d == point.OPPOSITE_DIRECTION[item.direction] { // never go back :-D
				continue
			}

			// do not leave field
			if !bounds.Contains(adj) {
				continue
			}

			// do not go into walls
			if field[adj.Y][adj.X] == "#" {
				continue
			}

			newPos := Pos{adj, d}

			costs := distances[item.Pos] + roationCosts(item.direction, d) + 1
			if currentCosts, ok := distances[newPos]; ok && costs > currentCosts {
				continue
			}
			distances[newPos] = costs
			heap.Push(&q, NewItem(adj, d, append(slices.Clone(item.path), newPos), costs))
		}
	}
	return solutions, costsSoFar
}
