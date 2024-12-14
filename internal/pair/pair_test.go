package pair

import (
	"reflect"
	"slices"
	"testing"
)

func TestNewIntPair(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want Pair[int]
	}{
		{"simple", args{3, 4}, Pair[int]{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntPair(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_All(t *testing.T) {
	type testCase[T any] struct {
		name string
		p    Pair[T]
		want []int
	}
	tests := []testCase[int]{
		{"simple", Pair[int]{42, 73}, []int{42, 73}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Collect(tt.p.All()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPair_Pick(t *testing.T) {
	type testCase[T any] struct {
		name  string
		p     Pair[T]
		want  T
		want1 T
	}
	tests := []testCase[int]{
		{"simple", Pair[int]{1, 3}, 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.Pick()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pick() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pick() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
