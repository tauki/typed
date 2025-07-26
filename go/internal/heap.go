package internal

import "container/heap"

type Comparator[T any] func(a, b T) bool

type Heap[T any] struct {
	items      []T
	comparator Comparator[T]
}

func NewHeap[T any](cmp func(a, b T) bool) *Heap[T] {
	h := &Heap[T]{comparator: cmp}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.items)
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.comparator(h.items[i], h.items[j])
}

func (h *Heap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *Heap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

func (h *Heap[T]) Pop() any {
	n := len(h.items)
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

func (h *Heap[T]) Peek() (T, bool) {
	var zero T
	if h.Len() == 0 {
		return zero, false
	}
	return h.items[0], true
}

func (h *Heap[T]) ItemsCopy() []T {
	cp := make([]T, len(h.items))
	copy(cp, h.items)
	return cp
}
