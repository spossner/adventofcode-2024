package utils

import (
	"reflect"
	"strconv"
	"testing"
)

func TestAny(t *testing.T) {
	type args[S interface{ ~[]E }, E any] struct {
		s S
		f func(E) bool
	}
	type testCase[S interface{ ~[]E }, E any] struct {
		name string
		args args[S, E]
		want bool
	}
	tests := []testCase[[]int, int]{
		{"simple",
			args[[]int, int]{
				[]int{1, 2, 3, 5, 8, 13, 21},
				func(v int) bool {
					return v == 13
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCut(t *testing.T) {
	type args[S ~[]T, T any] struct {
		slice S
		index int
	}
	type testCase[S ~[]T, T any] struct {
		name string
		args args[S, T]
		want S
	}
	tests := []testCase[[]int, int]{
		{"simple",
			args[[]int, int]{
				[]int{1, 2, 3, 5, 8, 13, 21},
				3,
			},
			[]int{1, 2, 3, 8, 13, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.args.slice, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[S ~[]T, T any] struct {
		slice S
		f     func(T) (bool, error)
	}
	type testCase[S ~[]T, T any] struct {
		name    string
		args    args[S, T]
		want    S
		wantErr bool
	}
	tests := []testCase[[]int, int]{
		{"simple",
			args[[]int, int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				func(i int) (bool, error) { return i%2 == 0, nil },
			},
			[]int{2, 4, 6, 8, 10},
			false,
		},
		{"none",
			args[[]int, int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				func(i int) (bool, error) { return false, nil },
			},
			[]int{},
			false,
		},
		{"all",
			args[[]int, int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				func(i int) (bool, error) { return true, nil },
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filter(tt.args.slice, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		slice []T
		f     func(T) (U, error)
	}
	type testCase[T any, U any] struct {
		name    string
		args    args[T, U]
		want    []U
		wantErr bool
	}
	tests := []testCase[string, int]{
		{"simple int conversion",
			args[string, int]{
				[]string{"1", "13", "42", "99"},
				func(s string) (int, error) {
					return strconv.Atoi(s)
				},
			},
			[]int{1, 13, 42, 99},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Map(tt.args.slice, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	type args[S ~[][]T, T any] struct {
		slice S
	}
	type testCase[S ~[][]T, T any] struct {
		name string
		args args[S, T]
		want S
	}
	tests := []testCase[[][]int, int]{
		{"simple", args[[][]int, int]{
			[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
			},
		},
			[][]int{
				{1, 6, 11, 16},
				{2, 7, 12, 17},
				{3, 8, 13, 18},
				{4, 9, 14, 19},
				{5, 10, 15, 20},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transpose(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
