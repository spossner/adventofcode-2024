package queue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPQ(t *testing.T) {
	type args[U any] struct {
		values []Item[int, string]
	}
	type testCase[U any] struct {
		name string
		args args[U]
		want PriorityQueue[int, U]
	}
	tests := []testCase[string]{
		{"simple", args[string]{
			[]Item[int, string]{{"initial", 0, 0}},
		}, PriorityQueue[int, string]{{"initial", 0, 0}}},
		{"multi", args[string]{
			[]Item[int, string]{{"a", 10, 0}, {"b", 5, 0}, {"c", 3, 0}},
		}, PriorityQueue[int, string]{{"c", 3, 0}, {"b", 5, 0}, {"a", 10, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPQ(tt.args.values...); !reflect.DeepEqual(got.queue, tt.want) {
				for _, item := range got.queue {
					fmt.Println(*item)
				}
				t.Errorf("NewPQ() = %v, want %v", got.queue, tt.want)
			}
		})
	}
}
