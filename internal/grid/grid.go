package grid

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"github.com/spossner/aoc2024/internal/queue"
	"github.com/spossner/aoc2024/internal/rectangle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
	"iter"
	"strings"
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

type Path []point.Point

func (p Path) String() string {
	return strings.Join(utils.MustMap[point.Point, string](p, func(p point.Point) (string, error) {
		return p.Icon(), nil
	}), "")
}

var root = point.Point{-1, -1}

type BfsOptions struct {
	directions bool
}
type BfsOptionFunc func(cfg *BfsOptions) *BfsOptions

func WithDirections() BfsOptionFunc {
	return func(cfg *BfsOptions) *BfsOptions {
		cfg.directions = true
		return cfg
	}
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

func (g Grid) BfsAll(start, end point.Point, options ...BfsOptionFunc) []Path {
	cfg := &BfsOptions{}
	for _, fn := range options {
		cfg = fn(cfg)
	}
	q := queue.NewQueue[point.Point](start)
	previous := make(map[point.Point][]point.Point)
	previous[start] = nil
	distances := make(map[point.Point]int)
	distances[start] = 0
	for distance := 1; !q.Empty(); distance++ {
		for range q.Len() { // do one step after the other
			item, _ := q.PopLeft()
			if item == end {
				return buildMultiPath(item, previous, cfg.directions)
			}
			for _, adj := range item.DirectAdjacents() {
				if !g.Contains(adj) {
					continue
				}
				if g.Get(adj.X, adj.Y) == g.Wall() {
					continue
				}
				d, ok := distances[adj]
				if ok && d < distance {
					continue // already seen earlier
				}
				distances[adj] = distance
				previous[adj] = append(previous[adj], item)
				if len(previous[adj]) == 1 {
					q.Append(adj) // first visitor
				}
			}
		}
	}
	return nil
}

func (g Grid) Bfs(start, end point.Point) Path {
	q := queue.NewQueue[point.Point](start)
	previous := make(map[point.Point]point.Point)
	previous[start] = root
	for !q.Empty() {
		item, _ := q.PopLeft()
		if item == end {
			return buildPath(item, previous)
		}
		for _, adj := range item.DirectAdjacents() {
			if !g.Contains(adj) {
				continue
			}
			if g.Get(adj.X, adj.Y) == g.Wall() {
				continue
			}
			if utils.Contains(previous, adj) {
				continue
			}
			previous[adj] = item
			q.Append(adj)
		}
	}
	return nil
}

func (g Grid) Dijkstra(start, end point.Point) (int, []point.Point) {
	distances := make(map[point.Point]int)
	previous := make(map[point.Point]point.Point)
	previous[start] = root

	q := queue.NewPQ[point.Point]()
	q.Push(0, start)

	for !q.Empty() {
		item := q.Pop()
		if item == end {
			return distances[item], buildPath(item, previous)
		}

		for _, adj := range item.DirectAdjacents() {
			if !g.Contains(adj) {
				continue
			}

			if utils.Contains(previous, adj) {
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
