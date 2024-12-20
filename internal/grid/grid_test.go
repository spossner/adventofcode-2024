package grid

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/point"
	"reflect"
	"strings"
	"testing"
)

func Test_buildPath(t *testing.T) {
	type args struct {
		p        point.Point
		previous map[point.Point]point.Point
	}
	tests := []struct {
		name string
		args args
		want []point.Point
	}{
		{"simple", args{p: point.Point{2, 2}, previous: map[point.Point]point.Point{
			{2, 2}: {1, 1},
			{1, 1}: {0, 0},
		}}, []point.Point{{0, 0}, {1, 1}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildPath(tt.args.p, tt.args.previous); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Dijkstra(t *testing.T) {
	type args struct {
		start point.Point
		end   point.Point
	}
	type testCase struct {
		name  string
		g     Grid
		args  args
		want  int
		want1 []point.Point
	}
	tests := []testCase{
		{"simple", Grid{data: [][]string{
			strings.Split("...#...", ""),
			strings.Split("..#..#.", ""),
			strings.Split("....#..", ""),
			strings.Split("...#..#", ""),
			strings.Split("..#..#.", ""),
			strings.Split(".#..#..", ""),
			strings.Split("#.#....", ""),
		},
		}, args{point.Point{0, 0}, point.Point{6, 6}}, 22, []point.Point{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.Dijkstra(tt.args.start, tt.args.end)
			if got != tt.want {
				t.Errorf("Dijkstra() got = %v, want %v", got, tt.want)
			}
			fmt.Println("path", got1)
			//if !reflect.DeepEqual(got1, tt.want1) {
			//	t.Errorf("Dijkstra() got1 = %v, want %v", got1, tt.want1)
			//}
		})
	}
}
