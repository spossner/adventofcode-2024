package grid

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
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
	g.DumpWithMarker(nil)
}

func (g Grid) DumpPath(path []point.Point) {
	g.DumpWithMarker(set.FromSlice(path))
}

func (g Grid) DumpWithMarker(markers set.Set[point.Point]) {
	for y, row := range g.data {
		for x, cell := range row {
			if markers != nil && markers.Contains(point.Point{x, y}) {
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

func (g Grid) Contains(p point.Point) bool {
	if len(g.data) == 0 {
		return false
	}
	if p.X >= 0 && p.Y >= 0 && p.X < len(g.data[0]) && p.Y < len(g.data) {
		return true
	}
	return false
}

func (g Grid) Get(x, y int) string {
	return g.data[y][x]
}

func (g Grid) Set(x, y int, value string) {
	g.data[y][x] = value
}

func (g Grid) Bfs(start, end point.Point) []point.Point {
	root := point.Point{-1, -1}
	q := queue.NewQueue[point.Point](start)
	parents := make(map[point.Point]point.Point)
	parents[start] = root
	for !q.Empty() {
		item, _ := q.PopLeft()
		if item == end {
			var route []point.Point
			var ok bool
			for {
				route = append(route, item)
				item, ok = parents[item]
				if !ok || item == root {
					slices.Reverse(route)
					return route
				}
			}
		}
		for _, adj := range item.DirectAdjacents() {
			if !g.Contains(adj) {
				continue
			}
			if utils.Contains(parents, adj) {
				continue
			}
			if g.Get(adj.X, adj.Y) == g.Wall() {
				continue
			}
			parents[adj] = item
			q.Append(adj)
		}
	}
	return nil
}

func (g Grid) Dijkstra(start, end point.Point) (int, []point.Point) {
	distances := make(map[point.Point]int)
	previous := make(map[point.Point]point.Point)
	visited := set.NewSet[point.Point]()
	bounds := g.Bounds()

	q := queue.NewPQ[point.Point]()
	q.Push(0, start)

	for !q.Empty() {
		item := q.Pop()
		if item == end {
			return distances[item], buildPath(item, previous)
		}
		if visited.Contains(item) {
			continue
		}
		visited.Add(item)

		for _, adj := range item.DirectAdjacents() {
			if !bounds.Contains(adj) || visited.Contains(adj) {
				continue
			}

			if g.data[adj.Y][adj.X] == g.Wall() {
				continue
			}
			costs := distances[item] + 1
			if v, ok := distances[adj]; ok && costs >= v {
				continue
			}
			distances[adj] = costs
			previous[adj] = item
			q.Push(costs, adj)
		}
	}
	return 0, nil
}

func (g Grid) Wall() string {
	if g.cfg != nil {
		return g.cfg.wall
	}
	return "#"
}

func (g Grid) Marker() string {
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
