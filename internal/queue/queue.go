package queue

import (
	"iter"
)

type (
	node[T any] struct {
		value          T
		next, previous *node[T]
	}

	Queue[T any] struct {
		start, end *node[T]
		length     int
	}
)

// NewQueue creates a new queue with the given elements (if any).
func NewQueue[T any](values ...T) *Queue[T] {
	q := &Queue[T]{nil, nil, 0}
	for _, v := range values {
		q.Append(v)
	}
	return q
}

// Pop removes and returns an element from the right side of the queue. If no elements are present, Pop returns the zero value with false.
func (q *Queue[T]) Pop() (T, bool) {
	var zero T
	if q.length == 0 {
		return zero, false
	}
	n := q.end
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.end = q.end.previous
		q.end.next = nil
	}
	q.length--
	return n.value, true
}

// PopLeft removes and returns an element from the left side of the queue. If no elements are present, PopLeft returns the zero value and false.
func (q *Queue[T]) PopLeft() (T, bool) {
	var zero T
	if q.length == 0 {
		return zero, false
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
		q.start.previous = nil
	}
	q.length--
	return n.value, true
}

// Append adds the value to the right side of the queue.
func (q *Queue[T]) Append(value T) {
	n := &node[T]{value, nil, q.end}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// AppendLeft adds the value to the left side of the queue.
func (q *Queue[T]) AppendLeft(value T) {
	n := &node[T]{value, q.start, nil}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.start.previous = n
		q.start = n
	}
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
	q.start = nil
	q.end = nil
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
func (q *Queue[T]) Peek() any {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}

func (q *Queue[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		el := q.start
		i := 0
		for el != nil {
			if !yield(i, el.value) {
				break
			}
			el = el.next
			i++
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
