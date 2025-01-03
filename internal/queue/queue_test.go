package queue

import (
	"github.com/stretchr/testify/assert"
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
		{"a lot", NewQueue(0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584), args[int]{4181}, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}},
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

func TestQueue_RingBuffer(t *testing.T) {
	q := NewQueue[int]()

	assert.Equal(t, 0, q.length, "queue should have a length of 10.000")
	assert.Equal(t, 0, q.head, "head should be at zero")
	assert.Equal(t, 0, q.tail, "tail should be at 10.000")
	assert.Equal(t, 0, q.Peek(), "peek non existing first element should return zero value of int")
	assert.Equal(t, 0, q.PeekLast(), "peek non existing last element should return zero value of int")
	assert.Equal(t, 0, len(q.data), "ring buffer should be nil")

	for i := range 10_000 {
		q.Append(i)
	}
	assert.Equal(t, 10_000, q.length, "queue should have a length of 10.000")
	assert.Equal(t, 0, q.head, "head should be at zero")
	assert.Equal(t, 10_000, q.tail, "tail should be at 10.000")
	assert.Equal(t, 0, q.Peek(), "peek first element should be 0")
	assert.Equal(t, 9_999, q.PeekLast(), "peek last element should be 9.999")
	assert.Equal(t, 16_384, len(q.data), "ring buffer should have a length of 2^14")

	q.Clear()
	assert.Equal(t, 0, q.length, "queue should have a length of 10.000")
	assert.Equal(t, 0, q.head, "head should be at zero")
	assert.Equal(t, 0, q.tail, "tail should be at 10.000")
	assert.Equal(t, 0, q.Peek(), "peek non existing first element should return zero value of int")
	assert.Equal(t, 0, q.PeekLast(), "peek non existing last element should return zero value of int")
	assert.Equal(t, 0, len(q.data), "ring buffer should be nil")
}

func TestQueue_PoppingAllItems(t *testing.T) {
	q := NewQueue(1, 2, 3, 4, 5)
	assert.Equal(t, 5, q.length, "length should be 5")
	assert.Equal(t, 8, len(q.data), "capacity should be 8")
	assert.Equal(t, 0, q.head, "head should point to first element of ring buffer")
	assert.Equal(t, 5, q.tail, "tail should point behind the last element")

	q.PopLeft()
	assert.Equal(t, 4, q.length, "length should now be 4")
	assert.Equal(t, 8, len(q.data), "capacity should still be 8")
	assert.Equal(t, 1, q.head, "head should point to the new first element")
	assert.Equal(t, 5, q.tail, "tail should point behind the last element")

	q.Pop()
	assert.Equal(t, 3, q.length, "length should now be 3")
	assert.Equal(t, 8, len(q.data), "capacity should still be 8")
	assert.Equal(t, 1, q.head, "head should point still to the first element")
	assert.Equal(t, 4, q.tail, "tail should now point behind the 4th element of the ring buffer")

	for range 3 {
		q.PopLeft()
	}

	assert.Equal(t, 0, q.length, "length should now be 0")
	assert.Equal(t, 8, len(q.data), "capacity should still be 8")
	assert.Equal(t, 4, q.head, "head points to 4")
	assert.Equal(t, 4, q.tail, "tail still points to 4")

	q.Clear()
	assert.Equal(t, 0, q.length, "length should now be 0")
	assert.Equal(t, 0, len(q.data), "capacity should be 0")
	assert.Equal(t, 0, q.head, "head points to 0")
	assert.Equal(t, 0, q.tail, "tail still points to 0")
}

func TestQueue_FloatingAround(t *testing.T) {
	q := NewQueue(1, 2, 3, 4, 5)
	q.PopLeft()
	q.Pop()
	assert.Equal(t, 3, q.length, "length should now be 3")
	assert.Equal(t, 8, len(q.data), "capacity should still be 8")
	assert.Equal(t, 1, q.head, "head should point still to the first element")
	assert.Equal(t, 4, q.tail, "tail should now point behind the 4th element of the ring buffer")

	for i := 6; i < 200; i++ {
		q.PopLeft()
		q.Append(i)
		assert.Equal(t, 3, q.length, "length should now be 3")
		assert.Equal(t, 8, len(q.data), "capacity should still be 8")
		assert.Equal(t, i, q.PeekLast(), "last element should be %d", i)
	}
}

func TestQueue_GrowBuffer(t *testing.T) {
	q := NewQueue(1, 2, 3, 4, 5)
	for range 3 {
		q.PopLeft()
	}

	for i := 6; i < 12; i++ {
		q.Append(i)
	}

	assert.Equal(t, 8, q.length, "queue should be full")
	assert.Equal(t, 8, len(q.data), "queue should still have capactity of 8")

	q.Append(12)
	assert.Equal(t, 9, q.length, "queue now has length of 9")
	assert.Equal(t, 16, len(q.data), "queue now grew to 16")
	assert.Equal(t, 4, q.Peek(), "first number should be 4")
	assert.Equal(t, 12, q.PeekLast(), "last number should be 12")
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
				t.Errorf("Clear() does not create Queue zero Value: %v", tt.q.List())
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
		q    *Queue[T]
		args args[T]
		want []T
	}
	tests := []testCase[byte]{
		{"simple",
			NewQueue[byte](0x23, 0x40, 0x6c),
			args[byte]{
				[]byte{0x01},
			},
			[]byte{0x01, 0x23, 0x40, 0x6c},
		},
		{"multi",
			NewQueue[byte](0x23, 0x40, 0x6c),
			args[byte]{
				[]byte{0x09, 0x05, 0x01},
			},
			[]byte{0x01, 0x05, 0x09, 0x23, 0x40, 0x6c},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.ExtendLeft(tt.args.values...)
			if !reflect.DeepEqual(tt.q.List(), tt.want) {
				t.Errorf("ExtendLeft() = %v, want %v", tt.q.List(), tt.want)
			}
		})
	}
}

func TestQueue_Len(t *testing.T) {
	type testCase[T any] struct {
		name string
		q    *Queue[T]
		want int
	}
	tests := []testCase[int]{
		{"simple", NewQueue[int](1, 2, 3), 3},
		{"empty", NewQueue[int](), 0},
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
		q    *Queue[T]
		want T
	}
	tests := []testCase[string]{
		{"simple", NewQueue[string]("Seppo", "Vera", "Carlotta"), "Seppo"},
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
		q     *Queue[T]
		want  T
		found bool
	}
	tests := []testCase[int]{
		{"simple", NewQueue[int](1, 2, 3, 4, 5), 5, true},
		{"none", NewQueue[int](), 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.Pop()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if ok != tt.found {
				t.Errorf("Pop() ok = %v, want %v", ok, tt.found)
			}
		})
	}
}

func TestQueue_PopLeft(t *testing.T) {
	type testCase[T any] struct {
		name  string
		q     *Queue[T]
		want  T
		found bool
	}
	tests := []testCase[int]{
		{"simple", NewQueue[int](1, 2, 3, 4, 5), 1, true},
		{"none", NewQueue[int](), 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := tt.q.PopLeft()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopLeft() got = %v, want %v", got, tt.want)
			}
			if ok != tt.found {
				t.Errorf("PopLeft() got1 = %v, want %v", ok, tt.found)
			}
		})
	}
}
