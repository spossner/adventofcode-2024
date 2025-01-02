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

func TestLCM(t *testing.T) {
	type args struct {
		a    int
		b    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simple", args{3, 7, nil}, 21},
		{"larger", args{13, 283, nil}, 3679},
		{"three", args{12, 15, []int{10}}, 60},
		{"four", args{35, 48, []int{42, 100}}, 8400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.a, tt.args.b, tt.args.nums...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMod(t *testing.T) {
	type args[T constraints.Integer] struct {
		x T
		y T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"simple", args[int]{7, 3}, 1},
		{"matching", args[int]{9, 3}, 0},
		{"little", args[int]{3, 9}, 3},
		{"big", args[int]{372387203, 9}, 8},
		{"negative", args[int]{-5, 9}, 4},
		{"big negative", args[int]{-28975, 9}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mod(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Mod() = %v, want %v", got, tt.want)
			}
		})
	}
}
