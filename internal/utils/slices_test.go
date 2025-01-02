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

func TestBatched(t *testing.T) {
	type args[S ~[]E, E any] struct {
		s S
		n int
	}
	type testCase[S ~[]E, E any] struct {
		name string
		args args[S, E]
		want []S
	}
	tests := []testCase[[]byte, byte]{
		{"simple", args[[]byte, byte]{[]byte("Sebastian"), 3}, [][]byte{[]byte("Seb"), []byte("ast"), []byte("ian")}},
		{"not aligned", args[[]byte, byte]{[]byte("Sebastian"), 4}, [][]byte{[]byte("Seba"), []byte("stia"), []byte("n")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, batch := range Batched(tt.args.s, tt.args.n) {
				if !reflect.DeepEqual(batch, tt.want[i]) {
					t.Errorf("Batched()[%d] = %v, want %v", i, batch, tt.want[i])
				}
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T, U any] struct {
		slice   []T
		fn      func(U, T) U
		initial U
	}
	type testCase[T, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, int]{
		{"sum", args[int, int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc, item int) int { return acc + item }, 0}, 55},
		{"square", args[int, int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(acc, item int) int { return acc + (item * item) }, 0}, 385},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.slice, tt.args.fn, tt.args.initial); got != tt.want {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt(t *testing.T) {
	type args[T any] struct {
		slice []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{"simple", args[int]{[]int{1, 2, 3}}, 6},
		{"negative", args[int]{[]int{-1, 0, 1}}, 0},
		{"large", args[int]{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat(t *testing.T) {
	type args[T any] struct {
		slice []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[float64]{
		{"simple", args[float64]{[]float64{1.1, 2.2, 3.3}}, 6.6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.slice); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
