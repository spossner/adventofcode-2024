package utils

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{18, 8}, 2},
		{"min", args{12, 3}, 3},
		{"standard", args{12, 18}, 6},
		{"none", args{4, 9}, 1},
		{"high", args{42, 70}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	type testCase[T constraints.Ordered] struct {
		name               string
		a, low, high, want T
	}
	tests := []testCase[int]{
		{"simple", 8, 6, 9, 8},
		{"lower", 3, 6, 9, 6},
		{"higher", 30, 6, 9, 9},
		{"lower bound", 6, 6, 9, 6},
		{"upper bound", 9, 6, 9, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bound(tt.a, tt.low, tt.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bound() = %v, want %v", got, tt.want)
			}
		})
	}
}
