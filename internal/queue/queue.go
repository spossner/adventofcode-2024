package queue

import (
	"fmt"
	"iter"
)

type (
	Node[T any] struct {
		Value          T
		next, previous *Node[T]
	}

	Queue[T any] struct {
		start, end *Node[T]
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

// Pop removes and returns an element from the right side of the queue. If no elements are present, Pop returns the zero Value with false.
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
	return n.Value, true
}

// PopLeft removes and returns an element from the left side of the queue. If no elements are present, PopLeft returns the zero Value and false.
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
	return n.Value, true
}

// Append adds the Value to the right side of the queue.
func (q *Queue[T]) Append(value T) {
	n := &Node[T]{value, nil, q.end}
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

// AppendLeft adds the Value to the left side of the queue.
func (q *Queue[T]) AppendLeft(value T) {
	n := &Node[T]{value, q.start, nil}
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
	return q.start.Value
}

// InsertBefore inserts the given Value before the given node.
// Note that InsertBefore expects the node to be part of this queue.
func (q *Queue[T]) InsertBefore(node *Node[T], value T) {
	n := &Node[T]{value, node, node.previous}
	if node.previous != nil {
		node.previous.next = n
	} else {
		q.start = n
	}
	node.previous = n

	q.length++
}

// InsertAfter inserts the given Value after the given node.
// Note that InsertAfter expects the node to be part of this queue.
func (q *Queue[T]) InsertAfter(node *Node[T], value T) {
	n := &Node[T]{value, node.next, node}
	if node.next != nil {
		node.next.previous = n
	} else {
		q.end = n
	}
	node.next = n
	q.length++
}

func (q *Queue[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		el := q.start
		i := 0
		for el != nil {
			if !yield(i, el.Value) {
				break
			}
			el = el.next
			i++
		}
	}
}

func (q *Queue[T]) AllNodes() iter.Seq2[int, *Node[T]] {
	return func(yield func(int, *Node[T]) bool) {
		el := q.start
		i := 0
		for el != nil {
			if !yield(i, el) {
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

func (q *Queue[T]) String() string {
	buf := make([]string, q.length)
	for i, v := range q.All() {
		buf[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%v", buf)
}
