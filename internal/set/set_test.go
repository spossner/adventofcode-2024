package set

import (
	"reflect"
	"testing"
)

func TestFromSlice(t *testing.T) {
	type args[T comparable] struct {
		slices [][]T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{"no sets", args[int]{[][]int{}}, Set[int]{}},
		{"simple", args[int]{[][]int{{1, 2, 3, 5, 8, 13, 21, 34}}}, Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 8: struct{}{}, 13: struct{}{}, 21: struct{}{}, 34: struct{}{}}},
		{"two sets", args[int]{[][]int{{1, 2, 3, 5}, {8, 13, 21, 34}}}, Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 8: struct{}{}, 13: struct{}{}, 21: struct{}{}, 34: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSlice(tt.args.slices...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSet(t *testing.T) {
	type args[T comparable] struct {
		items []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{"none", args[string]{[]string{}}, Set[string]{}},
		{"simple", args[string]{[]string{"Seppo"}}, Set[string]{"Seppo": struct{}{}}},
		{"double", args[string]{[]string{"Seppo", "Vera"}}, FromSlice([]string{"Seppo", "Vera"})},
		{"multi", args[string]{[]string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E"}}, FromSlice([]string{"a", "A", "b", "B", "c", "C", "d", "D", "e", "E"})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		item T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[float64]{
		{"simple", Set[float64]{1.2: struct{}{}, 2.3: struct{}{}}, args[float64]{3.14}, Set[float64]{1.2: struct{}{}, 2.3: struct{}{}, 3.14: struct{}{}}},
		{"already exists", Set[float64]{1.2: struct{}{}, 2.3: struct{}{}}, args[float64]{2.3}, Set[float64]{1.2: struct{}{}, 2.3: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.item)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("NewSet() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		item T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{"empty set", Set[string]{}, args[string]{"Seppo"}, false},
		{"empty item with empty set", Set[string]{}, args[string]{""}, false},
		{"found empty item", Set[string]{"": struct{}{}}, args[string]{""}, true},
		{"success", Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}, args[string]{"Seppo"}, true},
		{"does not exists", Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}, args[string]{"Carlotta"}, false},
		{"case sensitive mismatch", Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}, args[string]{"seppo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersect(t *testing.T) {
	type args[T comparable] struct {
		sets []Set[T]
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{"equal sets",
			args[int]{
				[]Set[int]{
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
				},
			}, NewSet(1, 2, 3, 5, 8, 13, 21, 34),
		},
		{"all items in all sets",
			args[int]{
				[]Set[int]{
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
				},
			}, NewSet(1, 2, 3, 5, 8, 13, 21, 34),
		},
		{"some items",
			args[int]{
				[]Set[int]{
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(1, 2, 34),
					NewSet(1, 34),
				},
			}, NewSet(1, 34),
		},
		{"no intersection",
			args[int]{
				[]Set[int]{
					NewSet(1, 2, 3, 5, 8, 13, 21, 34),
					NewSet(10, 20, 30),
					NewSet(340),
				},
			}, Set[int]{},
		},
		{"including empty set",
			args[int]{
				[]Set[int]{
					{},
					{10: struct{}{}, 20: struct{}{}, 30: struct{}{}},
					{340: struct{}{}}},
			}, Set[int]{},
		},
		{"intersect with empty set",
			args[int]{
				[]Set[int]{
					{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 8: struct{}{}, 13: struct{}{}, 21: struct{}{}, 34: struct{}{}},
					{},
				},
			},
			Set[int]{},
		},
		{"single set",
			args[int]{
				[]Set[int]{
					{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 8: struct{}{}, 13: struct{}{}, 21: struct{}{}, 34: struct{}{}},
				},
			},
			Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 5: struct{}{}, 8: struct{}{}, 13: struct{}{}, 21: struct{}{}, 34: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		item T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{"simple", Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}, args[string]{"Seppo"}, Set[string]{"Vera": struct{}{}}},
		{"not in set", Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}, args[string]{"Carlotta"}, Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.item)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Subtract(t *testing.T) {
	type args[T comparable] struct {
		others []Set[T]
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[byte]{
		{"simple", Set[byte]{0x32: struct{}{}, 0xff: struct{}{}}, args[byte]{[]Set[byte]{{0x32: struct{}{}}}}, Set[byte]{0xff: struct{}{}}},
		{"nothing to subtract", Set[byte]{0x32: struct{}{}, 0xff: struct{}{}}, args[byte]{}, Set[byte]{0x32: struct{}{}, 0xff: struct{}{}}},
		{"subtract multiple sets",
			Set[byte]{0x20: struct{}{}, 0x23: struct{}{}, 0x32: struct{}{}, 0x42: struct{}{}, 0xaa: struct{}{}, 0xff: struct{}{}},
			args[byte]{
				[]Set[byte]{
					{0x23: struct{}{}, 0xaa: struct{}{}},
					{0x42: struct{}{}},
					{0x20: struct{}{}},
				},
			},
			Set[byte]{0x32: struct{}{}, 0xff: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.s, tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args[T comparable] struct {
		sets []Set[T]
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want Set[T]
	}
	tests := []testCase[string]{
		{"simple",

			args[string]{
				[]Set[string]{
					{"Seppo": struct{}{}},
					{"Vera": struct{}{}},
				}},
			Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}},
		},

		{"multiple",
			args[string]{
				[]Set[string]{
					{"Seppo": struct{}{}},
					{"Vera": struct{}{}},
					{"Carlotta": struct{}{}},
					{"Emilia": struct{}{}, "Ami": struct{}{}},
				}},
			Set[string]{"Seppo": struct{}{}, "Vera": struct{}{}, "Carlotta": struct{}{}, "Emilia": struct{}{}, "Ami": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Pop(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want T
	}
	tests := []testCase[int]{
		{"simple", NewSet[int](1, 2, 3, 5, 8, 13, 21), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Pop()
			if len(tt.s) != tt.want {
				t.Errorf("Pop() = %v leaving %v elements in set", got, tt.want)
			}
		})
	}
}

func TestSet_MultiContains(t *testing.T) {
	type args[T comparable] struct {
		first T
		items []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{"simple", NewSet[int](1, 2, 3, 5, 8, 13), args[int]{1, []int{3, 5, 8}}, true},
		{"some", NewSet[int](1, 2, 3, 5, 8, 13), args[int]{1, []int{2, 3, 5, 8, 13}}, true},
		{"some not", NewSet[int](1, 2, 3, 5, 8, 13), args[int]{3, []int{5, 13, 21, 34}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.first, tt.args.items...); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
