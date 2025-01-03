package queue

import (
	"fmt"
	"github.com/spossner/aoc2024/internal/utils"
	"iter"
	"strings"
)

const INITIAL_SIZE = 8

type Queue[T any] struct {
	data       []T
	head, tail int
	length     int
}

// NewQueue creates a new queue with the given elements (if any).
func NewQueue[T any](values ...T) *Queue[T] {
	q := &Queue[T]{}
	for _, v := range values {
		q.Append(v)
	}
	return q
}

// Pop removes and returns an element from the right side of the queue. If no elements are present, Pop returns the zero Value with false.
func (q *Queue[T]) Pop() (T, bool) {
	var zero T
	if q.length == 0 {
		return zero, false
	}
	q.tail = q.getIndex(q.tail - 1)
	value := q.data[q.tail]
	q.data[q.tail] = zero
	q.length--
	q.checkShrink()
	return value, true
}

// PopLeft removes and returns an element from the left side of the queue. If no elements are present, PopLeft returns the zero Value and false.
func (q *Queue[T]) PopLeft() (T, bool) {
	var zero T
	if q.length == 0 {
		return zero, false
	}
	value := q.data[q.head]
	q.data[q.head] = zero
	q.head = q.getIndex(q.head + 1)
	q.length--
	q.checkShrink()
	return value, true
}

// Append adds the Value to the right side of the queue.
func (q *Queue[T]) Append(value T) {
	q.checkGrow()
	q.data[q.tail] = value
	q.tail = q.getIndex(q.tail + 1)
	q.length++
}

// AppendLeft adds the Value to the left side of the queue.
func (q *Queue[T]) AppendLeft(value T) {
	q.checkGrow()
	q.head = q.getIndex(q.head - 1)
	q.data[q.head] = value
	q.length++
}

// Extend adds the values in the given order to the right side of the queue.
func (q *Queue[T]) Extend(values ...T) {
	for _, v := range values {
		q.Append(v)
	}
}

// ExtendLeft add the values to the left side of the queue. Note, the series of left appends results in reversing the order of given elements.
func (q *Queue[T]) ExtendLeft(values ...T) {
	for _, v := range values {
		q.AppendLeft(v)
	}
}

// Clear removes all elements from the queue leaving it with length 0.
func (q *Queue[T]) Clear() {
	q.data = nil
	q.head = 0
	q.tail = 0
	q.length = 0
}

// Len returns the number of items in the queue.
func (q *Queue[T]) Len() int {
	return q.length
}

// Empty returns true of the queue is empty; false otherwise
func (q *Queue[T]) Empty() bool {
	return q.length == 0
}

// Peek returns the first item in the queue without removing it.
func (q *Queue[T]) Peek() T {
	var zero T
	if q.length == 0 {
		return zero
	}
	return q.data[q.head]
}

// PeekLast returns the last item in the queue without removing it.
func (q *Queue[T]) PeekLast() T {
	var zero T
	if q.length == 0 {
		return zero
	}
	return q.data[q.getIndex(q.tail-1)]
}

func (q *Queue[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		id := q.head
		for i := range q.length {
			if !yield(i, q.data[id]) {
				break
			}
			id = q.getIndex(id + 1)
		}
	}
}

func (q *Queue[T]) List() []T {
	slice := make([]T, q.Len())
	for i, v := range q.All() {
		slice[i] = v
	}
	return slice
}

func (q *Queue[T]) String() string {
	buf := make([]string, q.length)
	for i, v := range q.All() {
		buf[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(buf, ",")
}

func (q *Queue[T]) getIndex(i int) int {
	return i & (len(q.data) - 1) // modulo for queues with len(ring buffer) power of 2
}

func (q *Queue[T]) checkShrink() {
	if q.length < len(q.data)>>2 {
		q.setSize(len(q.data) >> 1)
	}
}

func (q *Queue[T]) checkGrow() {
	if q.length == len(q.data) {
		q.setSize(len(q.data) << 1)
	}
}

func (q *Queue[T]) setSize(newSize int) {
	newSize = utils.Max(newSize, INITIAL_SIZE)
	if len(q.data) == newSize {
		return
	}
	if newSize&(newSize-1) != 0 {
		panic("new size must be 2^n")
	}

	newData := make([]T, newSize)
	//for i, item := range q.All() {
	//	newData[i] = item
	//}

	if q.head < q.tail {
		copy(newData, q.data[q.head:q.tail])
	} else {
		copy(newData, q.data[q.head:])
		copy(newData[len(q.data)-q.head:], q.data[:q.tail])
	}

	q.data = newData
	q.head = 0
	q.tail = q.length
}
