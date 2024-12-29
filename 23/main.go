package _0

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/spossner/aoc2024/internal/config"
	"github.com/spossner/aoc2024/internal/puzzle"
	"github.com/spossner/aoc2024/internal/set"
	"github.com/spossner/aoc2024/internal/utils"
	"golang.org/x/exp/maps"
	"log"
	"slices"
	"strings"
)

const DEV_FILE = "dev.txt"

func createConfig(dev bool) *config.Config {
	return config.NewConfig(utils.GetPackageDir(), dev,
		config.WithSplitLines(),
		config.WithDevFile(DEV_FILE),
	)
}

func part1(dev bool) any {
	return solve(createConfig(dev), false)
}

func part2(dev bool) any {
	return solve(createConfig(dev), true)
}

func BronKerbosch(adj map[string]set.Set[string], R, P, X set.Set[string]) set.Set[string] {
	if len(P) == 0 && len(X) == 0 {
		return R // R maximal clique
	}
	var maxClique set.Set[string]
	for len(P) > 0 {
		v := P.Pop()
		clique := BronKerbosch(adj, R.Clone(v), set.Intersect(P, adj[v]), set.Intersect(X, adj[v]))
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
		X.Add(v)
	}
	return maxClique
}

func solve(cfg *config.Config, isPart2 bool) any {
	defer utils.Duration(fmt.Sprintf("DAY %d, PART %d", cfg.Day, utils.If(isPart2, 2, 1)))()

	p := puzzle.NewPuzzle(cfg)
	result := 0
	if isPart2 {
		fmt.Println("PART2")
	}

	g := graph.New(graph.StringHash)
	for _, row := range p.Rows {
		nodes := strings.Split(row, "-")
		if len(nodes) != 2 {
			log.Panicf("illegal input data %v", row)
		}
		_ = g.AddVertex(nodes[0])
		_ = g.AddVertex(nodes[1])
		_ = g.AddEdge(nodes[0], nodes[1])
	}

	adjMap := utils.Must(g.AdjacencyMap())
	if isPart2 {
		adj := make(map[string]set.Set[string])
		vertices := set.NewSet[string]()
		for node, neighbours := range adjMap {
			adj[node] = set.FromSlice(maps.Keys(neighbours))
			vertices.Add(node)
		}
		clique := BronKerbosch(adj, set.NewSet[string](), vertices, set.NewSet[string]())
		result := clique.List() // shadowing result as []string
		slices.Sort(result)
		return strings.Join(result, ",")
	}

	type triangle [3]string
	triangles := set.NewSet[triangle]()
	for node, adj := range adjMap {
		for a := range adj {
			if a == node {
				continue
			}
			for b := range adj {
				if a == b || b == node {
					continue
				}

				if utils.Contains(adjMap[a], b) {
					if a[0] == 't' || b[0] == 't' || node[0] == 't' {
						t := triangle{a, b, node}
						slices.Sort(t[:])
						triangles.Add(t)
					}
				}

			}
		}
	}
	result = len(triangles)

	return result
}
