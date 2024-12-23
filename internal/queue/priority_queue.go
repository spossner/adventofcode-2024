package queue

import (
	"container/heap"
	"fmt"
	"golang.org/x/exp/constraints"
)

type Item[T constraints.Ordered, U any] struct {
	Payload U
	Costs   T
	index   int // used by heap
}

func (i *Item[T, U]) String() string {
	return fmt.Sprintf("%v, %v (idx %d)", i.Costs, i.Payload, i.index)
}

func NewItem[T constraints.Ordered, U any](costs T, payload U) *Item[T, U] {
	return &Item[T, U]{Costs: costs, Payload: payload}
}

type PriorityQueue[T constraints.Ordered, U any] []*Item[T, U]

func (pq PriorityQueue[T, U]) Len() int { return len(pq) }
func (pq PriorityQueue[T, U]) Less(i, j int) bool {
	return pq[i].Costs < pq[j].Costs
}
func (pq PriorityQueue[T, U]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq PriorityQueue[T, U]) Empty() bool {
	return len(pq) == 0
}

func (pq *PriorityQueue[T, U]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T, U])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T, U]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type PQ[U any] struct {
	queue PriorityQueue[int, U]
}

func NewPQ[U any](values ...Item[int, U]) *PQ[U] {
	q := PQ[U]{
		queue: make(PriorityQueue[int, U], 0),
	}

	for _, value := range values {
		q.queue = append(q.queue, &value)
	}
	heap.Init(&q.queue)
	return &q
}

func (q *PQ[U]) Push(costs int, value U) {
	heap.Push(&q.queue, NewItem[int, U](costs, value))
}

func (q *PQ[U]) Pop() U {
	return heap.Pop(&q.queue).(*Item[int, U]).Payload
}

func (q *PQ[U]) Empty() bool {
	return q.queue.Empty()
}
