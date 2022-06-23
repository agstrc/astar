// Package pqueue provides a priority queue implementation.
package pqueue

import "github.com/agstrc/heuristic-search/pqueue/heap"

type item[T any] struct {
	value    T
	priority int
}

// innerQueue is a list on which heap.Interface is implemented
type innerQueue[T any] []item[T]

func (iq innerQueue[T]) Len() int { return len(iq) }

func (iq innerQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return iq[i].priority > iq[j].priority
}

func (iq innerQueue[T]) Swap(i, j int) {
	iq[i], iq[j] = iq[j], iq[i]
}

func (iq *innerQueue[T]) Push(x item[T]) {
	*iq = append(*iq, x)
}

func (iq *innerQueue[T]) Pop() item[T] {
	popped := (*iq)[len(*iq)-1]
	*iq = (*iq)[:len(*iq)-1]
	return popped
}

var _ heap.Interface[item[any]] = &innerQueue[any]{}

// PriorityQueue is a generic priority queue. It
type PriorityQueue[T any] struct {
	innerQueue[T]
}

// Push adds a value to the queue with a given priority.
func (pq *PriorityQueue[T]) Push(value T, priority int) {
	i := item[T]{value: value, priority: priority}
	heap.Push[item[T]](&pq.innerQueue, i)
}

// Pop removes the highest priority item from the queue and returns it.
func (pq *PriorityQueue[T]) Pop() T {
	i := heap.Pop[item[T]](&pq.innerQueue)
	return i.value
}

// Empty reports whether the queue is empty.
func (pq *PriorityQueue[T]) Empty() bool {
	return len(pq.innerQueue) == 0
}
