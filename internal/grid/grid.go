package grid

import (
	"container/heap"
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/set"
	"iter"
	"slices"
)

type GridConfig struct {
	marker string
	wall   string
}
type GridConfigFunc func(cfg *GridConfig) *GridConfig

func NewGridConfig(options ...GridConfigFunc) *GridConfig {
	cfg := &GridConfig{marker: "O", wall: "#"}
	for _, fn := range options {
		cfg = fn(cfg)
	}
	return cfg
}

type Grid struct {
	cfg  *GridConfig
	data [][]string
}

func NewGrid(width, height int, options ...GridConfigFunc) Grid {
	data := make([][]string, height)
	for i := 0; i < height; i++ {
		data[i] = make([]string, width)
	}
	return Grid{
		cfg:  NewGridConfig(options...),
		data: data,
	}
}

func AsGrid(matrix [][]string, options ...GridConfigFunc) Grid {
	return Grid{
		data: matrix,
		cfg:  NewGridConfig(options...),
	}
}

func (g Grid) FindMarker(marker string) (point.Point, bool) {
	for pos, value := range IterateGrid(g.data) {
		if value == marker {
			return pos, true
		}
	}
	return point.Point{}, false
}

func (g Grid) All() iter.Seq2[point.Point, string] {
	return IterateGrid(g.data)
}

func (g Grid) Dump() {
	g.DumpWithMarker(set.NewSet[point.Point]())
}

func (g Grid) DumpWithMarker(markers set.Set[point.Point]) {
	for y, row := range g.data {
		for x, cell := range row {
			if markers.Contains(point.Point{x, y}) {
				fmt.Print(g.cfg.marker)
			} else {
				fmt.Print(cell)
			}
		}
		fmt.Println()
	}
}

func (g Grid) Bounds() rectangle.Rectangle {
	return GetBounds(g.data)
}

func (g Grid) Set(x, y int, value string) {
	g.data[y][x] = value
}

func (g Grid) Dijkstra(start, end point.Point) (int, []point.Point) {
	distances := make(map[point.Point]int)
	previous := make(map[point.Point]point.Point)
	visited := set.NewSet[point.Point]()
	bounds := g.Bounds()

	q := make(queue.PriorityQueue, 0)
	q.Push(queue.NewItem(start, 0))
	heap.Init(&q)

	for !q.Empty() {
		item := heap.Pop(&q).(*queue.Item)
		if item.Pos == end {
			return distances[item.Pos], buildPath(item.Pos, previous)
		}
		if visited.Contains(item.Pos) {
			continue
		}
		visited.Add(item.Pos)

		for _, adj := range item.Pos.DirectAdjacents() {
			if !bounds.Contains(adj) || visited.Contains(adj) {
				continue
			}

			if g.data[adj.Y][adj.X] == g.wall() {
				continue
			}
			costs := distances[item.Pos] + 1
			if v, ok := distances[adj]; ok && costs >= v {
				continue
			}
			distances[adj] = costs
			previous[adj] = item.Pos
			heap.Push(&q, queue.NewItem(adj, costs))
		}
	}
	return 0, nil
}

func (g Grid) wall() string {
	if g.cfg != nil {
		return g.cfg.wall
	}
	return "#"
}

func (g Grid) marker() string {
	if g.cfg != nil {
		return g.cfg.marker
	}
	return "O"
}

func buildPath(p point.Point, previous map[point.Point]point.Point) []point.Point {
	path := []point.Point{p}
	for {
		v, ok := previous[p]
		if !ok {
			break
		}
		path = append(path, v)
		p = v
	}
	slices.Reverse(path)
	return path
}
