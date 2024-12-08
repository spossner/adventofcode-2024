package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args[T any] struct {
		values []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"empty", args[int]{nil}, []int{}},
		{"single", args[int]{[]int{42}}, []int{42}},
		{"multiple", args[int]{[]int{42, 99, 100}}, []int{42, 99, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.values...); !reflect.DeepEqual(got.List(), tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got.List(), tt.want)
			}
		})
	}
}

func TestQueue_Append(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		q    *Queue[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"simple", NewQueue(1, 2, 3, 5, 8, 13), args[int]{21}, []int{1, 2, 3, 5, 8, 13, 21}},
		{"first", NewQueue[int](), args[int]{21}, []int{21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Append(tt.args.value)
			if !reflect.DeepEqual(tt.q.List(), tt.want) {
				t.Errorf("Append() = %v, want %v", tt.q.List(), tt.want)
			}
		})
	}
}

func TestQueue_AppendLeft(t *testing.T) {
	type args[T any] struct {
		value T
	}
	type testCase[T any] struct {
		name string
		q    *Queue[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"simple", NewQueue(1, 2, 3, 5, 8, 13), args[int]{21}, []int{21, 1, 2, 3, 5, 8, 13}},
		{"first", NewQueue[int](), args[int]{21}, []int{21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.AppendLeft(tt.args.value)
			if !reflect.DeepEqual(tt.q.List(), tt.want) {
				t.Errorf("AppendLeft() = %v, want %v", tt.q.List(), tt.want)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    *Queue[T]
	}
	tests := []testCase[byte]{
		{"simple", NewQueue[byte](2, 3, 5, 8, 13)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Clear()
			if !reflect.DeepEqual(tt.q, &Queue[byte]{}) {
				t.Errorf("Clear() does not create Queue zero value: %v", tt.q.List())
			}
		})
	}
}

func TestQueue_Extend(t *testing.T) {
	type args[T any] struct {
		values []T
	}
	type testCase[T any] struct {
		name string
		q    *Queue[T]
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{"simple", NewQueue(1, 2, 3, 5), args[int]{[]int{8, 13, 21}}, []int{1, 2, 3, 5, 8, 13, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.Extend(tt.args.values...)
			if !reflect.DeepEqual(tt.q.List(), tt.want) {
				t.Errorf("Extend() = %v, want %v", tt.q.List(), tt.want)
			}
		})
	}
}

func TestQueue_ExtendLeft(t *testing.T) {
	type args[T any] struct {
		values []T
	}
	type testCase[T any] struct {
		name string
		q    Queue[T]
		args args[T]
	}
	tests := []testCase[byte]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.ExtendLeft(tt.args.values...)
		})
	}
}

func TestQueue_Len(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    Queue[T]
		want int
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    Queue[T]
		want any
	}
	tests := []testCase[string]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type testCase[T any] struct {
		name  string
		q     Queue[T]
		want  T
		want1 bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_PopLeft(t *testing.T) {
	type testCase[T any] struct {
		name  string
		q     Queue[T]
		want  T
		want1 bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.PopLeft()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopLeft() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PopLeft() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}